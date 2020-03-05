package funchttp

import (
	_ "fmt"
	"github.com/formeo/XlsForXMLHttp/commonfunc"
	"github.com/formeo/XlsForXMLHttp/config"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type HttpApp struct {
	log  *zap.Logger
	conf *config.Config
}

func NewHttpApp(log *zap.Logger, conf *config.Config) *HttpApp {
	return &HttpApp{
		log:  log,
		conf: conf,
	}
}

func (h *HttpApp) GetOnlyZBParse(w http.ResponseWriter, r *http.Request) {

	_, err := commonfunc.ParceZB(h.conf.PathToFiles)

	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	res, err := commonfunc.MakeBackupXML()
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	w.Write(res)

}

func (h *HttpApp) Test(w http.ResponseWriter, r *http.Request) {
	h.log.Info(h.conf.PathToFiles)
	io.WriteString(w, h.conf.PathToFiles)
}

//GetOnlyZB выдает XML файл со списком файлов для Зап.Сиб
func (h *HttpApp) GetOnlyZB(w http.ResponseWriter, r *http.Request) {
	var res []byte
	res, err := commonfunc.MakeXMLFromXLSZBvbs(h.conf.PathToFiles)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)

	}
	w.Write(res)

}

//GetOnly выдает XML файл со списком файл
func (h *HttpApp) GetOnly(w http.ResponseWriter, r *http.Request) {
	res, err := commonfunc.MakeXMLFromXLSvbs(h.conf.PathToFiles)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	w.Write(res)

}

//ToArch функция отправляет файлы в архив
func (h *HttpApp) ToArch(w http.ResponseWriter, r *http.Request) {
	err := commonfunc.CopyToArchive(h.conf.PathToFiles, h.conf.PathToBackupFolder)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	res, err := commonfunc.MakeBackupXML()
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	w.Write(res)
}

//ClearDir функция удаляет файлы в сетевой директории
func (h *HttpApp) ClearDir(w http.ResponseWriter, r *http.Request) {
	err := commonfunc.ClearDirectory(h.conf.PathToFiles, h.conf.PathToClearDir)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	res, err := commonfunc.MakeBackupXML()
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	w.Write(res)
}

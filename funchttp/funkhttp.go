package funchttp

import (
	"github.com/formeo/XlsForOra/commonfunc"
	"github.com/formeo/XlsForOra/config"
	_ "fmt"
	"net/http"
	"log"
	"io"
)

func GetOnlyZBParce(w http.ResponseWriter, r *http.Request) {

	_, err := commonfunc.ParceZB(config.Current().PathToFiles)

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

func Test(w http.ResponseWriter, r *http.Request) {	
	log.Println(config.Current().PathToFiles)	
	io.WriteString(w, config.Current().PathToFiles)
}

//GetOnlyZB выдает XML файл со списком файлов для Зап.Сиб
func GetOnlyZB(w http.ResponseWriter, r *http.Request) {
	var res []byte
	//res, err := commonfunc.MakeXMLFromXLSZB(config.Current().PathToFiles, "44")
	res, err := commonfunc.MakeXMLFromXLSZBvbs(config.Current().PathToFiles)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)

	}
	w.Write(res)
	
}

//GetOnly выдает XML файл со списком файл
func GetOnly(w http.ResponseWriter, r *http.Request) {
	res, err := commonfunc.MakeXMLFromXLSvbs(config.Current().PathToFiles)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	w.Write(res)

}

//ToArch функция отправляет файлы в архив
func ToArch(w http.ResponseWriter, r *http.Request) {
	err := commonfunc.CopyToArchive(config.Current().PathToFiles, config.Current().PathToBackupFolder)
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
func ClearDir(w http.ResponseWriter, r *http.Request) {
	err := commonfunc.ClearDirectory(config.Current().PathToFiles,config.Current().PathToClearDir)
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
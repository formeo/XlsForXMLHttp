package funchttp

import (
	"XlsForOra/commonfunc"
	"XlsForOra/config"
	_ "fmt"
	"net/http"
)

//GetOnlyZB выдает XML файл со списком файлов для Зап.Сиб
func GetOnlyZB(w http.ResponseWriter, r *http.Request) {
	res, err := commonfunc.MakeXMLFromXLSZB(config.Current().PathToFiles)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXML(err.Error())
		w.Write(resErr)
	}
	w.Write(res)
	//fmt.Fprintf(w, config.Current().PathToFiles)
}

//GetOnly выдает XML файл со списком файл
func GetOnly(w http.ResponseWriter, r *http.Request) {
	res, err := commonfunc.MakeXMLFromXLS(config.Current().PathToFiles)
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

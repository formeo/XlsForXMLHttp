package funchttp

import (
	"XlsForOra/commonfunc"
	"XlsForOra/config"
	"net/http"
)

//GetOnly выдает XML файл со списком файлов
func GetOnly(w http.ResponseWriter, r *http.Request) {
	res, err := commonfunc.MakeXmlFromXLS(config.Current().PathToFiles)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXml(err.Error())
		w.Write(resErr)
	}
	w.Write(res)

}

//ToArch функция отправляет файлы в архив
func ToArch(w http.ResponseWriter, r *http.Request) {
	err := commonfunc.CopyToArchive(config.Current().PathToFiles, config.Current().PathToBackupFolder)
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXml(err.Error())
		w.Write(resErr)
	}
	res, err := commonfunc.MakeBackupXml()
	if err != nil {
		resErr, _ := commonfunc.MakeErrorXml(err.Error())
		w.Write(resErr)
	}
	w.Write(res)
}

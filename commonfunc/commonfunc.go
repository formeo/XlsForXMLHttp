package commonfunc

import (
	"XlsForOra/filesutil"
	"XlsForOra/xmlstruck"
	"encoding/xml"
	"log"
	"os"
)

//CopyToArchive копирует все файлы в архив и удаляет из источника
func CopyToArchive(PathDir string, PathArchDir string) (err error) {
	dir, err := os.Open(PathDir)
	if err != nil {
		return err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fi := range fileInfos {
		if !fi.IsDir() {
			filesutil.FileCopy(PathDir+fi.Name(), PathArchDir+fi.Name(), true)
			err := os.Remove(PathDir + fi.Name())
			if err != nil {
				return err
			}

		}
	}

	return nil
}

//MakeErrorXML формирует XML с ответом если произошла ошибка
func MakeErrorXML(InErr string) (res []byte, err error) {
	v := &xmlstruck.Servers{Version: "1", Code: "1", Message: InErr}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		return nil, err
	}

	mySlice := []byte(xml.Header + string(output))
	res = mySlice
	return res, nil

}

//MakeBackupXML формирует ответ после функции бэкапа
func MakeBackupXML() (res []byte, err error) {
	v := &xmlstruck.Servers{Version: "1", Code: "200", Message: "Success"}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		return nil, err
	}
	mySlice := []byte(xml.Header + string(output))
	res = mySlice
	return res, nil

}

//MakeXMLFromXLS формирует окончательную XML
func MakeXMLFromXLS(PathDir string) (res []byte, err error) {

	var s string

	v := &xmlstruck.Servers{Version: "1", Code: "0", Message: ""}
	dir, err := os.Open(PathDir)
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		if !fi.IsDir() {
			item, err := filesutil.FileToRow(PathDir+fi.Name(), fi.Name())
			log.Println(item)
			if err != nil {
				return nil, err

			}
			v.Svs = append(v.Svs, *item)

		}
	}
	if len(v.Svs) == 0 {
		v.Code = "404"
		v.Message = "files not found"
	}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		return nil, err
	}
	/*s = xml.Header*/
	s += string(output)
	//mySlice := []byte(xml.Header + string(output))
	mySlice := []byte(xml.Header + string(output))
	res = mySlice
	return res, nil

}

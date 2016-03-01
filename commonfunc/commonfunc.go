package commonfunc

import (
	"github.com/formeo/XlsForXMLHttp/filesutil"
	"github.com/formeo/XlsForXMLHttp/xmlstruck"
	"encoding/xml"	     
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	_ "strconv"
	_ "strings"
	"time"
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

func ParceZB(PathDir string) (Folder string, err error) {
	t := time.Now()

	Folder = t.Format("20060102150405")
	log.Println("make folder ", Folder)
	if err := os.MkdirAll(PathDir+Folder, 0777); err != nil {
		return Folder, err
	}
	result, err := filesutil.DelForMask(PathDir+"\\"+Folder+"\\", "Sheet")
	if err != nil {
		return Folder, err
	}
	if !result {
		return Folder, err
	}
	log.Println("start parce")
	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\filename.vbs", PathDir+"drvscrp\\", PathDir, Folder)
	err = cmd.Run()
	if err != nil {
		return Folder, err
	}
	log.Println("finish parce")
	return Folder, nil
}

func MakeXMLtest() (res []byte, err error) {

	res, err = ioutil.ReadFile("C:\\paynotes\\20160127092609\\zapsib.xml")
	if err != nil {
		return nil, err
	}
	return res, nil
}

//MakeXMLFromXLSZBvbs
func MakeXMLFromXLSZBvbs(PathDir string) (res []byte, err error) {
	log.Println("MakeXMLFromXLSZB")
	Folder, err := ParceZB(PathDir)

	if err != nil {
		return nil, err
	}
	log.Println("start make xml")
	log.Println("PathDir", PathDir)
	log.Println("Folder", Folder)

	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\2.vbs", PathDir+Folder+"\\", Folder)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	log.Println("finish xml")
	res, err = ioutil.ReadFile(PathDir + Folder + "\\zapsib.xml")
	if err != nil {
		return nil, err
	}
	return res, nil
}

//MakeXMLFromXLSZB формирует окончательную XML
func MakeXMLFromXLSZB(PathDir string, Folder string) (res []byte, err error) {
	/*var (
		s   string
		exl *excel.MSO
	)
	v := &xmlstruck.Servers{Version: "1", Code: "0", Message: ""}
	dir, err := os.Open(PathDir + "\\" + Folder + "\\")
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	//чистим папку

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	log.Println("start parce")

	/*ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("Excel.Application")
	excel, _ := unknown.QueryInterface(ole.IID_IDispatch)
	oleutil.PutProperty(excel, "Visible", false)
	workbooks := oleutil.MustGetProperty(excel, "Workbooks").ToIDispatch()

	log.Println("exl", exl)

	for _, fi := range fileInfos {
		if !fi.IsDir() {

			item, err := filesutil.FileToRowZPNew(PathDir+"\\"+Folder+"\\"+fi.Name(), "\\"+Folder+"\\"+fi.Name(), excel, workbooks)
			log.Println(item)
			if err != nil {
				return nil, err
			}
			if item != nil {
				v.Svs = append(v.Svs, *item)
			}

		}
	}
	//err = exl.Quit()

	workbooks.Release()
	excel.Release()
	ole.CoUninitialize()*/

	/*log.Println("err Quit", err)

	if len(v.Svs) == 0 {
		v.Code = "404"
		v.Message = "files not found"
	}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		return nil, err
	}
	s += string(output)
	mySlice := []byte(xml.Header + string(output))
	res = mySlice*/
	return nil, nil
}

//MakeXMLFromXLS формирует окончательную XML
/*func MakeXMLFromXLS(PathDir string) (res []byte, err error) {
	var s string
	log.Println("MakeXMLFromSBER")
	bres, err := filesutil.DelForMask(PathDir, "zap")
	if bres != true {
		if err != nil {
			return nil, err
		}
	}
	v := &xmlstruck.Servers{Version: "1", Code: "0", Message: ""}
	dir, err := os.Open(PathDir)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	for _, fi := range fileInfos {
		if !fi.IsDir() {
			if fi.Name() != "RurPaymentDemand.xls" {
				item, err := filesutil.FileToRow(PathDir+fi.Name(), fi.Name())
				log.Println(item)
				if err != nil {
					return nil, err

				}
				v.Svs = append(v.Svs, *item)
			}
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
/*	s += string(output)
	//mySlice := []byte(xml.Header + string(output))
	mySlice := []byte(xml.Header + string(output))
	res = mySlice
	return res, nil

}*/

func MakeXMLFromXLSvbs(PathDir string) (res []byte, err error) {
	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\sber.vbs", PathDir)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	res, err = ioutil.ReadFile(PathDir + "\\sber.xml")
	if err != nil {
		return nil, err
	}
	return res, nil

}

func ClearDirectory(PathDir string, PathDirToClear string) (err error) {

	log.Println("start delete")
	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\delete.vbs", PathDirToClear)
	err = cmd.Run()
	if err != nil {
		return err
	}
	log.Println("finish delete")
	return nil

}

package commonfunc

import (
	"encoding/xml"
	"github.com/formeo/XlsForXMLHttp/filesutil"
	"github.com/formeo/XlsForXMLHttp/xmlstruck"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	_ "strconv"
	_ "strings"
	"time"
)

type CommFunc struct {
	filesUtil *filesutil.UtilsApp
	log       *zap.Logger
}

func NewCommFunc(filesUtil *filesutil.UtilsApp, log *zap.Logger) *CommFunc {
	return &CommFunc{
		log:       log,
		filesUtil: filesUtil,
	}
}

//CopyToArchive копирует все файлы в архив и удаляет из источника
func (c *CommFunc) CopyToArchive(PathDir string, PathArchDir string) (err error) {
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
			c.filesUtil.FileCopy(PathDir+fi.Name(), PathArchDir+fi.Name(), true)
			err := os.Remove(PathDir + fi.Name())
			if err != nil {
				return err
			}

		}
	}

	return nil
}

//MakeErrorXML формирует XML с ответом если произошла ошибка
func (c *CommFunc) MakeErrorXML(InErr string) (res []byte, err error) {
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
func (c *CommFunc) MakeBackupXML() (res []byte, err error) {
	v := &xmlstruck.Servers{Version: "1", Code: "200", Message: "Success"}
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		return nil, err
	}
	mySlice := []byte(xml.Header + string(output))
	res = mySlice
	return res, nil

}

func (c *CommFunc) ParceZB(PathDir string) (Folder string, err error) {
	t := time.Now()

	Folder = t.Format("20060102150405")
	log.Println("make folder ", Folder)
	if err := os.MkdirAll(PathDir+Folder, 0777); err != nil {
		return Folder, err
	}
	result, err := c.filesUtil.DeleteAtMask(PathDir+"\\"+Folder+"\\", "Sheet")
	if err != nil {
		return Folder, err
	}
	if !result {
		return Folder, err
	}
	c.log.Info("start parce")
	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\filename.vbs", PathDir+"drvscrp\\", PathDir, Folder)
	err = cmd.Run()
	if err != nil {
		return Folder, err
	}
	c.log.Info("finish parce")
	return Folder, nil
}

//MakeXMLFromXLSZBvbs
func (c *CommFunc) MakeXMLFromXLSZBvbs(PathDir string) (res []byte, err error) {
	c.log.Info("MakeXMLFromXLSZB")
	Folder, err := c.ParceZB(PathDir)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\2.vbs", PathDir+Folder+"\\", Folder)
	err = cmd.Run()
	if err != nil {
		return nil, err
	}
	c.log.Info("finish xml")
	res, err = ioutil.ReadFile(PathDir + Folder + "\\zapsib.xml")
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *CommFunc) MakeXMLFromXLSvbs(PathDir string) (res []byte, err error) {
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

func (c *CommFunc) ClearDirectory(PathDir string, PathDirToClear string) (err error) {
	c.log.Info("start delete")
	cmd := exec.Command("c:\\Windows\\System32\\cscript.exe", PathDir+"drvscrp\\delete.vbs", PathDirToClear)
	err = cmd.Run()
	if err != nil {
		return err
	}
	c.log.Info("finish delete")
	return nil

}

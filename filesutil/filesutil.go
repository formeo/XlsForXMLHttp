package filesutil

import (
	_ "errors"
	_ "github.com/formeo/XlsForXMLHttp/xmlstruck"
	"io"
	_ "log"
	"os"
	_ "strconv"
	"strings"
)

type UtilsApp struct {

}

func NewUtilsApp() *UtilsApp {

	return &UtilsApp{}
}


//копирует файлы
func(f *UtilsApp) FileCopy(source, dest string, overwrite bool) (bool, error) {
	in, err := os.Open(source)
	if in == nil {
		return false, nil
	}
	defer in.Close()
	if err != nil {
		return false, err
	}
	if _, e := os.Stat(dest); overwrite && (e != nil) && os.IsNotExist(err) {
		e = os.Remove(dest)
		if e != nil {
			return false, e
		}
	}
	out, eout := os.Create(dest)
	if out == nil {
		return false, nil
	}
	defer out.Close()
	if eout != nil {
		return false, eout
	}
	if _, err = io.Copy(out, in); err != nil {
		return false, err
	}
	if err = out.Sync(); err != nil {
		return false, err
	}
	return true, nil
}

//DelForMask удаляет файл по маске
func(f *UtilsApp) DeleteAtMask(source, mask string) (bool, error) {

	dir, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return false, err
	}
	for _, fi := range fileInfos {
		if !fi.IsDir() {
			if strings.Index(fi.Name(), mask) != -1 {
				err := os.Remove(source + fi.Name())
				if err != nil {
					return false, err
				}

			}

		}
	}
	return true, nil

}

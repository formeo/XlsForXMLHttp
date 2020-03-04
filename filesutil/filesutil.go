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

//копирует файлы
func FileCopy(source, dest string, overwrite bool) (bool, error) {
	in, err := os.Open(source)
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
func DelForMask(source, mask string) (bool, error) {

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

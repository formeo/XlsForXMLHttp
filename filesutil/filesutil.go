package filesutil

import (
	"XlsForOra/xmlstruck"
	"errors"
	"github.com/aswjh/excel"
	"io"
	"os"
	"strconv"
	"strings"
)

//FileToRow Извлекает из XLS файла необходимую информацию
func FileToRow(FilePath string, FileName string) (*xmlstruck.Files, error) {

	option := excel.Option{"Visible": false, "DisplayAlerts": false}
	xl, err := excel.Open(FilePath, option)
	if err != nil {

		return nil, err
	}
	defer xl.Quit()
	sheet, _ := xl.Sheet(1)
	res := &xmlstruck.Files{Number: sheet.MustCells(6, 12), Date: sheet.MustCells(5, 15), Summ: sheet.MustCells(11, 21), Filename: FileName}
	return res, nil
}

//DelSheet разбивает файл на несколько фрагментов по листам
func DelSheet(number int, sourceFile string, dirFiles string) (err error) {
	option := excel.Option{"Visible": false, "DisplayAlerts": false}
	res, err := FileCopy(sourceFile, dirFiles+"zap"+strconv.Itoa(number)+".xls", true)
	if res != true {
		if err != nil {
			return err
		} else {
			err := errors.New("files not copied")
			return err
		}

	}
	xlnew, _ := excel.Open(dirFiles+"zap"+strconv.Itoa(number)+".xls", option)
	defer xlnew.Quit()
	j := 1
	for xlnew.CountSheets() > 1 {
		shh, _ := xlnew.Sheet(j)
		if shh.Name() != "Sheet"+strconv.Itoa(number-1) {
			s, _ := xlnew.Sheet("Sheet" + strconv.Itoa(number-1))
			s.Delete()

		} else {
			j = j + 1
		}
	}
	xlnew.Save()
	return nil
}

//FileToRowZP Извлекает из XLS файла необходимую информацию
func FileToRowZP(FilePath string, FileName string) (*xmlstruck.Files, error) {
	var res *xmlstruck.Files
	option := excel.Option{"Visible": false, "DisplayAlerts": false}
	xlparse, err := excel.Open(FilePath, option)
	if err != nil {
		return nil, err
	}
	defer xlparse.Quit()
	sheet, _ := xlparse.Sheet(1)
	if sheet.MustCells(3, 1) == "ПЛАТЕЖНОЕ ПОРУЧЕНИЕ" {
		s := strings.Trim(sheet.MustCells(3, 4), "№ ")

		res = &xmlstruck.Files{Number: s, Date: sheet.MustCells(3, 6), Summ: sheet.MustCells(8, 7), Filename: FileName}

	}

	if sheet.MustCells(5, 1) == "ПЛАТЕЖНОЕ ТРЕБОВАНИЕ №" {
		res = &xmlstruck.Files{Number: sheet.MustCells(5, 12), Date: sheet.MustCells(5, 18), Summ: sheet.MustCells(14, 24), Filename: FileName}

	}
	return res, nil
}

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

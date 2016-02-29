package filesutil

import (
	"github.com/formeo/XlsForXMLHttp/xmlstruck"
	"errors"
	"github.com/aswjh/excel"
	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"io"
	"log"
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
func FileToRowZP(FilePath string, FileName string, Exl *excel.MSO) (*xmlstruck.Files, error) {
	var res *xmlstruck.Files

	log.Println("OpenWorkBook")
	wb, err := Exl.OpenWorkBook(FilePath)
	log.Println(err)
	log.Println("Activate")
	err = wb.Activate()
	log.Println(err)
	sheet, err := wb.SelectSheet(1)
	log.Println(sheet)
	if err != nil {
		return nil, err
	}
	if sheet.MustCellsNew(3, 1) == "ПЛАТЕЖНОЕ ПОРУЧЕНИЕ" {
		s := strings.Trim(sheet.MustCellsNew(3, 4), "№ ")

		res = &xmlstruck.Files{Number: s, Date: sheet.MustCellsNew(3, 6), Summ: sheet.MustCellsNew(8, 7), Filename: FileName}

	}	

	err = wb.Close()
	log.Println(err)
	return res, nil
}

func FileToRowZPNew(FilePath string, FileName string, excel, workbooks *ole.IDispatch) (*xmlstruck.Files, error) {
	var res *xmlstruck.Files

	log.Println("OpenWorkBookNEEEEEW")
	log.Println("Openworkbook")
	workbook, err := oleutil.CallMethod(workbooks, "Open", FilePath)
	log.Println(err)
	defer workbook.ToIDispatch().Release()
	log.Println("Openworksheet")
	worksheet := oleutil.MustGetProperty(workbook.ToIDispatch(), "Worksheets", 1).ToIDispatch()
	defer worksheet.Release()
	log.Println(worksheet)	
	log.Println(err)
	log.Println("Activate")
	

	cell := oleutil.MustGetProperty(worksheet, "Cells", 3, 1).ToIDispatch()
	val, err := oleutil.GetProperty(cell, "Value")

	if val.ToString() == "ПЛАТЕЖНОЕ ПОРУЧЕНИЕ" {
		cell := oleutil.MustGetProperty(worksheet, "Cells", 3, 4).ToIDispatch()
		val, err := oleutil.GetProperty(cell, "Value")
		log.Println(err)
		s := strings.Trim(val.ToString(), "№ ")

		cell = oleutil.MustGetProperty(worksheet, "Cells", 3, 6).ToIDispatch()
		date, err := oleutil.GetProperty(cell, "Value")

		cell = oleutil.MustGetProperty(worksheet, "Cells", 8, 7).ToIDispatch()
		summ, err := oleutil.GetProperty(cell, "Value")

		res = &xmlstruck.Files{Number: s, Date: date.ToString(), Summ: summ.ToString(), Filename: FileName}

	}

	
	log.Println(err)
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

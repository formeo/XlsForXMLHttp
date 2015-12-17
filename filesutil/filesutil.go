package filesutil

import (
	"XlsForOra/xmlstruck"
	"github.com/aswjh/excel"
	"io"
	"os"
)

//FileCopy Копирует файлы в паку бэкапа и уладяет их из целевой папки
func FileCopy(source, dest string, overwrite bool) (bool, error) {
	in, err := os.Open(source)
	if err != nil {
		return false, err
	}
	defer in.Close()
	if _, e := os.Stat(dest); overwrite && (e != nil) && os.IsNotExist(err) {
		e = os.Remove(dest)
		if e != nil {
			return false, e
		}
	}
	out, eout := os.Create(dest)
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
	//{Version: "1", Code: "0", Message: ""}
	/*res.Number = sheet.MustCells(6, 12)
	res.Date = sheet.MustCells(5, 15)
	res.Summ = sheet.MustCells(11, 21)
	res.Filename = FileName*/

	return res, nil

}

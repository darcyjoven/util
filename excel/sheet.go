package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

// 返回工作簿的列和行
func getMaxHV(f *excelize.File, v string) (int, int) {
	s, err := f.Rows(v)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	//列，行
	i, j := 0, 1
	if s.Next() {
		temp, err := s.Columns()
		if err != nil {
			log.Println(err)
			return 0, 0
		}
		i = len(temp)
	}
	for s.Next() {
		j++
	}
	return i, j
}

// 没有上面的快
func getByRows(f *excelize.File, v string) (int, int) {
	s, err := f.GetRows(v)
	if err != nil {
		log.Println(err)
		return 0, 0
	}
	i, j := len(s), 0
	if i > 0 {
		j = len(s[0])
	}
	return i, j
}

// delete 删除默认的Sheet1
func clearSheet(f *excelize.File) error {
	v, err := f.GetCellValue("Sheet1", "A1")
	if err != nil {
		// 可能无这个sheet直接返回
		return nil
	}
	if v != "" {
		// 不为空，这个sheet在使用，不清理
		return nil
	}
	// 有sheet1，且A1为空
	v, err = f.GetCellValue("Sheet1", "A2")
	if err != nil {
		// 遇到异常
		// 删除sheet
		f.DeleteSheet("Sheet1")
		return nil
	}
	if v != "" {
		// 不为空，这个sheet在使用，不清理
		return nil
	}
	f.DeleteSheet("Sheet1")
	return nil
}

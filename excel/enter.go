package excel

import (
	"fmt"
	"log"
	"strconv"

	"github.com/darcyjoven/util"
	"github.com/xuri/excelize/v2"
)

type excelFile struct {
	file *excelize.File
	name string
}

// name string 文件目录
func NewExcel(name string) *excelFile {
	return &excelFile{
		file: excelize.NewFile(),
		name: name,
	}
}

func (f *excelFile) AddSheet(name string, data [][]any) error {
	//  新增sheet
	s := f.file.NewSheet(name)
	if s == 0 {
		return fmt.Errorf("创建shet失败:%s", name)
	}
	// 流式写入
	stream, err := f.file.NewStreamWriter(name)
	defer stream.Flush()
	if err != nil {
		log.Println(err)
		return err
	}
	for i, v := range data {
		err := stream.SetRow("A"+strconv.Itoa(i+1), v)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func (f *excelFile) Save() string {
	defer f.file.Close()
	err := clearSheet(f.file)
	if err != nil {
		log.Println(err)
		return " "
	}
	err = f.file.SaveAs(f.name)
	if err != nil {
		log.Println(err)
		return ""
	}
	return f.file.Path
}

// 有资料的位置设置边框
func (f *excelFile) SetBorder() error {
	for _, v := range f.file.GetSheetList() {
		i, j := getMaxHV(f.file, v)
		if i*j == 0 {
			return fmt.Errorf("无资料")
		}
		err := f.file.SetCellStyle(v, "A1", util.NumToLetter(i)+strconv.Itoa(j), allborder(f.file))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

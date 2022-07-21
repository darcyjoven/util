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
func SetBorder(f1 *excelize.File) error {
	for _, v := range f1.GetSheetList() {
		i, j := getMaxHV(f1, v)
		if i*j == 0 {
			return fmt.Errorf("无资料")
		}
		err := f1.SetCellStyle(v, "A1", util.NumToLetter(i)+strconv.Itoa(j), allborder(f1))
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// open
func OpenFile(f string) (*excelize.File, error) {
	return excelize.OpenFile(f)
}

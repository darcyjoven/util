package excel

import (
	"log"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestBorder(t *testing.T) {
	f, err := excelize.OpenFile("demo.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	e := NewExcel("name")
	e.file = f
	SetBorder(f)
	defer f.Save()
}

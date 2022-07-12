package excel

import (
	"log"
	"testing"

	"github.com/xuri/excelize/v2"
)

func BenchmarkRows(b *testing.B) {
	f, err := excelize.OpenFile("demo.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		getMaxHV(f, "Sheet1")
		// fmt.Print(i, ":", j)
	}
}

func BenchmarkGetRows(b *testing.B) {
	f, err := excelize.OpenFile("demo.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < b.N; i++ {
		getByRows(f, "Sheet1")
	}
}

package excel

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func allborder(f *excelize.File) int {
	sty, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{
				Type:  "top",
				Style: 1,
				Color: "808080",
			},
			{
				Type:  "left",
				Style: 1,
				Color: "808080",
			},
			{
				Type:  "right",
				Style: 1,
				Color: "808080",
			},
			{
				Type:  "bottom",
				Style: 1,
				Color: "808080",
			},
		},
	})
	if err != nil {
		log.Println(err)
		return 0
	}
	return sty
}

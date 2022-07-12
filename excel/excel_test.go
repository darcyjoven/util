package excel

import "testing"

func TestMakeExcel(t *testing.T) {
	f := NewExcel("demo.xlsx")
	f.AddSheet("test", [][]any{
		{"A1", "A2", "A3", "A4", "A5", "A6"},
		{12, 123, 64, 65, 134, 967},
		{12, 123, 64, 65, 134, 967},
		{12, 123, 64, 65, 134, 967},
		{12, 123, 64, 65, 134, 967},
		{12, 123, 64, 65, 134, 967},
		{12, 123, 64, 65, 134, 967},
		{12, 123, 64, 65, 134, 967},
	})
	f.Save()
}

package util

import (
	"github.com/russross/blackfriday/v2"
)

func MarkTotHtml(md []byte) []byte {
	a := blackfriday.Run(md,
		blackfriday.WithExtensions(blackfriday.Extensions(blackfriday.Table)),
		blackfriday.WithExtensions(blackfriday.Extensions(blackfriday.FencedCode)),
		blackfriday.WithExtensions(blackfriday.Extensions(blackfriday.CommonHTMLFlags)),
		blackfriday.WithExtensions(blackfriday.Extensions(blackfriday.CommonExtensions)),
	)
	return a
}

func NumToLetter(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每⼀位数据的值，然后通过索引的⽅式匹配A-Z
	)
	//⽤来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if Num > 26 { //数据⼤于26需要进⾏拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这⾥必须为A-Z中的⼀个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后⼀位数的值，因为已经记录在temp中
			if Num <= 26 {       //⼩于等于26直接进⾏匹配，不需要进⾏数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后⾯
	}
	return Str
}

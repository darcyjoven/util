package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {

	if GetShortDate() != fmt.Sprintf(
		"%04d%02d%02d",
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
	) {
		t.Error(GetLongDate())
	} else {
		fmt.Print("success")
	}

	fmt.Print(GetLongDate())

}

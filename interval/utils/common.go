package utils

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func AddError(top string, err error) {
	fmt.Sprintf("[%s]\t[%s]\t[err = %v]\n", NowStr(), top, err)
}

func NowStr() string {
	return time.Now().Format(timeFormat)
}

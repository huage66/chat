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

func Map2Strings(retMap map[string]bool) []string {
	retList := make([]string, len(retMap))
	for k, _ := range retMap {
		retList = append(retList, k)
	}
	return retList
}
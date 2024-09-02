package godemo

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func FromUnixTime() {
	unixTimeStr := "1677721627"
	tmpTime, err := strconv.ParseInt(unixTimeStr, 10, 64)

	if err != nil {
		log.Fatal(err)
	}
	unixTime := time.Unix(tmpTime, 0)
	dateTime := unixTime.Format("2006-01-02 15:04:05")
	fmt.Println(dateTime)
}

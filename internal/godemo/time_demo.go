package godemo

import (
	"encoding/json"
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

	now := time.Now()
	day, month, year := now.Date()
	fmt.Printf("%d%d%d", day, month, year)

	fmt.Println(now.Format("20060102"))

	newTime := "2023-01-20 19:80:80"
	tmp, _ := time.Parse("2006-01-02 15:04:05", newTime)
	fmt.Println(tmp)

	b := "\\"

	c, _ := json.Marshal(b)
	fmt.Println(b, string(c))

	fmt.Printf(b)
}

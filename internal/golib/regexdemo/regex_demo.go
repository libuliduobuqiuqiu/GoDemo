package regexdemo

import (
	"fmt"
	"log"
	"regexp"
)

func InterpreteDataGroup() {

	dgStr := "    if { [class match [IP::client_addr] eq NB_test ] } {"

	r, err := regexp.Compile(`.*?class\s+match.*?eq\s+(\S+)`)
	if err != nil {
		log.Fatal(err)
	}

	match := r.FindStringSubmatch(dgStr)
	if len(match) > 0 {
		fmt.Println(string(match[1]))
	}

}

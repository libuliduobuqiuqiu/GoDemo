package test

import (
	"fmt"
	"net"
	"testing"
)

type MainPerson struct {
	Name string `json:"name"`
}

func JudgeMainPerson(p interface{}) {
	if person, ok := p.(*MainPerson); ok {
		fmt.Println(person.Name)
		person.Name = "linsan"
	}
}

func JudgeList(notExistPerson []string) {
	fmt.Println(len(notExistPerson))
	if len(notExistPerson) > 0 {
		fmt.Println("Not Exist Person")
	}
	fmt.Println("done.")
}

func TestJudegeMainPerson(t *testing.T) {
	p := MainPerson{
		Name: "zhangsan",
	}
	JudgeMainPerson(&p)
	fmt.Println(p.Name)
}

func TestJudgeList(t *testing.T) {
	JudgeList(nil)
}

func TestNetParseIP(t *testing.T) {
	ip := "240e:6b1:10:1::46"
	tmpIP := net.ParseIP(ip)
	fmt.Println(tmpIP)
	fmt.Println(IPToBinary(tmpIP))
}

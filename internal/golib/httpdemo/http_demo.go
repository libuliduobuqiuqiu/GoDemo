package httpdemo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type routerManger struct{}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonList struct {
	Items []Person `json:"items,omitempty"`
	Code  int      `json:"code"`
	Msg   string   `json:"msg"`
}

var Persons PersonList

func (r *routerManger) HandelePerson(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		r.PostPersonInfo(w, req)
	} else if req.Method == "GET" {
		r.GetPersonInfo(w, req)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorContent := PersonList{Code: int(http.StatusNoContent), Msg: "Cannot support this method. "}
		if data, err := json.Marshal(errorContent); err == nil {
			fmt.Println(string(data))
			w.Write(data)
		}
	}
}

func (r *routerManger) GetPersonInfo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if data, err := json.Marshal(Persons); err == nil {
		w.Write(data)
	}
}

func (r *routerManger) PostPersonInfo(w http.ResponseWriter, req *http.Request) {
	var person Person
	data, _ := io.ReadAll(req.Body)
	w.Header().Set("Content-Type", "application/json")

	if data == nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorContent := PersonList{Code: int(http.StatusInternalServerError), Msg: "Body is not empty."}
		if errData, err := json.Marshal(errorContent); err == nil {
			fmt.Println(string(errData))
			w.Write(errData)
			return
		}
	}

	if err := json.Unmarshal(data, &person); err != nil {
		log.Fatal(err)
	}

	Persons.Items = append(Persons.Items, person)
	successContent := PersonList{Code: int(http.StatusOK), Msg: "success."}
	if successData, err := json.Marshal(successContent); err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Println(string(successData))
		w.Write(successData)
		return
	}
}

func HandleHttpRequest() {
	r := routerManger{}
	http.HandleFunc("/person", r.HandelePerson)
	log.Fatal(http.ListenAndServe("0.0.0.0:8989", nil))
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Response18nError struct {
	Context *gin.Context `json:"-"`
	Code    string       `json:"code"`
	Data    interface{}  `json:"data"`
	Msg     string       `json:"message"`
	Extra   interface{}  `json:"extra"`
}

func main() {
	test := &Response18nError{Context: &gin.Context{Accepted: []string{"AAA", "BBB", "CCC"}}, Code: "Code", Data: "Data", Msg: "Msg", Extra: "Extra"}
	testJSON, err := json.Marshal(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(testJSON))
}

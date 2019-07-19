package response

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSillyHatError(t *testing.T) {
	test1 := SillyHatError{Code: 1, Data: "data", Msg: "Msg", Extra: "Extra"}
	test1JSON, err := json.Marshal(test1)
	fmt.Println(string(test1JSON))
	fmt.Println(test1.Error())
	assert.Nil(t, err)
	assert.EqualValues(t, string(test1JSON), test1.Error())
	//[]interface{}
	fmt.Println(fmt.Sprintf("test %s %s", []string{"aaa", "bbb"}))
	//test2 := SillyHatError{Code: 1, Data: "data", Msg: "test message %s %t %d", Extra: "Extra",Args:}
	//test2JSON, err := json.Marshal(test1)
	//fmt.Println(string(test2JSON))
	//fmt.Println(test2.Error())
	//assert.Nil(t, err)
	//assert.EqualValues(t, string(test2JSON), test2.Error())
}

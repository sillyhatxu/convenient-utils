package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	type Department struct {
		ID   int    `validate:"Required"`
		Name string `validate:"Required;Length(20)"`
	}
	type User struct {
		ID                    int
		TestRequiredAndLength string `validate:"Required;Length(20)"`
		TestRangeRange        int64  `validate:"Range(20, 80)"`
		TestMax               int32  `validate:"Max(20)"`
		TestMin               int    `validate:"Min(20)"`
		TestMinMax            int    `validate:"Min(20);Max(80)"`
		Dept                  Department
		DeptRequired          Department `validate:"Required"`
	}
	u := &User{
		ID:                    1001,
		TestRequiredAndLength: "TestRequiredAndLength",
		TestRangeRange:        25,
		TestMax:               20,
		TestMin:               20,
		Dept:                  Department{ID: 1, Name: "DEPT"},
		DeptRequired:          Department{ID: 1, Name: "DEPT"},
	}
	e := Entry{}
	b, err := e.Validate(u)
	assert.Nil(t, err)
	assert.EqualValues(t, b, true)
}

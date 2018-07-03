package validator_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"git.ycyz.org/kolter/validator"
)

type obj1 struct {
	Account string `valid:"Required;"`
}

func TestValidator(t *testing.T) {
	Convey("", t, func() {
		temp := &obj1{Account: "admin"}
		v := validator.Validation{}
		v.Valid(temp)
	})
}

package validator

import (
	"reflect"
	"fmt"
)

const keyWord = "valid"

type Validation struct {
}

type Result struct {
	Message   string
	ErrorCode int
	Ok        bool
}

func (v *Validation) register(chk Validator, obj interface{}) *Result {
	if ok := chk.Valid(obj); ok {
		return &Result{Ok: true}
	}
	return &Result{}
}

func (v *Validation) Valid(obj interface{}) (bool, error) {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case isStruct(objT):
	case isPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return false, fmt.Errorf("%v must be struct or ptr", obj)
	}
	for i := 0; i < objT.NumField(); i++ {
		fmt.Println(objT.Field(i).Name)
	}
	return false, nil
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr
}


func (v *Validation)getFuncs(f reflect.StructField)(bool,funcs [])
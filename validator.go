package validator

type Validator interface {
	Valid(interface{}) bool
	Key() string
}

type Required struct {
	Key string
}

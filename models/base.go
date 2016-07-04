package models

type Envelope struct {
	Result bool
	ErrMsg string
	Data   interface{}
}

package main

import "fmt"

type App interface {
	// Hi is my first comment.
	// and this is second line.
	// we went to the third line :)
	// @myAnnotation `json:"a"`
	// @tx `retry:"4"`
	Hi(username string, password string) (sentence string, err error)
	By() (string, error)
	B
}

type B interface {
	// ByFromB says By from the B method :)
	// @trace
	ByFromB() (string, error)
}

type appImpl struct {
}

func (a *appImpl) ByFromB() (string, error) {
	return fmt.Sprint("by from b."), nil
}

func (a *appImpl) Hi(username string, password string) (sentence string, err error) {
	return fmt.Sprintf("hi %s", username), nil
}

func (a *appImpl) By() (string, error) {
	return "By", nil
}

var _ App = &appImpl{}

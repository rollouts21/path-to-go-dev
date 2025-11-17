package main

import (
	"errors"
	"fmt"
)

var ErrSomething = errors.New("something")

type SomethinError struct {
	code int
}

func (s SomethinError) Error() string {
	return fmt.Sprintf("Some error [%d]", s.code)
}

func (s SomethinError) Code() int {
	return s.code
}

type constSttringError string

func (e constSttringError) Error() string {
	return string(e)
}

const ErrConstSomething constSttringError = "some erorr"

type constIntError int

func (e constIntError) Error() string {
	return fmt.Sprintf("some error %d", e)
}

const (
	ErrFirst constIntError = iota

	SecondErr
)

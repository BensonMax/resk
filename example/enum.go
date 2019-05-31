package main

import "fmt"

type Status int

const (
	StatusOk      Status = 201
	StatusFailed  Status = 400
	StatusTimeout Status = 500
)

func main() {
	var s Status
	s = StatusFailed

	fmt.Println(s)
}

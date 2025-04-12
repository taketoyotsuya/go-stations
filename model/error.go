package model

import (
	"fmt"
	"time"
)

type ErrNotFound struct {
	When time.Time
	What string
}

func (e *ErrNotFound) Error() string {
	return "not found"
}

func run() error {
	return &ErrNotFound{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		//エラー検知
		fmt.Println(err)
	}
}

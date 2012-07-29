package main

import (
	"fmt"
	"os"
	"time"
	"math"
)

type MyError struct {
	When time.Time
	What string
}

func ( e *MyError ) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What ) 
}

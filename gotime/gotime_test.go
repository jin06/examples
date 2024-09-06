package gotime

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	time.Now()
	time.Now().Add(time.Second)
	time.Now().Format()
	time.Now().In()
	time.Now().Format()

	time.After()
	time.Since()
}

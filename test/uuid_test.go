package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestUUid(t *testing.T) {
	s := uuid.NewV4().String()
	fmt.Println(s)
}

package jg

import (
	"fmt"
	"testing"
)

func TestGetMob(t *testing.T) {
	p, err := GetMobile("", "")
	fmt.Println(p, err)
}

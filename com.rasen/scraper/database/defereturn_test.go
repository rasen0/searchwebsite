package database_test

import (
	"errors"
	"fmt"
	"testing"
)

func returnFlag()bool{
	fmt.Println("in return fn")
	defer func() bool{
		fmt.Println("in defer")
		return true
	}()
	panic(errors.New("panic"))
	return false
}

func TestDefereturn(t *testing.T)  {
	flag := returnFlag()
	fmt.Println("flag:",flag)
}

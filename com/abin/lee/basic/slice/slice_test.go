package slice

import (
	"fmt"
	"testing"
)

func Test_slice1(t *testing.T){
	slice1 := make([]string, 5)
	fmt.Println("slice1=",slice1)
}

func Test_slice2(t *testing.T){
	slice2 := make([]string, 5)
	slice2 = append(slice2, "one")
	fmt.Println("slice2=",slice2)
}








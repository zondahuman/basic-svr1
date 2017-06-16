package basic

import (
	"testing"
	"fmt"
)

func Test_array1(t *testing.T){
	var array1 [5]int
	fmt.Println("array1=",array1)
}
func Test_array2(t *testing.T){
	array2 := [5]int{1,5,4,2,3}
	fmt.Println("array2=",array2)
}

func Test_array3(t *testing.T){
	array3 := [...]int{9,4,7,5,6,8}
	fmt.Println("array3=",array3)
}

func Test_array4(t *testing.T){
	array4 := [...]int{1:5, 2:9, 5:12}
	fmt.Println("array4=",array4)
}
func Test_array5(t *testing.T){
	array5 := [6]int{1:5, 2:9, 5:12}
	fmt.Println("array5=",array5)
}
func Test_array6(t *testing.T){
	array6 := [5]*int{0:new(int), 1:new(int)}
	*array6[0] = 66
	*array6[1] = 77
	fmt.Println("array6=",array6)
}
func Test_array7(t *testing.T){
	var array7 [5]string
	array8 := [5]string{"abin","lee","stevenjohn","zondahuman","paul"}
	fmt.Println("array8=",array8)
	array7 = array8
	fmt.Println("array7=",array7)

}



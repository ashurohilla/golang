package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Println("hello i am there and lets code")

	// data types   are  3 type  in go   1 Numeric and sub parts are

	//1.1  integer types

	var a int = 342 //  size of  int is 8 bytes -128 to 127

	var b int8 = 127          //  size of int8 is 1 bytes -128 to 127
	var ba int16 = 32767      //  size of int16 is 2 bytes  -32,768 to 32,767
	var bb int32 = 2147483647 //  size of int32 is 4 bytes  -2,147,483,648 to 2,147,483,647

	var bc int64 = 9223372036854775807 //  size of int64  is 8 bytes -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

	fmt.Printf("a (int): %d\n", a)
	fmt.Printf("a (int): %d\n", b)
	fmt.Printf("a (int): %d\n", ba)
	fmt.Printf("a (int): %d\n", bb)
	fmt.Printf("a (int): %d\n", bc)

	//checking size using unsafe.sizeof

	fmt.Printf("Size of int: %d bytes\n", unsafe.Sizeof(int64(0)))

	// integers are also have types ..

	// 1.1.1  signed integers

	// signed inters are as int , int8 , int16 , int32 , int64

	//unsigned inters are as uint , uint8 , uint16 , uint32 , uint64, unintptr

	// 1.1.2  unsigned integers

	// 1.2-->   floating points

	// 1.2.1  float32 , float64

	var c float64 = 34.2

	fmt.Printf("a (float64): %f\n", c)

	// 1.3 -->   complex numbers

	var d complex128 = 1 + 2i

	fmt.Printf("c (complex128): %.1f + %.1fi\n", real(d), imag(d))

	// 1.3.1  complex64 , complex128

	// 2 bolean types

	var isGoFun bool = true

	fmt.Println(isGoFun)

	// 3 STRINGS

	var name string = "Golang"

	fmt.Println(name)

	// integer type

	// flaot type

}

package main

import (
	"fmt"
	"unsafe"
)

/* Size and types:
* string - 16 bytes
* int32 - 4 bytes
* int64 - 8 bytes
* int - depends on system architecture (32 bits - 4 bytes or 64 bits - 8 bytes)
* uint - equal rules for int type
* float32 - 4 bytes
* float64 - 8 bytes
* bool - 1 byte
 */

type StructA struct {
	i32 int32  // 4 bytes
	s   string // 16 bytes
	b   bool   // 1 bytes
}

type StructAOptimized struct {
	i32 int32  // 4 bytes
	b   bool   // 1 bytes
	s   string // 16 bytes
}

type StructB struct {
	i32 int32   // 4 bytes
	s   string  // 16 bytes
	f32 float32 // 4 bytes
}

type StructBOptimized struct {
	i32 int32   // 4 bytes
	f32 float32 // 4 bytes
	s   string  // 16 bytes
}

type StructC struct {
	b   bool  // 1 byte
	i64 int64 // 8 bytes
	i32 int32 // 4 bytes

}

type StructCOptimized struct {
	b   bool  // 1 byte
	i32 int32 // 4 bytes
	i64 int64 // 8 bytes
}

func main() {
	a := StructA{}
	aOptimized := StructAOptimized{}

	b := StructB{}
	bOptimized := StructBOptimized{}

	c := StructC{}
	cOptimized := StructCOptimized{}

	fmt.Println("Size of Struct A: ", unsafe.Sizeof(a))
	fmt.Println("Size of Struct A Optimized: ", unsafe.Sizeof(aOptimized))

	fmt.Println("Size of Struct B: ", unsafe.Sizeof(b))
	fmt.Println("Size of Struct B Optimized: ", unsafe.Sizeof(bOptimized))

	fmt.Println("Size of Struct C: ", unsafe.Sizeof(c))
	fmt.Println("Size of Struct C Optimized: ", unsafe.Sizeof(cOptimized))
}

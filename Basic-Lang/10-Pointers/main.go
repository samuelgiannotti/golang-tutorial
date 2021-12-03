package main

import "fmt"

func main() {
	basicPointer()
	basicPointer2()
	pointerNilSample()
	arrayOfPointers()
}

func basicPointer() {
	var a int = 10
	fmt.Printf("Address of a variable: %x\n", &a)
}

func basicPointer2() {
	fmt.Printf("basicPointer2 \n")
	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in pointer variable*/

	fmt.Printf("Address of a variable: %x\n", &a)
	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)
	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)
}

func pointerNilSample() {
	fmt.Printf("pointerNilSample \n")
	var ptr *int
	fmt.Printf("The value of ptr is : %x\n", ptr)
}

const MAX int = 3

func arrayOfPointers() {
	fmt.Printf("arrayOfPointers \n")
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer. */
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *ptr[i])
	}
}

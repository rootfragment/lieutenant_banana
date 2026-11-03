package main

import "fmt"

const Pi = 3.14159

func main() {
	var isReady bool = true
	fmt.Println("bool:", isReady)

	var a int = 42
	var b int8 = 127
	var c uint = 10
	var d int64 = 1 << 40
	var by byte = 65
	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("uint:", c)
	fmt.Println("int64:", d)
	fmt.Println("byte:", by)

	var pi float64 = 3.14159
	var f float32 = 2.5
	fmt.Println("float64:", pi)
	fmt.Println("float32:", f)

	var name string = "Gopher"
	fmt.Println("string:", name)
	fmt.Println("len(name):", len(name))
	fmt.Println("name + \"!\":", name+"!")

	var x int = 10
	var y = 20
	z := 30
	var count int
	fmt.Println("x:", x, "y:", y, "z:", z, "count:", count)

	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroPtr *int
	fmt.Println("int:", zeroInt)
	fmt.Println("float64:", zeroFloat)
	fmt.Println("bool:", zeroBool)
	fmt.Printf("string: %q (empty)\n", zeroString)
	fmt.Println("pointer:", zeroPtr)

	fmt.Println("Pi:", Pi)

	num := 10
	ptr := &num
	fmt.Println("before: num =", num, "| *ptr =", *ptr)
	*ptr = 20
	fmt.Println("after *ptr = 20: num =", num)

	gName := "Gopher"
	gAge := 5
	gHeight := 1.2
	isCute := true

	fmt.Printf("%s is %d years old, %.1fm tall, cute: %t\n", gName, gAge, gHeight, isCute)
	fmt.Printf("Types: %T %T %T %T\n", gName, gAge, gHeight, isCute)
}

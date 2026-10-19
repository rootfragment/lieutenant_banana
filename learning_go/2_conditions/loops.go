package main
import "fmt"
func main(){
	fmt.Println("classic for")
	for i :=0; i<5; i++{
		fmt.Println(i)
	}
	fmt.Println("while style")
	n := 1
	for n < 50{
		n *= 2}
	fmt.Println(n)
	fmt.Println("infinte loop + break")
	count := 0
	for {
		count++
		if count == 3 {
			fmt.Println("uve reached 3")
			break
		}
	}
	for i:= range 3{
		fmt.Println(i)
	}
}

package main 
import "fmt"
func main(){
	fmt.Println("Switch statement")
	d := "saturday"
	switch d{
	case "saturday", "sun":
		fmt.Println("weekend")
	case "fri", "mon", "tue", "wed", "thu":
		fmt.Println("weekday")

}}

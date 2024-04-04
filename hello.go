package main 
import "fmt"

// import "rsc.io/quote"

func main() {
	const greetingStart string = "Application started at port" //space automatically gets added before and after variable; this isn't JavaScript
	const greetingEnd string = "! Use Postman to test the API routes."
	const port int = 3001
	fmt.Println(greetingStart, port, greetingEnd)
}
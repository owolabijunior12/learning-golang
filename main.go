package main
import "fmt"

func main(){
	messagesFromUsers :=[]string{
		"hello world it is me iboytech",
		"currenting leraing go language",
		"didficult to learn go",
		"pleaase help me to learn go",
		"how to use go routines",
	}

	numMessages := float64(len(messagesFromUsers))
	costPerMessage := 0.02
	totalCost := numMessages * costPerMessage
	fmt.Printf("Total cost for processing %.2f messages: $%.2f\n", numMessages, totalCost)
	fmt.Println("hello world it is me iboytech")
}
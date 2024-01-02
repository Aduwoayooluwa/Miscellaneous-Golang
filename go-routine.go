package main
import (
	"fmt"
	 "time"
)


func main() {
	go test("When")
	test("Let's see")
	fmt.Println("run")
}

func test(s string) {
	for i:=0; i<5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
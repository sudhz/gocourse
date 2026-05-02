package basics

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

	if err != nil {
		fmt.Println("error", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("HTTP response status", resp.Status)
}

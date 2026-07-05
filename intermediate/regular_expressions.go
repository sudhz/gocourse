package intermediate

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("He said, \"I am great\"")
	fmt.Println(`He said, "I am great"`)

	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	email1 := "user@email.com"
	email2 := "invalid_email"

	fmt.Println("email1:", re.MatchString(email1))
	fmt.Println("email2:", re.MatchString(email2))

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2024-07-30"
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)

	str := "Hello World"
	re = regexp.MustCompile("[aeiou]")
	res := re.ReplaceAllString(str, "*")
	fmt.Println(res)

	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)
	text := "Golang is going great"
	fmt.Println("match:", re.MatchString(text))
}

package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!" // In go, \r moves to the left most position in the line

	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("length of rawMessage variable is", len(rawMessage))

	fmt.Println("the first char in message variable is", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"

	fmt.Println(greeting + name)

	str := "apple"   // a has an ASCII value of 97
	str1 := "Apple"  // A has an ASCII value of 65
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97

	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for i, char := range message {
		fmt.Printf("character at index %d is %c with ASCII value %v\n", i, char, char)
	}

	fmt.Println("rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	jch := '日'

	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)
	fmt.Printf("type of cstr is %T\n", cstr)

	const NIHONGO = "日本語"
	fmt.Println(NIHONGO)

	jhello := "こんにちは"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	emoji := '🙂'
	fmt.Printf("%v\n", emoji)
	fmt.Printf("%c\n", emoji)
}

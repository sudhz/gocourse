package intermediate

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	fmt.Println(str[0])

	fmt.Println(str[1:7])

	// Conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(str3)
	fmt.Println(len(str3))

	// Splitting
	fruits1 := "apple,orange,banana"
	fruits2 := "apple-orange-banana"
	parts1 := strings.Split(fruits1, ",")
	parts2 := strings.Split(fruits2, "-")
	fmt.Println(parts1)
	fmt.Println(parts2)

	// Joining
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	// Contains
	fmt.Println(strings.Contains(str, "Go"))
	fmt.Println(strings.Contains(str, "Go?"))

	// Replace
	replaced := strings.Replace(str, "Go", "Universe", 1)
	fmt.Println(replaced)

	// Trim
	strwithspaces := "    Hello Everyone!   "
	fmt.Println(strwithspaces)
	fmt.Println(strings.TrimSpace(strwithspaces))

	// Changing case
	fmt.Println(strings.ToLower(strwithspaces))
	fmt.Println(strings.ToUpper(strwithspaces))

	// Repeat
	fmt.Println(strings.Repeat("foo ", 3))

	// Count string in a string
	fmt.Println(strings.Count("Hello", "l"))
	fmt.Println(strings.HasPrefix("Hello", "he"))
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	// Regex
	str5 := "Hello, 123 Go! 11"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1)
	fmt.Println(matches)

	// UTF8
	str6 := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(str6))

	// String building
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	builderResult := builder.String()
	fmt.Println(builderResult)

	builder.WriteRune(' ')
	builder.WriteString("How are you")
	builderResult = builder.String()
	fmt.Println(builderResult)
	
	builder.Reset()
	builder.WriteString("Starting fresh!")

	builderResult = builder.String()
	fmt.Println(builderResult)

	 
}

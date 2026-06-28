package intermediate

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func main() {
	tmpl1, err := template.New("example").Parse("Welcome {{.name}}! How are you doing?\n")

	// No error handling required for this, it will automatically panic
	tmpl2 := template.Must(template.New("example").Parse("Hi {{.name}}! How are you doing?\n"))

	if err != nil {
		panic(err)
	}

	// Data for the welcome message template
	data := map[string]any{
		"name": "John",
	}

	err = tmpl1.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	tmpl2.Execute(os.Stdout, data)

	// Small console app
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We're glad you joined.",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Oops! An error occured: {{.errorMessage}}",
	}

	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu")
		fmt.Println("1. Join")
		fmt.Println("2. Get notification")
		fmt.Println("3. Get error")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]any{
				"name": name,
			}
		case "2":
			fmt.Print("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]any{
				"name": name,
				"notification": notification,
			}
		case "3":
			fmt.Print("Enter your error message: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["error"]
			data = map[string]any{
				"errorMessage": errorMessage,
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
			continue
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template:", err)
		}
	}
}

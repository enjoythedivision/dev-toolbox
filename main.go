package main
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to Dev Toolbox")
	text := "This is a simple tool to help developers with various tasks."
	fmt.Println(text)
	words := strings.Fields(text)
	fmt.Println("Words:", len(words))
}
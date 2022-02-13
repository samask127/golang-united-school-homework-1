package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	return emoji.Sprint("Hello :world_map:!")
}

func main() {
	var welcomeMessage = GetMessage()
	fmt.Println(welcomeMessage)
}

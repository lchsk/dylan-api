package main

import (
	"fmt"
)

func main() {
	songs := LoadData()

	fmt.Println(songs[0].Name)
}

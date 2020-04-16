package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	name := "/Users/suhun-uimac/Desktop/Card.txt"

	Content := read(name)
	remove(name)

	if len(Content) == 7 {
		fmt.Println(Content[3], " / ", Content[4], " / ", Content[5], " / ", Content[6])
	}

}

func read(name string) []string {
	var Content []string

	file, err := os.Open(name)
	if err != nil {
		return Content
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		Content = append(Content, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
	}
	return Content
}

func remove(name string) {
	err3 := os.Remove(name)
	if err3 != nil {
		fmt.Println("reading standard input:", err3)
	}
}

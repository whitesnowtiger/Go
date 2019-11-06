package main

import (
	"C:\\Users\\John\\Desktop\\Program\\Go\\LinkedList\\dataStruct"
	"fmt"
)

func main() {
	list := &dataStruct.LinkedList{}

	list.AddNode(0)

	for i := 1; i < 10; i++ {
		list.AddNode(i)
	}

	list.PrintNodes()
}

package customTypes

import "fmt"

type aliesStr string

func (text aliesStr) log() {
	fmt.Println(text)
}

func customTypes() {
	var name aliesStr = "Max"
	name.log()
}

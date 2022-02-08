package library

import "fmt"

func iniprivate() {
	fmt.Println("saya di private")
}

func Inipublic() {
	fmt.Println("saya di publik")
	iniprivate()
}

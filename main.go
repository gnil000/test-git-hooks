package main

import "fmt"

func main() {
	rabbit := &Rabbit{}

	rabbit.Poop()
}

type Rabbit struct {
}

func (r *Rabbit) Poop() {
	fmt.Println("pooped!")
}

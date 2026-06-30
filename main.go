package main

import "fmt"

func main() {
	rabbit := &Rabbit{}

	_ = rabbit.Poop(true)
}

type Rabbit struct {
}

func (r *Rabbit) Poop(success bool) error {
	if success {
		return nil
	}

	return fmt.Errorf("rabbit can't poop")
}

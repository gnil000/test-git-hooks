package main

import "fmt"

func main() {
	rabbit := &Rabbit{}

	_ = rabbit.Poop(false)

	fmt.Println("very long row. this row longer than 140 characters and go vet must return error for this. After need remove or cut this row. Yes complete.")
}

type Rabbit struct {
}

func (r *Rabbit) Poop(success bool) error {
	if success {
		return nil
	}

	return fmt.Errorf("rabbit can't poop")
}

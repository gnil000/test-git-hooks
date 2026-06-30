package main

import "testing"

func TestRabbitPoop(t *testing.T) {
	rabbit := &Rabbit{}

	err := rabbit.Poop(false)
	if err != nil {
		t.Fatal(err)
	}
}

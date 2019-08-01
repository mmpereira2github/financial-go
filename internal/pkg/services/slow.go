package services

import (
	"log"
	"time"
)

// SlowServiceInput is the input for slow service
type SlowServiceInput struct {
	Delay time.Duration
}

// SlowServiceOutput is the output for slow service
type SlowServiceOutput struct {
	Delay time.Duration
}

func init() {
	inputFactory := func() interface{} { return &SlowServiceInput{} }
	invoke := func(input interface{}) (interface{}, *Status) {
		if inputCasted, ok := input.(*SlowServiceInput); ok {
			return Slow(*inputCasted)
		}
		return Status{0, nil}, nil
	}
	Manager.Register(&ServiceEntry{
		ID:           "Slow",
		InputFactory: inputFactory,
		Invoke:       invoke,
	})
}

// Slow simulate a slow service
func Slow(input SlowServiceInput) (SlowServiceOutput, *Status) {
	log.Println("input", input)
	start := time.Now()
	if input.Delay != 0 {
		time.Sleep(input.Delay * time.Millisecond)
	} else {
		time.Sleep(100 * time.Millisecond)
	}
	end := time.Now()
	return SlowServiceOutput{end.Sub(start) / time.Millisecond}, nil
}

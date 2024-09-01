package gopromise

import (
	"log"
	"reflect"
	"sync"
)

type promises struct {
	fn     interface{}
	params interface{}
}

// execute is a function to execute promise
func (p promises) execute(wg *sync.WaitGroup, result chan []reflect.Value) {
	log.Println("Executing promise....")
	defer wg.Done()

	results, err := applyFunction(p.fn, p.params)
	if err != nil {
		log.Println(err)
	}

	result <- results
}

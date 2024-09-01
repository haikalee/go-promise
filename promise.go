package gopromise

import (
	"log"
	"reflect"
	"sync"
)

type Promises struct {
	fn     interface{}
	params interface{}
}

// execute is a function to execute promise
func (p Promises) execute(wg *sync.WaitGroup, result chan []reflect.Value) {
	log.Println("Executing promise....")
	defer wg.Done()

	results, err := applyFunction(p.fn, p.params)
	if err != nil {
		log.Println(err)
	}

	result <- results
}

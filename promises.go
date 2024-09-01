package gopromise

import (
	"errors"
	"log"
	"reflect"
	"sync"
)

type promise struct {
	promises []promises
}

// Promises is a struct to store function and parameters
func NewPromise() *promise {
	return &promise{
		promises: []promises{},
	}
}

// add is a function to add a function to promise
func (p *promise) Add(fn interface{}, params ...interface{}) error {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		return errors.New("fn is not a function")
	}

	p.promises = append(p.promises, promises{
		fn:     fn,
		params: params,
	})

	return nil
}

// Fetch is a function to fetch all data from promise
func (p *promise) Fetch() (datas [][]reflect.Value) {
	log.Println("Fetching data....")

	// create wait group for waiting after all promise action is completed
	wg := new(sync.WaitGroup)

	// create channel for check is action is complete or not
	done := make(chan bool)

	// create channel for get data from execute function
	data := make(chan []reflect.Value)

	go func() {
		// waiting for all executed action
		wg.Wait()

		// close data channel after all promise is executed
		close(data)

		// send channel to stop consuming channel
		done <- true
	}()

	// process execute action
	for _, promise := range p.promises {
		// add 1 wait group each for loop
		wg.Add(1)

		// execute process
		go promise.execute(wg, data)
	}

	// variable for consuming channel
	isContinue := true

	for isContinue {
		// process get data from channel
		select {
		case res, ok := <-data:
			// if data is empty, then stop consuming channel
			if !ok {
				isContinue = false
				break
			}

			// append result to container variable
			datas = append(datas, res)
		case <-done:
			isContinue = false
			close(done)
		}
	}

	return
}

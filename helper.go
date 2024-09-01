package gopromise

import (
	"errors"
	"reflect"
)

// applyFunction is a function to call a function with parameters
//
// fn is a function to be called
// params is a list of parameters
//
// return a list of reflect.Value
func applyFunction(fn interface{}, params interface{}) ([]reflect.Value, error) {
	/* get function */
	function := reflect.ValueOf(fn)

	/* check if fn is a function */
	if function.Kind() != reflect.Func {
		return nil, errors.New("fn is not a function")
	}

	/* check if fn has input parameters */
	args := []reflect.Value{}
	if reflect.TypeOf(fn).NumIn() > 0 {
		for _, param := range params.([]interface{}) {
			args = append(args, reflect.ValueOf(param))
		}
	}

	/* call function */
	results := function.Call(args)

	/* return result */
	return results, nil
}

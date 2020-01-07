package main

import "fmt"

type Any interface {}
type EvalFunc func(Any) (Any, Any)

func main()  {
	evenFunc := func(state Any) (Any, Any) {
		os := state.(int)
		ns := os + 2
		return os, ns
	}
	even := generator(evenFunc, 0)
	for i:=0; i<10; i++{
		fmt.Printf("%vth even: %v\n", i, even())
	}
}

func BuildLazyEvaluator(evalFunc EvalFunc, initstate Any) func() Any {
	retValChan := make(chan Any)
	loopFunc := func() {
		var actState = initstate
		var retVal Any
		for {
			retVal, actState = evalFunc(actState)
			retValChan <- retVal
		}
	}

	retFunc := func() Any{
		return <- retValChan
	}
	go loopFunc()
	return retFunc
}

func BuildLazyIntEvaluator(evalFunc EvalFunc, initState Any) func() int {
	ef := BuildLazyEvaluator(evalFunc, initState)
	return func() int {
		return ef().(int)
	}
}

func generator(evalFunc EvalFunc, initstate Any) func() Any {
	state := make(chan Any)
	go func() {state <- initstate}()
	return func() Any {
		var actState = <- state
		var retVal Any
		retVal, actState = evalFunc(actState)
		go func() {state <- actState}()
		return retVal
	}
}
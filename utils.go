package main

import (
	"fmt"
	"os"
	"reflect"
)

func FatalOnErr(err error) {
	if err != nil {
		fmt.Printf("Got Error: Exiting. %v\n", err.Error())
		os.Exit(1)
	}
}

// GetVals pulls the values out of a map into a slice of the
// appropriate type.
func GetVals(mapval, sliceptr interface{}) {
	mv := reflect.ValueOf(mapval)
	sv := reflect.ValueOf(sliceptr).Elem()

	for _, key := range mv.MapKeys() {
		sv.Set(reflect.Append(sv, mv.MapIndex(key)))
	}
}
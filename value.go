package sizer

import (
	"strconv"
	"strings"
)

type Value struct {
	value float64
}


func (v Value) String() string {
	return strconv.FormatFloat(v.value, 'f', -1, 64)
}

func ParseValue(b []byte) (Value, error) {

	var v Value
	var err error
	var s = strings.Replace(string(b), ",", ".", 1)
	v.value, err = strconv.ParseFloat(s, 64)
	return v, err
}

func (v Value) toBit() Value {
	return v.multiply(8)
}

func (v Value) toByte() Value {
	return v.divide(8)	
}

func (v Value) multiply(f float64) Value {
	v.value = v.value * f
	return v
}

func (v Value) divide(f float64) Value {
	v.value = v.value / f
	return v
}

func (v Value) Float() float64 {
	return v.value
}

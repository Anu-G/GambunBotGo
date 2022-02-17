package main

import (
	"reflect"
	"strings"
	"testing"
)

func Split() []string {
	var res []string
	splitter := strings.Split("$tl jpid text", " ")
	for i := range splitter {
		res = append(res, splitter[i])
	}

	return res
}

func Test_Split(t *testing.T) {
	exp := ""
	res := Split()

	if !reflect.DeepEqual(exp, res) {
		t.Errorf("res=%v", res[2])
	}
}

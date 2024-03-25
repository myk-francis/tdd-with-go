package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	// t.Run("initial test", func(t *testing.T) {
	// 	expected := "Chris"
	// 	var got []string

	// 	x := struct {
	// 		Name string
	// 	}{expected}

	// 	walk(x, func(input string) {
	// 		got = append(got, input)
	// 	})

	// 	if got[0] != expected {
	// 		t.Errorf("wrong number of function calls, got %q want %q", got[0], expected)
	// 	}
	// })

	cases := []struct {
		Name string
		Input interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
				Age int
			}{"Chris", 33},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
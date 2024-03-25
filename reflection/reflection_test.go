package main

import "testing"

func TestWalk(t *testing.T) {
	t.Run("initial test", func(t *testing.T) {
		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %q want %q", got[0], expected)
		}
	})
}
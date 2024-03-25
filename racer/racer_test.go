package main

import "testing"

// func TestInit(t *testing.T) {

// 	t.Run("initial test", func(t *testing.T) {

// 		want := fastUrl
// 		got := Racer(slowUrl, fastUrl)

// 		if got != want {
// 			t.Errorf("got %q want %q", got, want)
// 		}

// 	})
// }

func TestRacerOne(t *testing.T) {
	t.Run("initial test", func(t *testing.T) {
		slowUrl := "http://www.google.com"
		fastUrl := "http://www.facebook.com"

		want := fastUrl
		got := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestRacerTwo(t *testing.T) {
	t.Run("initial test", func(t *testing.T) {
		slowUrl := "http://www.google.com"
		fastUrl := "http://www.facebook.com"

		want := fastUrl
		got := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
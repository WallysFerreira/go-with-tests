package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "struct with non string field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 23},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "nested fields",
			Input: Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "using slices",
			Input: []Profile{
				{33, "London"},
				{34, "Los Angeles"},
			},
			ExpectedCalls: []string{"London", "Los Angeles"},
		},
		{
			Name: "using arrays",
			Input: [2]Profile{
				{33, "London"},
				{34, "Los Angeles"},
			},
			ExpectedCalls: []string{"London", "Los Angeles"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v expected %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("using maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow":   "Moo",
			"Sheep": "Baa",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Baa")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		expected := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{22, "Katowice"}
		}

		var got []string
		expected := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}

func assertContains(t testing.TB, got []string, expected string) {
	t.Helper()

	contains := false
	for _, x := range got {
		if x == expected {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", got, expected)
	}
}

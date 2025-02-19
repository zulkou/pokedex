package main

import (
    "fmt"
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: " hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "Pistachio  ",
            expected: []string{"pistachio"},
        },
        {
            input: "        24 CARAT lego",
            expected: []string{"24", "carat", "lego"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                err := fmt.Errorf("%v != %v", word, expectedWord)
                fmt.Println(err)
            }
        }
    } 
}

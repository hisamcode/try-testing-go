package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_alpha_isPrime(t *testing.T) {
	primeTest := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -1, false, "Negative number are not prime by definition!"},
	}

	for _, e := range primeTest {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func Test_alpha_prompt(t *testing.T) {
	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close out write
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of out prompt() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-> " {
		t.Errorf("incorrect propmt: expected -> but got %s", string(out))
	}
}

func Test_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	intro()

	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("intro text not correct; got %s", string(out))
	}
}

func Test_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty", "", "Please enter a whole number!"},
		{"zero", "0", "0 is not prime, by definition!"},
		{"one", "1", "1 is not prime, by definition!"},
		{"two", "2", "2 is a prime number!"},
		{"three", "3", "3 is a prime number!"},
		{"negative", "-1", "Negative number are not prime by definition!"},
		{"typed", "three", "Please enter a whole number!"},
		{"decimal", "1.1", "Please enter a whole number!"},
		{"quit", "q", ""},
		{"QUIT", "Q", ""},
		{"emot", "👍", "Please enter a whole number!"},
	}

	for _, e := range tests {
		input := strings.NewReader(e.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, e.expected) {
			t.Errorf("%s: expected %s, but got %s", e.name, e.expected, res)
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this function, we need a channel, and an instance of an io.Reader
	donech := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	// simulate
	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, donech)
	<-donech
	close(donech)
}

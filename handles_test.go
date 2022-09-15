package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"
)

func TestMutantPostRequestSuccess(t *testing.T) {

	jsonFile, err := os.Open("dnaOne.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened dnaOne.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonBody := []byte(byteValue)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonBody)))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	MutantPostRequest(w, r)

	if err != nil {
		log.Fatal(err)
	}
}

func TestMutantPostRequestFailure(t *testing.T) {

	jsonFile, err := os.Open("dnaThree.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened dnaThree.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonBody := []byte(byteValue)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonBody)))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	MutantPostRequest(w, r)

	if err != nil {
		log.Fatal(err)
	}
}

func TestMutantPostRequestFailureBody(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("asdf"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	MutantPostRequest(w, r)
}

func TestRooty(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	HandleRoot(w, r)
}

func TestMutantPostRequestFailureSizeBody(t *testing.T) {

	jsonFile, err := os.Open("dnaTwo.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened dnaTwo.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonBody := []byte(byteValue)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(jsonBody)))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	MutantPostRequest(w, r)
}

func TestReverse(t *testing.T) {
	test := Reverse("ASDF")

	if reflect.ValueOf(test).Kind() != reflect.String {
		t.Errorf("expected res to be and string")
	}
}

func TestTriangleAlgot(t *testing.T) {
	dna := DnaData{
		Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
	}

	test := TriangleAlgot(dna.Dna)

	if test {
		t.Errorf("expected res to be false")
	}
}

func TestEvalDiagonal(t *testing.T) {
	dna := DnaData{
		Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
	}
	var w http.ResponseWriter
	var wg sync.WaitGroup
	c := make(chan int)
	evalDiagonal(dna, w, &wg, c)

	if len(c) > 0 {
		t.Errorf("expected res channel to be empty")
	}
}

func TestEvalHorizontal(t *testing.T) {
	dna := DnaData{
		Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
	}
	var w http.ResponseWriter
	var wg sync.WaitGroup
	c := make(chan int)
	evalHorizontal(dna, w, &wg, c)

	if len(c) > 0 {
		t.Errorf("expected res channel to be empty")
	}
}

func TestEvalVertical(t *testing.T) {
	dna := DnaData{
		Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
	}
	var w http.ResponseWriter
	var wg sync.WaitGroup
	c := make(chan int)
	evalVertical(dna, w, &wg, c)

	if len(c) > 0 {
		t.Errorf("expected res channel to be empty")
	}
}

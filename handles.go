package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func TriangleAlgot(dna []string) bool {

	var evalPosition []string

	for prinIndex := 0; prinIndex < 6; prinIndex++ {

		evalPosition = []string{}

		for secIndex := 0; secIndex <= prinIndex; secIndex++ {
			if secIndex == 0 && len(evalPosition) == 0 {
				evalPosition = append(evalPosition, string(dna[secIndex][prinIndex-secIndex]))
			} else if len(evalPosition)-1 >= 0 && (evalPosition[len(evalPosition)-1] == string(dna[secIndex][prinIndex-secIndex])) {
				evalPosition = append(evalPosition, string(dna[secIndex][prinIndex-secIndex]))
			}
		}
		if len(evalPosition) >= 4 {
			return true
		}
	}
	return false
}

func evalDiagonal(dna DnaData, w http.ResponseWriter, wg *sync.WaitGroup, c chan int) {
	wg.Add(1)

	if w == nil || wg == nil || c == nil {
		return
	}

	var invertedArray []string
	mirrorArray := make([]string, 6)
	var invertedMirrorArray []string

	for position := 5; position > -1; position-- {
		invertedArray = append(invertedArray, Reverse(dna.Dna[position]))
	}

	for position := 0; position < 6; position++ {
		for indexL := 5; indexL > -1; indexL-- {
			mirrorArray[position] = mirrorArray[position] + string(dna.Dna[position][indexL])
		}
	}

	for position := 5; position > -1; position-- {
		invertedMirrorArray = append(invertedMirrorArray, Reverse(mirrorArray[position]))
	}

	evalDiagonal := TriangleAlgot(dna.Dna)
	evalInverted := TriangleAlgot(invertedArray)
	evalMirror := TriangleAlgot(mirrorArray)
	evalInvertedMirror := TriangleAlgot(invertedMirrorArray)

	if evalDiagonal || evalInverted || evalMirror || evalInvertedMirror {
		c <- 1
		return
	}

	c <- 0
	defer wg.Done()
}

func evalHorizontal(dna DnaData, w http.ResponseWriter, wg *sync.WaitGroup, c chan int) {
	wg.Add(1)

	if w == nil || wg == nil || c == nil {
		return
	}

	var evalHorizontal []string

	for _, v := range dna.Dna {

		var divide = strings.Split(v, "")

		for index, val := range divide {
			if index == 0 {
				evalHorizontal = []string{}
				evalHorizontal = append(evalHorizontal, val)
			} else if evalHorizontal[len(evalHorizontal)-1] == val {
				evalHorizontal = append(evalHorizontal, val)
			} else if evalHorizontal[len(evalHorizontal)-1] != val && len(evalHorizontal) < 4 {
				evalHorizontal = []string{val}
			}
		}

		if len(evalHorizontal) >= 4 {
			c <- 1
			return
		}
	}
	c <- 0
	defer wg.Done()
}

func evalVertical(dna DnaData, w http.ResponseWriter, wg *sync.WaitGroup, c chan int) {
	wg.Add(1)

	if w == nil || wg == nil || c == nil {
		return
	}
	var evalVertical []string

	for prinIndex := 0; prinIndex < len(dna.Dna); prinIndex++ {

		for secIndex := 0; secIndex < 6; secIndex++ {

			if len(evalVertical) == 0 {
				evalVertical = []string{}
				evalVertical = append(evalVertical, string(dna.Dna[secIndex][prinIndex]))
			} else if evalVertical[len(evalVertical)-1] == string(dna.Dna[secIndex][prinIndex]) {
				evalVertical = append(evalVertical, string(dna.Dna[secIndex][prinIndex]))
			} else if evalVertical[len(evalVertical)-1] != string(dna.Dna[secIndex][prinIndex]) && len(evalVertical) < 4 {
				evalVertical = []string{string(dna.Dna[secIndex][prinIndex])}
			}
		}

		if len(evalVertical) >= 4 {
			c <- 1
			return
		}
	}
	c <- 0
	defer wg.Done()
}

func MutantPostRequest(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	if w == nil || r == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var dna DnaData

	err := decoder.Decode(&dna)

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	if len(dna.Dna) != 6 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid size body error")
	}

	var wg sync.WaitGroup
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	go evalHorizontal(dna, w, &wg, c1)
	go evalVertical(dna, w, &wg, c2)
	go evalDiagonal(dna, w, &wg, c3)

	wg.Wait()

	mutantConfirm := <-c1 + <-c2 + <-c3

	if mutantConfirm >= 1 {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "true")
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "false")
		return
	}

}

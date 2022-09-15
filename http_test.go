package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

type Client struct {
	url string
}

func NewClient(url string) Client {
	return Client{url}
}

func (c *Client) MakeRequest() (string, error) {

	jsonFile, err := os.Open("dnaOne.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened dnaOne.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	jsonBody := []byte(byteValue)
	bodyReader := bytes.NewReader(jsonBody)

	client := &http.Client{}

	fmt.Println(c.url + "/mutant")

	res, _ := http.NewRequest("POST", c.url+"/mutant", bodyReader)
	res.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(res)

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	fmt.Println(string(b))

	return string(b), nil
}

func TestPostMutantApiSuccess(t *testing.T) {
	expected := "true"
	c := NewClient("http://localhost:3000")
	res, err := c.MakeRequest()
	if err != nil {
		t.Errorf("expected err to be nil got %v", err)
	}
	if res != expected {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}

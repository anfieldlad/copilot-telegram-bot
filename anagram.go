package main

// Method call this HTTP GET API http://www.anagramica.com/all/:parameter
// where :parameter is the word to be checked for anagrams.
// The API returns a JSON object with the following fields:
// all: all anagrams of the word
//
// Example:
// http://www.anagramica.com/all/cheese
// {
//  "all": [
//    "cheese"
//  ]
//}

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "fmt"
	"io/ioutil"
	_ "io/ioutil"
	"net/http"
	_ "net/http"
	"os"
	_ "os"
	"strings"
)

type Anagram struct {
	All []string `json:"all"`
}

// start function
func anagram(word string) string {
	url := "http://www.anagramica.com/all/" + word
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(responseData))
	var anagram Anagram
	err = json.Unmarshal(responseData, &anagram)
	if err != nil {
		return ""
	}

	return strings.Join(anagram.All, ", ")

}

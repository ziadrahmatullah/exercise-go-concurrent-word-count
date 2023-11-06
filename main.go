package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

var wordMap = map[string]int{
	"pellentesque" :0,
	"aliquet" :0,
	"at" : 0,
	"ut": 0,
	"nec": 0,
}
var mutex sync.Mutex

func calculateWords(input []string, expected string){
	mutex.Lock()
	for _, word := range input {
		if word == expected{
			wordMap[expected] += 1
		}
	}	
	mutex.Unlock()
}

func readInput() []string {
	content, _ := ioutil.ReadFile("file.txt")
	return strings.Fields(string(content))
}

func display(){
	fmt.Println("Word count result:")
	fmt.Println("====================")
	for word,  count:= range wordMap{
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main(){
	var wg sync.WaitGroup
	words := readInput()
	wg.Add(5)
	go func(){
		calculateWords(words, "pellentesque")
		wg.Done()
	}()
	go func(){
		calculateWords(words, "aliquet")
		wg.Done()
	}()
	go func(){
		calculateWords(words, "at")
		wg.Done()
	}()
	go func(){
		calculateWords(words, "ut")
		wg.Done()
	}()
	go func(){
		calculateWords(words, "nec")
		wg.Done()
	}()
	wg.Wait()
    display()
}

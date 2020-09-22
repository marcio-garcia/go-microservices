package main

import (
	"fmt"
	"sync"
)

func main() {
	inputChanel := make(chan string)
	outputChanel := make(chan string)
	defer close(outputChanel)

	var wg sync.WaitGroup
	go handleInput(&wg, inputChanel, outputChanel)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go sendLine(inputChanel, i)
	}

	wg.Wait()
	close(inputChanel)

	result := <-outputChanel

	fmt.Println("RESULT")
	fmt.Println(result)
}

func handleInput(wg *sync.WaitGroup, input chan string, output chan string) {

	var allLines string
	for incomingEvent := range input {
		fmt.Println(incomingEvent)
		printedLine := fmt.Sprintf("Printed: %s", incomingEvent)
		allLines = fmt.Sprintf("%s: %s - ", allLines, printedLine)
		wg.Done()
	}
	output <- allLines
}

func sendLine(input chan string, index int) {
	input <- fmt.Sprintf("LINE %v", index)
}

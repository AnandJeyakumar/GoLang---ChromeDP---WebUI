package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
)

func main() {

	headless := os.Getenv("HEADLESS") != "false"
	// / Set headless to false for headed mode
	// headless := false // Set to false for headed mode

	// Chrome options
	opts := chromedp.DefaultExecAllocatorOptions[:]
	opts = append(opts,
		// chromedp.ExecPath("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"), // Make sure this is your Chrome path
		chromedp.Flag("headless", headless),    // Toggle headless mode based on env
		chromedp.Flag("disable-gpu", headless), // Disable GPU in headless mode
		chromedp.Flag("no-sandbox", true),      // Disable sandbox
	)

	// Create Chrome allocator and context
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Set a timeout for the test
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Variables to hold test result
	var result string

	// Perform the UI automation tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.google.com"),
		chromedp.WaitVisible(`//textarea[@name="q"]`),
		chromedp.SendKeys(`//textarea[@name="q"]`, "Golang"),
		chromedp.SendKeys(`//textarea[@name="q"]`, kb.Enter),
		chromedp.WaitVisible(`#search`),
		chromedp.Text(`#search`, &result),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the result contains "Golang"
	if strings.Contains(result, "Golang") {
		fmt.Println("Test Passed: 'Golang' found in search results")
	} else {
		fmt.Println("Test Failed: 'Golang' not found in search results")
	}

	// Normal basic code
	// fmt.Println("Hey , This is the First code written in Go....")
	// Strings
	// var nameOne string = "Hey"
	// var nameTwo = "Welcome"
	// var nameThree string
	// nameThree = "To Gooo"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFOur := "Is a Gooo"

	// fmt.Println(nameFOur)

	// int

	// Strings
	// var age int = 90
	// var salary = 25000
	// // var year int

	// fmt.Println("The Salary is ", salary)

	// bits & memory
	// var numOne int8 = 25
	// var numTow int8 = -128
	// var numThree uint16 = 256

	// fmt.Println(numOne, numTow, numThree)

	// var ages [3]int = [3]int{90, 80, 85}
	// var ages = [3]int{90, 80, 85}

	// names := [3]string{"Hai", "Everything will be ", "good"}
	// fmt.Println(names, len(names))

	// fmt.Println(ages, len(ages))

	// slices under the hood
	// var scores = []int{90, 80, 85}
	// scores[1] = 222
	// scores = append(scores, 8888)

	// fmt.Println(scores, len(scores))

	// ranges

	// names := [4]string{"Hai", "Everything will be ", "good", "G"}
	// // fmt.Println(names, len(names))
	// rangeOne := names[1:3]

	// fmt.Println(rangeOne)

	// The standaert Library

	// greeting := "hello there all!!"

	// fmt.Println(strings.Contains(greeting, "all!!"))
	// fmt.Println(strings.ReplaceAll(greeting, "all!!", "Friends!!"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "th"))

}

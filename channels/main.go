package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for {
	//	go checkLink(<-c, c)
	//}

	// alternate way of doing the for loop above.
	for l := range c {
		go checkLink(l, c)
		// this is a function literal. parenthesis add the end of the function are to invoke the function.
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send message into channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

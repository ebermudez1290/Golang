package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	ch := make(chan string)

	for _, link := range sites {
		go checkLink(link, ch)
	}

	// for {
	// 	go checkLink(<-ch, ch)
	// }

	for l := range ch {
		//this will pause the main go routine. will block the other routines.
		// We share information by argument or communicate between ch , not directly to avoid confusion.
		go func(val string) {
			time.Sleep(time.Second)
			checkLink(val, ch)
		}(l) //This will invoque the function.
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		ch <- link
		return
	}
	fmt.Println(link, "is up")
	ch <- link
}

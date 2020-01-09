package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	//arbitrary slice since read does not resize daa
	bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	fmt.Println(string(bs))
	io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

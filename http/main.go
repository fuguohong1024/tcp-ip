package main

import "fmt"

func main() {
	reqBuf := buildHttpRequest()
	fmt.Printf("http request packet:\n%s", reqBuf.String())
	respBuf := buildHttpResponse()
	fmt.Printf("http response packet:\n%s", respBuf.String())
}

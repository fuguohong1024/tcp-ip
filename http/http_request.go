package main

import (
	"bytes"
	"unsafe"
)

const (
	helloWorld    = "hello world!"
	helloWorldLen = "12"
)

/*
http request packet:

HTTP Method | SP | URL
Header field name | : | Value | CR | LF
....| CR | LF
 CR | LF
HTTP body
*/

type Request struct {
	Method  string
	URL     string
	Version string
	Headers map[string]string
	Body    string
}

func buildHttpRequest() *bytes.Buffer {
	method := "GET"
	url := "/openapi/ping"
	version := "HTTP/1.1"
	header := make(map[string]string, 4)
	header["Host"] = "127.0.0.1"
	header["User-Agent"] = "Go-http-client/1.1"
	header["Content-Length"] = helloWorldLen
	header["Content-Type"] = "application/json;charset=UTF-8"
	body := helloWorld
	req := buildRequest(method, url, version, header, body)
	return req.BuildToBuffer()
}

func buildRequest(method, url, version string, headers map[string]string, body string) *Request {
	return &Request{
		Method:  method,
		URL:     url,
		Version: version,
		Headers: headers,
		Body:    body,
	}
}

// BuildToBuffer buffer write by copy not append
func (r *Request) BuildToBuffer() *bytes.Buffer {
	var req bytes.Buffer
	req.Write(Str2byte(r.Method))
	req.WriteByte(SP)
	req.Write(Str2byte(r.URL))
	req.WriteByte(SP)
	req.Write(Str2byte(r.Version))
	req.Write([]byte{CR, LF})
	for k, v := range r.Headers {
		req.Write(Str2byte(k))
		req.Write([]byte{COLON, SP})
		req.Write(Str2byte(v))
		req.Write([]byte{CR, LF})
	}
	req.Write([]byte{CR, LF})
	req.Write(Str2byte(r.Body))
	return &req
}

func Str2byte(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s
	// cap = len
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)
	return
}

func byte2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

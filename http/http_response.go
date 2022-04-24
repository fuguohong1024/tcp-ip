package main

import "bytes"

/*
http response packet:

HTTP Version | SP | Response Status Code | Response Phrase [可省略]
Header field name | : | Value | CR | LF
....| CR | LF
 CR | LF
HTTP body
*/

type Response struct {
	Version    string
	StatusCode string
	Phrase     string
	Headers    map[string]string
	Body       string
}

func buildHttpResponse() *bytes.Buffer {
	version := "HTTP/1.1"
	statusCode := "200"
	header := make(map[string]string, 4)
	header["Date"] = "Wed, 13 Apr 2022 10:41:11 GMT"
	header["Connection"] = "close"
	header["Content-Length"] = helloWorldLen
	header["Content-Type"] = "application/json;charset=UTF-8"
	body := helloWorld
	resp := buildResponse(version, statusCode, "", header, body)
	return resp.BuildToBuffer()
}

func buildResponse(version, statusCode, phrase string, headers map[string]string, body string) *Response {
	return &Response{
		Version:    version,
		StatusCode: statusCode,
		Phrase:     phrase,
		Headers:    headers,
		Body:       body,
	}
}

// BuildToBuffer buffer write by copy not append
func (r *Response) BuildToBuffer() *bytes.Buffer {
	var resp bytes.Buffer
	resp.Write(Str2byte(r.Version))
	resp.WriteByte(SP)
	resp.Write(Str2byte(r.StatusCode))
	if r.Phrase != "" {
		resp.WriteByte(SP)
		resp.Write(Str2byte(r.Phrase))
	}
	resp.Write([]byte{CR, LF})
	for k, v := range r.Headers {
		resp.Write(Str2byte(k))
		resp.Write([]byte{COLON, SP})
		resp.Write(Str2byte(v))
		resp.Write([]byte{CR, LF})
	}
	resp.Write([]byte{CR, LF})
	resp.Write(Str2byte(r.Body))
	return &resp
}

package client

import (
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"bytes"
	"go-http-client/entityMapper"
)

type CRequest struct {
	client *CClient
	request *http.Request
	uriBuilder *url.URL
	entityMapper entityMapper.EntityMapper
}


func(r *CRequest) Execute()(*CResponse, error) {
	request, error := r.httpRequest()
	if error != nil {
		return nil, error
	}
	response, error := r.client.execute(request)
	if error != nil {
		return nil, error
	}
	response.EntityMapper = r.client.entityMapper
	return response, nil
}

// this is just for test
func(r *CRequest) ExecuteForEntity(v interface{})(error) {
	response, error := r.Execute()
	if error != nil {
		return error
	}
	return response.ForEntity(&v)
}

func(r *CRequest) httpRequest()(*http.Request, error) {
	uri := r.uriBuilder
	r.request.URL = uri
	r.request.Host = uri.Host
	return r.request, nil
}

func(r *CRequest) Path(path ...string)(*CRequest)  {
	relativePath := strings.Join(path, "")
	newUrl := r.uriBuilder.String() + relativePath
	newURI, _ := url.Parse(newUrl)
	r.uriBuilder = newURI
	fmt.Println("Path", r.request.Method, r.uriBuilder)
	return r
}

func(r *CRequest) JsonEntity(param interface{}) (*CRequest)  {
	r.Header("Content-Type", "application/json")
	body, _ := json.Marshal(param)
	return r.Entity(bytes.NewReader(body))
}

func(r *CRequest) Entity(body io.Reader) (*CRequest) {
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	r.request.Body = rc
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			r.request.ContentLength = int64(v.Len())
		case *bytes.Reader:
			r.request.ContentLength = int64(v.Len())
		case *strings.Reader:
			r.request.ContentLength = int64(v.Len())
		}
	}
	return r
}

func(r *CRequest) Header(key string, value string) (*CRequest) {
	r.request.Header.Add(key, value)
	return r
}

func(r *CRequest) AddParam(key string, value string)(*CRequest) {
	q := r.uriBuilder.Query()
	q.Add(key, value)
	r.uriBuilder.RawQuery = q.Encode()
	return r
}






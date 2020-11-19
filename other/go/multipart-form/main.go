package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	go func() {
		http.Handle("/", &Server{})
		err := http.ListenAndServe(":8080", nil)
		checkErr(err)
	}()

	req, err := newRequest()
	checkErr(err)
	rsp, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer rsp.Body.Close()

	content, err := ioutil.ReadAll(rsp.Body)
	checkErr(err)
	fmt.Println(string(content))
}

func newRequest() (*http.Request, error) {
	requestBody := NewRequestBody()
	body, err := requestBody.GetRequestBody(map[string]string{}, map[string]string{"file": "./main.go"})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/", body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", requestBody.ContentType())
	return req, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

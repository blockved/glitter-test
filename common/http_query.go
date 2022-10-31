package common

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func DoPost(url, data string) (resp []byte, err error) {
	body := bytes.NewReader([]byte(data))
	contentType := "text/plain"
	r, err := http.Post(url, contentType, body)
	resp, _ = ioutil.ReadAll(r.Body)
	return
}

package xhttp

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/andybalholm/brotli"
)

var DefaultClient = &http.Client{}

func Get(url string, resp func(*http.Response) error) (err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add(UserAgent, UserAgentWindowsChrome93_0_4577_63)

	return Do(req, resp)
}

func Do(req *http.Request, resp func(*http.Response) error) (err error) {
	r, err := DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if resp == nil {
		return
	}

	return resp(r)
}

func JsonUnmarshalResponse(resp interface{}, body []byte) error {
	err := json.Unmarshal(body, resp)
	if err != nil {
		err = nil
		var reader io.Reader
		// gzip 解码
		reader, err := gzip.NewReader(bytes.NewReader(body))
		if err != nil {
			err = nil
			// brotli 解码
			reader = brotli.NewReader(bytes.NewReader(body))
		}
		body, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, resp)
		if err != nil {
			return err
		}
	}
	return nil
}

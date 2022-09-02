package xhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w}
}

func (r *Response) JSONWithStatus(statusCode int, body interface{}) error {
	r.Header().Add(ConContentType, ConContentTypeJson)
	r.WriteHeader(statusCode)

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = r.Write(b)
	return err
}

func (r *Response) JSONRespWithStatus(statusCode int, body *Resp) error {
	return r.JSONWithStatus(statusCode, body)
}

func (r *Response) JSON(body interface{}) error {
	return r.JSONWithStatus(http.StatusOK, body)
}

func (r *Response) JSONResp(body *Resp) error {
	return r.JSON(body)
}

func (r *Response) JSONP(callbackFunc string, body interface{}) error {
	r.Header().Add(ConContentType, ConContentTypeJavaScript)
	r.WriteHeader(http.StatusOK)

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = r.Write([]byte(fmt.Sprintf("%s(%s)", callbackFunc, string(b))))
	return err
}

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Err  string      `json:"err,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

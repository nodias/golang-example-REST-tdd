package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ResponseError struct {
	Err error `json:"err"`
}

func (r ResponseError) MarshalJSON() ([]byte, error) {
	if r.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, r.Error())), nil
}

func (r *ResponseError) UnmarshalJSON(b []byte) error {
	var v interface{}
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	if v == nil {
		r.Err = nil
		return nil
	}

	switch t := v.(type) {
	case string:
		r.Err = errors.New(t)
		return nil
	default:
		return errors.New("ResponseError unmarshal failed")
	}
}

func (r ResponseError) Error() string {
	if r.Err == nil {
		return ""
	}
	return r.Err.Error()
}

type Response struct {
	Products []product     `json:"products"`
	Err      ResponseError `json:"err"`
}

func (r Response) Error() string {
	return r.Err.Error()
}

package main

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		Products []product
		Err      ResponseError
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Products: []product{
					product{
						ID:    1,
						Name:  "tire",
						Price: 10,
					},
				},
				Err: ResponseError{},
			},
			want:    `{"products":[{"id":1,"name":"tire","price":10}],"err":null}`,
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				Products: nil,
				Err:      ResponseError{Err: errors.New("test error!")},
			},
			want:    `{"products":null,"err":"test error!"}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{
				Products: tt.fields.Products,
				Err:      tt.fields.Err,
			}
			got, err := json.Marshal(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}

func TestResponse_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name      string
		jsonValue string
		want      Response
		wantErr   bool
	}{
		{
			name:      "success",
			jsonValue: `{"products":[{"id":1,"name":"tire","price":10}],"err":null}`,
			want: Response{
				Products: []product{
					product{
						ID:    1,
						Name:  "tire",
						Price: 10,
					},
				},
				Err: ResponseError{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Response{}
			err := json.Unmarshal([]byte(tt.jsonValue), &r)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", r, tt.want)
			}
		})
	}
}

func TestResponseError_Error(t *testing.T) {
	type fields struct {
		Err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "success",
			fields: fields{Err: nil},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := ResponseError{
				Err: tt.fields.Err,
			}
			if got := r.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseError_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		val      string
		wantErr  bool
		wantRerr ResponseError
	}{
		{
			name:    "success",
			val:     `"test error!"`,
			wantErr: false,
			wantRerr:    ResponseError{Err: errors.New("test error!")},
		},
		{
			name:    "success : null",
			val:     "null",
			wantErr: false,
			wantRerr:    ResponseError{Err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rerr ResponseError
			err := rerr.UnmarshalJSON([]byte(tt.val));
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rerr, tt.wantRerr) {
				t.Errorf("UnmarshalJSON() rerr = %v expected = %v", tt.wantRerr, rerr)
			}
		})
	}
}

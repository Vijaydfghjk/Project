package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// WE are not using this package in in ths project

func Parsebody(r *http.Request, x interface{}) {

	/*

	 user will give some data in json format that we are unmarshal

	*/

	if body, err := io.ReadAll(r.Body); err == nil {

		if err := json.Unmarshal([]byte(body), x); err != nil {

			return
		}
	}
}

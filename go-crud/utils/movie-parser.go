package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParserequestToAnyDataModel(r *http.Request, x interface{}) {
	fmt.Printf("Inside parser for request %v", r)
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

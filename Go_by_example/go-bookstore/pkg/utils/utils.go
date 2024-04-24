//code for unmarshalling
package utils

import (
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func ParseBody(w *http.Request, x interface){
	if body, err := ioutil.ReadAll(r.body); err== nil{
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
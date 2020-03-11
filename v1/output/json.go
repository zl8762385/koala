package output

import (
	"encoding/json"
	"net/http"
)

type Json struct {

}

var jsonType = []string{"application/json; charset=utf-8"}

func (j Json) Content(rw http.ResponseWriter, Value interface{}) error {
	writeContentType(rw,jsonType)

	jsonBytes, err := json.Marshal(Value)
	if err != nil {
		return err
	}

	//rw.WriteHeader(200)
	rw.Write([]byte(jsonBytes))
	return nil
}

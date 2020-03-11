package output

import (
	"fmt"
	"net/http"
)

type Text struct {
}

func(t Text) Content(rw http.ResponseWriter, value interface{}) error {

	fmt.Fprint(rw, value)
	return nil
}

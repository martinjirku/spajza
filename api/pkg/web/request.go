package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func BindBody(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, &target)
	return nil
}

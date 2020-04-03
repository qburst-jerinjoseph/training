package app

import (
	"encoding/json"
	"net/http"
)

func (s *server) success(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(map[string]interface{}{"result": payload})
	if err != nil {
		s.Error(err)
		s.fail(w, http.StatusInternalServerError, "something went error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (s *server) fail(w http.ResponseWriter, code int, message string) {
	s.success(w, code, map[string]string{"error": message})
}

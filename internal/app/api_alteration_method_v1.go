package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (s *server) handleAlterationMethodList(w http.ResponseWriter, r *http.Request) {
	list, err := s.AlterationMethodList()
	if err != nil {
		s.Error("handleAlterationMethodList:", err)
		s.fail(w, http.StatusInternalServerError, "something went wrong")
	}
	s.success(w, http.StatusOK, list)
}

func (s *server) handleAlterationMethodGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		s.fail(w, http.StatusBadRequest, "invalid id")
		return
	}
	a, err := s.AlterationMethodGet(id)
	if err != nil {
		s.Error("AlterationMethodGet:", err)
		s.fail(w, http.StatusInternalServerError, "something went wrong")
	}
	if a == nil {
		s.fail(w, http.StatusNotFound, fmt.Sprintf("alteration method not found with id: %d", id))
		return
	}

	s.success(w, http.StatusOK, a)
}

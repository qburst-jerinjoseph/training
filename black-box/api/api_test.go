package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"training/internal/data"
)

func TestHandleAlterationMethodList(t *testing.T) {
	resp, err := http.Get("http://localhost:3005/training/v1/alteration-methods")
	if err != nil {
		t.Fatal("error while making request", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("error while reading response", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatal("Response code not matching expected:", http.StatusOK, "got:", resp.StatusCode)
	}
	data := map[string][]data.AlterationMethod{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Fatal("error while unmarshaling response", err)
	}
	if len(data["result"]) == 0 {
		t.Fatal("empty response")
	}
	for _, d := range data["result"] {
		if d.ID == 0 {
			t.Fatal("got empty id")
		}
		if d.DefaultName == "" {
			t.Fatal("got empty default_name")
		}
	}
}
func TestHandleAlterationMethodGet(t *testing.T) {
	request := func(t *testing.T, id string, code int) []byte {
		resp, err := http.Get("http://localhost:3005/training/v1/alteration-methods" + "/" + id)
		if err != nil {
			t.Fatal("error while making request", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error while reading response", err)
		}
		if resp.StatusCode != code {
			t.Fatal("Response code not matching expected:", code, "got:", resp.StatusCode)
		}
		return body
	}
	responseChecker := func(t *testing.T, body []byte) {
		data := map[string]data.AlterationMethod{}
		err := json.Unmarshal(body, &data)
		if err != nil {
			t.Fatal("error while unmarshaling response", err)
		}

		d := data["result"]
		if d.ID == 0 {
			t.Fatal("got empty id")
		}
		if d.DefaultName == "" {
			t.Fatal("got empty default_name")
		}

	}
	test := []struct {
		name string
		id   string
		code int
	}{
		{
			"success",
			"1",
			http.StatusOK,
		},
		{
			"invalid id",
			"aaa",
			http.StatusBadRequest,
		},
		{
			"invalid id",
			"100",
			http.StatusNotFound,
		},
	}
	for _, tc := range test {
		body := request(t, tc.id, tc.code)
		if tc.code == http.StatusOK {
			responseChecker(t, body)
		}
	}
}

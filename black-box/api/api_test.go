package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
	"training/internal/app"
	"training/internal/data"
)

func connectDB(dbName string) *sql.DB {
	connectionStr := "postgres://training:training@localhost:5432/%s?sslmode=disable"

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			connectionStr,
			dbName,
		),
	)
	if err != nil {
		log.Panicf(
			"error while connection to database:%s err:%v",
			dbName,
			err,
		)
	}
	return db
}
func TestMain(m *testing.M) {
	exitCode := 0
	testDBName := "training_test"
	defer func() {
		if r := recover(); r != nil {
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	adminDB := connectDB("postgres")

	_, err := adminDB.Exec("DROP DATABASE IF EXISTS " + testDBName)
	if err != nil {
		log.Panicf(
			"error while drop database:%s err:%v",
			testDBName,
			err,
		)
	}

	_, err = adminDB.Exec("CREATE DATABASE " + testDBName)
	if err != nil {
		log.Panicf(
			"error while creating database:%s err:%v",
			testDBName,
			err,
		)
	}
	testDB := connectDB(testDBName)

	defer func() {
		testDB.Close()

		qTerminateConnections := `
		SELECT
			pg_terminate_backend(pid)
		FROM
			pg_stat_activity
		WHERE
			datname = $1
	`
		_, err = adminDB.Exec(qTerminateConnections, testDBName)
		if err != nil && err != sql.ErrNoRows {
			log.Panicf(
				"failed to terminate active connections to test database %s, err: %v",
				testDBName,
				err,
			)
		}

		_, err := adminDB.Exec("DROP DATABASE " + testDBName)
		if err != nil {
			log.Panicf(
				"error while drop database:%s err:%v",
				testDBName,
				err,
			)
		}
		adminDB.Close()
	}()
	os.Setenv("DB_NAME", testDBName)
	os.Setenv("API_PORT", "3006")
	s := app.NewServer()
	go func() {
		s.Start()
	}()
	time.Sleep(10 * time.Millisecond)
	exitCode = m.Run()
}
func TestHandleAlterationMethodList(t *testing.T) {
	resp, err := http.Get("http://localhost:3006/training/v1/alteration-methods")
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
		resp, err := http.Get("http://localhost:3006/training/v1/alteration-methods" + "/" + id)
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

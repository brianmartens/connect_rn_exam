package main

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUsersHandle(t *testing.T) {
	var cases = []struct {
		input    string
		expected string
	}{
		{
			`{"data": [{"user_id": 2, "name": "Jane Doe", "date_of_birth": "1990-08-06", "created_on": 1642612034}]}`,
			`{"data": [{"user_id": 2, "name": "Jane Doe", "date_of_birth": "Monday", "created_on": "2022-01-19T12:07:14-05:00"}]}`,
		},
		{
			`{"data": [{"user_id": 1, "name": "Joe Smith", "date_of_birth": "1983-05-12", "created_on": 1642612034}]}`,
			`{"data": [{"user_id": 1, "name": "Joe Smith", "date_of_birth": "Thursday", "created_on": "2022-01-19T12:07:14-05:00"}]}`,
		},
	}

	e := echo.New()
	for _, test := range cases {
		body := bytes.NewReader([]byte(test.input))
		req := httptest.NewRequest("POST", "/users", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, UsersHandle(c)) {
			var left interface{}
			var right interface{}
			if err := json.Unmarshal([]byte(test.expected), &left); err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &right); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, left, right)
		}
	}
}

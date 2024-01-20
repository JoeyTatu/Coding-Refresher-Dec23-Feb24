package handlers

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostHandler() {
	c := http.Client{}

	r, err := c.Get("http://localhost:1323/posts")
	if err != nil {
		s.Fail(err.Error())
		return
	}
	defer r.Body.Close()

	s.Equal(http.StatusOK, r.StatusCode)
}

func (s *EndToEndSuite) TestPostNoResult() {
	c := http.Client{}

	r, err := c.Get("http://localhost:1323/post/5334")
	if err != nil {
		s.Fail(err.Error())
		return
	}
	defer r.Body.Close()

	s.Equal(http.StatusOK, r.StatusCode)
	b, err := io.ReadAll(r.Body)
	if err != nil {
		s.Fail(err.Error())
		return
	}

	s.JSONEq(`{"status": "ok", "data": []}`, string(b))
}

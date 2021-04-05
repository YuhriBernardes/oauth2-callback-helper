package server

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yuhribernardes/oauth2-callback-helper/testutils"
)

/* -------------------------------------------------------------------------- */
/*                            Test Instrumentation                            */
/* -------------------------------------------------------------------------- */

type OptionSuite struct {
	suite.Suite
	logger *log.Logger
	logs   *testutils.MemoryWriter
}

func (s *OptionSuite) SetupSuite() {
	s.logs = testutils.NewMemoryWriter()
	s.logger = log.New(s.logs, "", 0)
}

func (s *OptionSuite) SetupTest() {
	s.logs.Reset()
}

func (s *OptionSuite) GenerateHandlerInput(query map[string]string, headers map[string]string, body interface{}) (r *http.Request, w *httptest.ResponseRecorder) {
	uri := url.URL{
		Scheme: "http",
		Host:   "localhost",
		Path:   "/callback",
	}

	q := make(url.Values, 0)
	for k, v := range query {
		q.Set(k, v)
	}

	uri.RawQuery = q.Encode()

	var bodyReader io.Reader = nil
	if body != nil {
		binaryBody, _ := json.Marshal(body)
		bodyReader = bytes.NewReader(binaryBody)
	}

	s.T().Logf("Setting req uri to %s", uri.RequestURI())
	request := httptest.NewRequest(http.MethodGet, uri.String(), bodyReader)

	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}

	return request, httptest.NewRecorder()
}

/* -------------------------------------------------------------------------- */
/*                                Query Options                               */
/* -------------------------------------------------------------------------- */

type QueryOptionsSuite struct {
	OptionSuite
}

func TestQueryOptions(t *testing.T) {
	suite.Run(t, new(QueryOptionsSuite))
}

func (s *QueryOptionsSuite) TestOptionEnabled() {
	handler := CreateHandler(Options{
		Logger:    s.logger,
		ShowQuery: true,
	})

	query := map[string]string{
		"entry1": "value1",
		"entry2": "value2",
	}
	req, w := s.GenerateHandlerInput(query, nil, nil)

	handler.ServeHTTP(w, req)

	for _, v := range query {
		s.Containsf(s.logs.Data, v, "Expected table to contains %d, but it doesn't", v)
	}

}

func (s *QueryOptionsSuite) TestAllOptionsDisabled() {
	handler := CreateHandler(Options{
		Logger: s.logger,
	})

	query := map[string]string{
		"entry1": "value1",
		"entry2": "value2",
	}
	req, w := s.GenerateHandlerInput(query, nil, nil)

	handler.ServeHTTP(w, req)

	for _, v := range query {
		s.Containsf(s.logs.Data, v, "Expected table to contains %d, but it doesn't", v)
	}

}

func (s *QueryOptionsSuite) TestOnlyQueryOptionDisabled() {
	handler := CreateHandler(Options{
		Logger:   s.logger,
		ShowBody: true,
	})

	query := map[string]string{
		"entry1": "value1",
		"entry2": "value2",
	}
	req, w := s.GenerateHandlerInput(query, nil, nil)

	handler.ServeHTTP(w, req)

	for _, v := range query {
		s.NotContainsf(s.logs.Data, v, "Expected table to contains %d, but it doesn't", v)
	}
}

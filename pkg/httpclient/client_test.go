package httpclient

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
	client *HttpClient
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

func (s *ClientSuite) SetupTest() {
	s.client = New(time.Second * 1)
}

func (s *ClientSuite) TestVerifyURL() {
	testCases := []struct {
		desc        string
		URL         string
		expectError bool
	}{
		{
			desc:        "invalid_scheme",
			URL:         "httpz://example.com",
			expectError: true,
		},
		{
			desc:        "empty_scheme",
			URL:         "example.com",
			expectError: true,
		},
		{
			desc:        "valid_no_domain",
			URL:         "https://example/?",
			expectError: false,
		},
		{
			desc:        "valid",
			URL:         "https://example.com",
			expectError: false,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.desc, func(t *testing.T) {
			err := s.client.VerifyURL(tC.URL)
			s.Equal(tC.expectError, err != nil)
		})
	}
}

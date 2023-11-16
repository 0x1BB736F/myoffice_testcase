package filereader

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReaderSuite struct {
	suite.Suite

	reader   *Reader
	ioreader *bytes.Buffer
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(ReaderSuite))
}

func (s *ReaderSuite) SetupTest() {
	s.ioreader = &bytes.Buffer{}
	s.reader = New(s.ioreader)
}

func (s *ReaderSuite) TestEmptyReader() {
	// arrange
	result := ""
	// act
	c := s.reader.ReadChan()
	for result = range c {
	}
	<-c

	// assert
	s.Empty(result)
}

func (s *ReaderSuite) TestNotEmptyReader() {
	// arrange
	expected := "hello"
	s.ioreader.WriteString("hello")
	result := ""
	// act
	c := s.reader.ReadChan()
	for result = range c {
	}
	<-c

	// assert
	s.Equal(expected, result)
}

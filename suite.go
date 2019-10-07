package comodo

import (
	"github.com/stretchr/testify/suite"
)

// ComodoSuite is wrapper over suite.Suite with some addon methods
type ComodoSuite struct {
	suite.Suite
}

// Fatalf is equivalent to Logf followed by FailNow.
func (s *ComodoSuite) Fatalf(msg string, args ...interface{}) {
	s.T().Fatalf(msg, args...)
}

// Fatal is equivalent to Log followed by FailNow.
func (s *ComodoSuite) Fatal(args ...interface{}) {
	s.T().Fatal(args...)
}

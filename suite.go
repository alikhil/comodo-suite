// Package comodo is a wrapper over github.com/stretchr/testify/suite aimed to make it's use easier and comfortable
package comodo

import (
	"github.com/stretchr/testify/suite"
)

// Suite is wrapper over suite.Suite with some addon methods
type Suite struct {
	suite.Suite
}

// Fatalf is equivalent to Logf followed by FailNow.
func (s *Suite) Fatalf(msg string, args ...interface{}) {
	s.T().Fatalf(msg, args...)
}

// Fatal is equivalent to Log followed by FailNow.
func (s *Suite) Fatal(args ...interface{}) {
	s.T().Fatal(args...)
}

// NotEqualFail is equivalent to NotEqual followed by FailNow if assert fails.
func (s *Suite) NotEqualFail(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	var notEqual = s.Suite.NotEqual(expected, actual, msgAndArgs...)
	if !notEqual {
		s.Suite.T().FailNow()
	}
	return notEqual
}

// EqualFail is equivalent to Equal followed by FailNow if assert fails.
func (s *Suite) EqualFail(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	var equal = s.Suite.Equal(expected, actual, msgAndArgs...)
	if !equal {
		s.Suite.T().FailNow()
	}
	return equal
}

// NoErrorFail is equivalent to NoError followed by FailNow if assert fails.
func (s *Suite) NoErrorFail(err error, msgAndArgs ...interface{}) bool {
	var noError = s.Suite.NoError(err, msgAndArgs...)
	if !noError {
		s.Suite.T().FailNow()
	}
	return noError
}

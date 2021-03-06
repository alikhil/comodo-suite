// Package comodo is a wrapper over github.com/stretchr/testify/suite aimed to make it's use easier and comfortable
package comodo

import (
	"github.com/stretchr/testify/suite"
)

// Suite is wrapper over suite.Suite with some addon methods
type Suite struct {
	suite.Suite
}

func (s *Suite) fail() {
	s.Suite.T().FailNow()
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
		s.fail()
	}
	return notEqual
}

// EqualFail is equivalent to Equal followed by FailNow if assert fails.
func (s *Suite) EqualFail(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	var equal = s.Suite.Equal(expected, actual, msgAndArgs...)
	if !equal {
		s.fail()
	}
	return equal
}

// NoErrorFail is equivalent to NoError followed by FailNow if assert fails.
func (s *Suite) NoErrorFail(err error, msgAndArgs ...interface{}) bool {
	var noError = s.Suite.NoError(err, msgAndArgs...)
	if !noError {
		s.fail()
	}
	return noError
}

// TrueFail is equivalent to True followed by FailNow if assert fails
func (s *Suite) TrueFail(value bool, msgAndArgs ...interface{}) bool {
	var trueA = s.Suite.True(value, msgAndArgs...)
	if !trueA {
		s.fail()
	}
	return trueA
}

// FalseFail is equivalent to False followed by FailNow if assert fails
func (s *Suite) FalseFail(value bool, msgAndArgs ...interface{}) bool {
	var falseA = s.Suite.False(value, msgAndArgs...)
	if !falseA {
		s.fail()
	}
	return falseA
}

// NotNilFail is equivalent to NotNil followed by FailNow if assert fails
func (s *Suite) NotNilFail(value interface{}, msgAndArgs ...interface{}) bool {
	var notNil = s.Suite.NotNil(value, msgAndArgs...)
	if !notNil {
		s.fail()
	}
	return notNil
}

// NilFail is equivalent to Nil followed by FailNow if assert fails
func (s *Suite) NilFail(value interface{}, msgAndArgs ...interface{}) bool {
	var isNil = s.Suite.Nil(value, msgAndArgs...)
	if !isNil {
		s.fail()
	}
	return isNil
}

// EmptyFail is equivalent to Empty followed by FailNow if assert fails
func (s *Suite) EmptyFail(object interface{}, msgAndArgs ...interface{}) bool {
	var empty = s.Suite.Empty(object, msgAndArgs...)
	if !empty {
		s.fail()
	}
	return empty
}

// NotEmptyFail is equivalent to NotEmpty followed by FailNow if assert fails
func (s *Suite) NotEmptyFail(object interface{}, msgAndArgs ...interface{}) bool {
	var notEmpty = s.Suite.NotEmpty(object, msgAndArgs...)
	if !notEmpty {
		s.fail()
	}
	return notEmpty
}

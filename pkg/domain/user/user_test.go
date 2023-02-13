package user_test

import (
	"testing"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/user"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name          string
	ShouldFail    bool
	ExpectedError error
}

var TestCases = []TestCase{
	{
		Name:          "John",
		ShouldFail:    false,
		ExpectedError: nil,
	},
	{
		Name:          "diego armando maradona",
		ShouldFail:    true,
		ExpectedError: appError.ErrorInvalidArguments{},
	},
	{
		Name:          "he",
		ShouldFail:    true,
		ExpectedError: appError.ErrorInvalidArguments{},
	},
	{
		Name:          "Reza",
		ShouldFail:    false,
		ExpectedError: nil,
	},
}

func TestNewUser(t *testing.T) {
	t.Parallel()
	for _, testCase := range TestCases {
		user, err := user.NewUser(testCase.Name)
		if testCase.ShouldFail {
			assert.NotEqual(t, err, nil)
			assert.EqualError(
				t,
				err,
				"[user] invalid arguments: name length must be greater than 2 and less than 20",
			)
			assert.IsType(t, appError.ErrorInvalidArguments{}, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.Name, user.Name())
		}
	}
}

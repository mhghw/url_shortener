package shortUrl_test

import (
	"testing"
	"time"
	appError "urlShortener/error"
	"urlShortener/pkg/domain/shortUrl"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	OriginalUrl   string
	UserID        string
	CreatedAt     time.Time
	ShouldFail    bool
	ExpectedError error
}

var TestCases = []TestCase{
	{
		OriginalUrl:   "https://www.example.com/",
		UserID:        "PedSdefasdD",
		CreatedAt:     time.Now().UTC(),
		ShouldFail:    false,
		ExpectedError: nil,
	},
	{
		OriginalUrl:   "https:www.example.comsdfs/sdf",
		UserID:        "0gfdnfe37",
		CreatedAt:     time.Now().UTC(),
		ShouldFail:    false,
		ExpectedError: appError.ErrorInvalidArguments{},
	},
	{
		OriginalUrl:   "https://www.example.com/apparel",
		UserID:        "n6YDFBIDG",
		CreatedAt:     time.Now().UTC(),
		ShouldFail:    false,
		ExpectedError: nil,
	},
	{
		OriginalUrl:   "http://example.combrother#ThisSummer",
		UserID:        "IB25Twefasd",
		CreatedAt:     time.Now().UTC(),
		ShouldFail:    true,
		ExpectedError: appError.ErrorInvalidArguments{},
	},
	{
		OriginalUrl:   "http://www.example.com/apparel?books=basketball",
		UserID:        "TgnqwsdfpI",
		CreatedAt:     time.Now().UTC(),
		ShouldFail:    false,
		ExpectedError: nil,
	},
	{
		OriginalUrl:   "http://air.example.com/",
		UserID:        "g2iCxcvasdfrx",
		CreatedAt:     time.Now().UTC(),
		ShouldFail:    false,
		ExpectedError: nil,
	},
}

func TestNewShortUrl(t *testing.T) {
	t.Parallel()
	for _, testCase := range TestCases {
		st, err := shortUrl.NewShortUrl(testCase.OriginalUrl, testCase.UserID)
		if testCase.ShouldFail {
			assert.NotEqual(t, err, nil)
			assert.IsType(t, appError.ErrorInvalidArguments{}, err)

		} else {
			assert.Equal(t, testCase.OriginalUrl, st.OriginalUrl())
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.UserID, st.UserID())
		}
	}
}

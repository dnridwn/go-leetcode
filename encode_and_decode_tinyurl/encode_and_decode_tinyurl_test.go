// https://leetcode.com/problems/encode-and-decode-tinyurl/

package encode_and_decode_tinyurl

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeDecodeTinyUrl(t *testing.T) {
	codec := Constructor()

	testCases := []struct {
		name    string
		longUrl string
	}{
		{
			name:    "Test Encode Decode Tiny Url #1",
			longUrl: "https://google.com",
		},
		{
			name:    "Test Encode Decode Tiny Url #2",
			longUrl: "https://google.com/this/is/path",
		},
		{
			name:    "Test Encode Decode Tiny Url #3",
			longUrl: "https://google.com?query=params",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shortUrl := codec.encode(tc.longUrl)
			fmt.Println(shortUrl)
			assert.Equal(t, tc.longUrl, codec.decode(shortUrl))
		})
	}
}

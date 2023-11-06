// https://leetcode.com/problems/encode-and-decode-tinyurl/

package encode_and_decode_tinyurl

import "math/rand"

type Codec struct {
	defaultHostname   string
	encodedUrlMapping map[string]string
}

func Constructor() Codec {
	return Codec{
		defaultHostname:   "https://tinyurl.com/",
		encodedUrlMapping: map[string]string{},
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	shortUrl := c.defaultHostname + RandString(8)
	for ok := true; ok; ok = c.encodedUrlMapping[shortUrl] != "" {
		shortUrl = c.defaultHostname + RandString(8)
	}

	c.encodedUrlMapping[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	return c.encodedUrlMapping[shortUrl]
}

// Return random string with given nth length
func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

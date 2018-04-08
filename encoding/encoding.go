/*
Package encoding provides Decoding implementations.

Decode decodes HTTP responses:

	resp, _ := http.Get("http://api.example.com/")
	...
	var data map[string]interface{}
	err := JSONDecoder(resp.Body, &data)

*/
package encoding

import (
	"io"
)

// A Decoder is a function that reads from the reader and decodes it
// into an map of interfaces
type Decoder func(io.Reader, *map[string]interface{}) error

// A DecoderFactory is a function that returns CollectionDecoder or an EntityDecoder
type DecoderFactory func(bool) Decoder

// Deprecated: Register is deprecated
func Register(name string, dec DecoderFactory) error {
	return decoders.Register(name, dec)
}

// Deprecated: Get is deprecated
func Get(name string) DecoderFactory {
	return decoders.Get(name)
}

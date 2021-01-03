package option_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/lestrrat-go/option"
)

type identFoo struct{}
type identBar struct{}

func TestOption(t *testing.T) {
	options := []option.Interface{
		option.New(identFoo{}, "foo"),
		option.New(identBar{}, 1),
	}

	expected := []struct {
		ident interface{}
		value interface{}
	}{
		{
			ident: identFoo{},
			value: "foo",
		},
		{
			ident: identBar{},
			value: 1,
		},
	}

	for i := 0; i < len(options); i++ {
		if !assert.Equal(t, expected[i].ident, options[i].Ident(), `identities should match`) {
			return
		}
		if !assert.Equal(t, expected[i].value, options[i].Value(), `values should match`) {
			return
		}
	}
}

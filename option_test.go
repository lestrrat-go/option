package option_test

import (
	"testing"

	"github.com/lestrrat-go/option/v2"
	"github.com/stretchr/testify/require"
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
		require.Equal(t, expected[i].ident, options[i].Ident(), `identities should match`)

		var value interface{}
		require.NoError(t, options[i].Value(&value), `should be able to retrieve value`)
		require.Equal(t, expected[i].value, value, `values should match`)
	}
}

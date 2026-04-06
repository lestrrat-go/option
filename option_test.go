package option_test

import (
	"testing"

	"github.com/lestrrat-go/option/v3"
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
		ident any
		value any
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
	}

	// Test typed Value() on concrete *Option[T]
	strOpt := option.New(identFoo{}, "foo")
	require.Equal(t, "foo", strOpt.Value())

	intOpt := option.New(identBar{}, 42)
	require.Equal(t, 42, intOpt.Value())
}

func TestGet(t *testing.T) {
	opt := option.New(identFoo{}, "hello")

	// Successful extraction
	v, ok := option.Get[string](opt)
	require.True(t, ok)
	require.Equal(t, "hello", v)

	// Type mismatch
	_, ok = option.Get[int](opt)
	require.False(t, ok)
}

func TestMustGet(t *testing.T) {
	opt := option.New(identFoo{}, "hello")

	// Successful extraction
	v := option.MustGet[string](opt)
	require.Equal(t, "hello", v)

	// Panics on type mismatch
	require.Panics(t, func() {
		option.MustGet[int](opt)
	})
}

func TestGetThroughWrapper(t *testing.T) {
	// Simulate jwx's wrapper pattern: a struct embedding option.Interface
	type wrappedOption struct {
		option.Interface
	}

	inner := option.New(identFoo{}, "wrapped-value")
	wrapped := &wrappedOption{Interface: inner}

	// Get should work through the wrapper
	v, ok := option.Get[string](wrapped)
	require.True(t, ok)
	require.Equal(t, "wrapped-value", v)

	// Ident should also work through the wrapper
	require.Equal(t, identFoo{}, wrapped.Ident())
}

func TestGetWithPointerValue(t *testing.T) {
	type payload struct {
		Name string
		Age  int
	}

	p := &payload{Name: "test", Age: 30}
	opt := option.New(identFoo{}, p)

	v, ok := option.Get[*payload](opt)
	require.True(t, ok)
	require.Equal(t, "test", v.Name)
	require.Equal(t, 30, v.Age)
}

func BenchmarkGet(b *testing.B) {
	opt := option.New(identFoo{}, "hello")
	var iface option.Interface = opt

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _ := option.Get[string](iface)
		_ = v
	}
}

func BenchmarkDirectValue(b *testing.B) {
	opt := option.New(identFoo{}, "hello")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := opt.Value()
		_ = v
	}
}

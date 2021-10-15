package reflectx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseTag(t *testing.T) {
	type _model struct {
		ID int `tag:"message(this is tag meta...)"`
	}
	type _tag struct {
		Message string `alias:"message"`
	}
	_, _, tags := ParseTag(&_model{}, &_tag{}, "alias", "tag", false)
	require.Equal(t, "this is tag meta...", tags[0].(*_tag).Message)
}

func TestParseTagWithRe(t *testing.T) {
	type _model struct {
		ID int `tag:"message(this is tag meta...)"`
	}
	type _tag struct {
		Message string `alias:"message"`
	}
	_, _, tags := ParseTagWithRe(&_model{}, &_tag{}, "alias", "tag", false, `([a-zA-Z0-9_]+)\(([^()]+)\)`)
	require.Equal(t, "this is tag meta...", tags[0].(*_tag).Message)
}

package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/jumpsiegel/radiance/pkg/sbf"
)

func TestSymbolHash_Entrypoint(t *testing.T) {
	assert.Equal(t, sbf.EntrypointHash, sbf.SymbolHash("entrypoint"))
}

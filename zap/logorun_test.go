package adapters

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestNewLoggerLevel(t *testing.T) {
	l := Log{}.New()
	j := Log{}.New()
	assert.NotEqual(t, l, nil)
	assert.Equal(t, &l, &j)
	assert.Equal(t, l.GetLevel(), zap.DebugLevel.String())
}

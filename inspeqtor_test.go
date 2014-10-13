package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInspeqtorParse(t *testing.T) {
	t.Parallel()
	i, err := New("test", "")
	assert.Nil(t, err)
	err = i.Parse()
	assert.Nil(t, err)
	assert.Equal(t, uint16(15), i.GlobalConfig.Top.CycleTime)

	i, err = New("testx", "")
	assert.NotNil(t, i)
	assert.Nil(t, err)
	err = i.Parse()
	assert.NotNil(t, err)
	assert.True(t, strings.Index(err.Error(), "Missing required file") >= 0)
}

func TestCreateSocket(t *testing.T) {
	// not parallelizable since it uses a Unix sock
	i, err := New("test", "")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	err = i.openSocket("/tmp/tmp.sock")
	assert.Nil(t, err)
	assert.NotNil(t, i.Socket)
	defer i.Socket.Close()
}

func TestTestAlertRoutes(t *testing.T) {
	i, err := New("test", "")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	badCount := i.TestAlertRoutes()
	assert.Equal(t, badCount, 1)
}

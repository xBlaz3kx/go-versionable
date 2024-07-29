package go_versionable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVersionList_Add_WithLimit(t *testing.T) {
	list := NewVersionList[string](WithLimit[string](2))
	currentTime := time.Now()
	err := list.Add("test")
	assert.NoError(t, err)

	err = list.Add("test2")
	assert.NoError(t, err)

	err = list.Add("test3")
	assert.Error(t, err)

	versionList := list.GetAll()
	expectedList := []Version[string]{{Version: 1, Data: "test"}, {Version: 2, Data: "test2"}}
	for i, version := range versionList {
		assert.EqualValues(t, expectedList[i].Data, version.Data)
		assert.EqualValues(t, expectedList[i].Version, version.Version)
		assert.InDelta(t, currentTime.Second(), version.InsertedAt.Second(), 1)
	}
}

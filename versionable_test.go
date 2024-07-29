package go_versionable

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVersionList_Add(t *testing.T) {
	list := NewVersionList[string]()

	currentTime := time.Now()
	err := list.Add("test")
	assert.NoError(t, err)

	err = list.Add("test2")
	assert.NoError(t, err)

	err = list.Add("test3")
	assert.NoError(t, err)

	versionList := list.GetAll()
	expectedList := []Version[string]{{Version: 1, Data: "test", InsertedAt: currentTime}, {Version: 2, Data: "test2", InsertedAt: currentTime}, {Version: 3, Data: "test3", InsertedAt: currentTime}}
	for i, version := range versionList {
		assert.EqualValues(t, expectedList[i].Data, version.Data)
		assert.EqualValues(t, expectedList[i].Version, version.Version)
		assert.InDelta(t, currentTime.Second(), version.InsertedAt.Second(), 1)
	}
}

func TestVersionList_Get(t *testing.T) {
	nodeList := []Version[string]{{Version: 1, Data: "test"}, {Version: 2, Data: "test1"}, {Version: 3, Data: "test2"}}
	list := NewFromVersions[string](nodeList)

	data, isFound := list.Get(1)
	assert.True(t, isFound)
	assert.EqualValues(t, "test", *data)

	data, isFound = list.Get(2)
	assert.True(t, isFound)
	assert.EqualValues(t, "test1", *data)

	_, isFound = list.Get(123)
	assert.False(t, isFound)
}

func TestVersionList_GetLatest(t *testing.T) {
	nodeList := []Version[string]{{Version: 1, Data: "test"}, {Version: 2, Data: "test1"}, {Version: 3, Data: "test2"}}
	list := NewFromVersions[string](nodeList)

	latest, err := list.GetLatest()
	assert.NoError(t, err)
	assert.EqualValues(t, "test2", latest.Data)
	assert.EqualValues(t, 3, latest.Version)

	list = NewVersionList[string]()
	_, err = list.GetLatest()
	assert.Error(t, err)
	assert.EqualValues(t, ErrNoVersions, err)
}

func TestVersionList_GetVersions(t *testing.T) {
	nodeList := []Version[string]{{Version: 1, Data: "test"}, {Version: 2, Data: "test1"}, {Version: 3, Data: "test2"}}
	list := NewFromVersions[string](nodeList)

	versions := list.GetAll()
	assert.EqualValues(t, nodeList, versions)
}

func TestVersionList_Remove(t *testing.T) {
	nodeList := []Version[string]{{Version: 1, Data: "test"}, {Version: 2, Data: "test1"}, {Version: 3, Data: "test2"}}
	list := NewFromVersions[string](nodeList)

	err := list.Remove(2)
	assert.NoError(t, err)

	_, isFound := list.Get(2)
	assert.False(t, isFound)

	err = list.Remove(123)
	assert.Error(t, err)
}

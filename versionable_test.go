package go_versionable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionList_Add(t *testing.T) {
	list := NewVersionList[string]()

	err := list.Add("test")
	if err != nil {
		t.Error("Expected nil, got error")
	}

	err = list.Add("test2")
	if err != nil {
		t.Error("Expected nil, got error")
	}

	err = list.Add("test3")
	if err != nil {
		t.Error("Expected nil, got error")
	}

	versionList := list.GetAll()
	expectedList := []Version[string]{{1, "test"}, {2, "test2"}, {3, "test3"}}
	if !reflect.DeepEqual(list.GetAll(), expectedList) {
		t.Errorf("Expected %v, got %v", expectedList, versionList)
	}
}

func TestVersionList_Get(t *testing.T) {
	nodeList := []Version[string]{{1, "test"}, {2, "test1"}, {3, "test2"}}
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
	nodeList := []Version[string]{{1, "test"}, {2, "test1"}, {3, "test2"}}
	list := NewFromVersions[string](nodeList)

	latest, err := list.GetLatest()
	assert.NoError(t, err)
	assert.EqualValues(t, "test2", latest.Data)

	list = NewVersionList[string]()
	_, err = list.GetLatest()
	assert.Error(t, err)
	assert.EqualValues(t, ErrNoVersions, err)
}

func TestVersionList_GetVersions(t *testing.T) {
	nodeList := []Version[string]{{1, "test"}, {2, "test1"}, {3, "test2"}}
	list := NewFromVersions[string](nodeList)

	versions := list.GetAll()
	assert.EqualValues(t, nodeList, versions)
}

func TestVersionList_Remove(t *testing.T) {
	nodeList := []Version[string]{{1, "test"}, {2, "test1"}, {3, "test2"}}
	list := NewFromVersions[string](nodeList)

	err := list.Remove(2)
	assert.NoError(t, err)

	_, isFound := list.Get(2)
	assert.False(t, isFound)

	err = list.Remove(123)
	assert.Error(t, err)
}

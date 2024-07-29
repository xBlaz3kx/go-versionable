package go_versionable

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionList_Add_WithLimit(t *testing.T) {
	list := NewVersionList[string](WithLimit[string](2))

	err := list.Add("test")
	assert.NoError(t, err)

	err = list.Add("test2")
	assert.NoError(t, err)

	err = list.Add("test3")
	assert.Error(t, err)

	versionList := list.GetAll()
	expectedList := []Version[string]{{1, "test"}, {2, "test2"}}
	if !reflect.DeepEqual(list.GetAll(), expectedList) {
		t.Errorf("Expected %v, got %v", expectedList, versionList)
	}
}

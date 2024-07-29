package go_versionable

import (
	"errors"

	"github.com/emirpasic/gods/sets/linkedhashset"
)

var (
	ErrLimitReached    = errors.New("limit reached")
	ErrVersionExists   = errors.New("version already exists")
	ErrVersionNotFound = errors.New("version not found")
	ErrNoVersions      = errors.New("no versions found")
)

type Version[Data comparable] struct {
	Version VersionNum `json:"version"`
	Data    Data       `json:"data"`
}

type VersionNum int

func (n VersionNum) Equals(version int) bool {
	return int(n) == version
}

// VersionList is a linked list of versions
type VersionList[Data comparable] struct {
	// Limit of total versions stored in the list. Optional.
	limit *int

	// Hash set for quick access to set
	set *linkedhashset.Set
}

// NewFromVersions creates a new VersionList from a list of set (useful when deserializing from a database)
func NewFromVersions[Data comparable](versions []Version[Data], opts ...Option[Data]) *VersionList[Data] {
	set := linkedhashset.New()

	for _, version := range versions {
		// todo validate the version has a version
		set.Add(version)
	}

	versionList := &VersionList[Data]{
		set: set,
	}

	versionList.applyOpts(opts...)

	return versionList
}

// NewVersionList creates a new VersionList with the provided options
func NewVersionList[Data comparable](opts ...Option[Data]) *VersionList[Data] {
	versionList := &VersionList[Data]{
		set:   linkedhashset.New(),
		limit: nil,
	}

	versionList.applyOpts(opts...)

	return versionList
}

// applyOpts applies the provided options to the VersionList
func (vl *VersionList[Data]) applyOpts(opts ...Option[Data]) {
	for _, opt := range opts {
		opt(vl)
	}
}

// Add adds a new version of data.
func (vl *VersionList[Data]) Add(data Data) error {
	// Check if the limit is reached
	if vl.limit != nil && vl.set.Size() >= *vl.limit {
		return ErrLimitReached
	}

	version := vl.set.Size() + 1
	node := Version[Data]{
		Version: VersionNum(version),
		Data:    data,
	}

	// Check if the version already exists
	hasVersion := vl.set.Any(func(index int, value interface{}) bool {
		return value.(Version[Data]).Version.Equals(version)
	})
	if hasVersion {
		return ErrVersionExists
	}

	// Add the version to the set
	vl.set.Add(node)
	return nil
}

// GetLatest returns the latest version of the data.
func (vl *VersionList[Data]) GetLatest() (*Version[Data], error) {
	latestIndex := vl.set.Size() - 1
	if latestIndex < 0 {
		return nil, ErrNoVersions
	}

	latestNode := vl.set.Values()[latestIndex].(Version[Data])

	return &latestNode, nil
}

// Get returns a Version from the list
func (vl *VersionList[Data]) Get(version int) (*Data, bool) {
	isFound, node := vl.set.Find(func(index int, value interface{}) bool {
		val, canCast := value.(Version[Data])
		if !canCast {
			return false
		}
		return val.Version.Equals(version)
	})

	if isFound == -1 {
		return nil, false
	}

	data := node.(Version[Data]).Data
	return &data, true
}

// Remove deletes a Version from the list
func (vl *VersionList[Data]) Remove(version int) error {
	var err = ErrVersionNotFound

	vl.set.Each(func(index int, value interface{}) {
		if value.(Version[Data]).Version.Equals(version) {
			vl.set.Remove(value)
			err = nil
		}
	})

	return err
}

// GetAll returns all versions of data.
func (vl *VersionList[Data]) GetAll() []Version[Data] {
	values := []Version[Data]{}

	vl.set.Each(func(index int, value interface{}) {
		values = append(values, value.(Version[Data]))
	})

	return values
}

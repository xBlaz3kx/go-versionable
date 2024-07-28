package go_versionable

// Option is a functional option for VersionList
type Option[Data comparable] func(*VersionList[Data])

// WithLimit sets the limit of total versions stored in the list
func WithLimit(limit int) func(*VersionList[any]) {
	return func(vl *VersionList[any]) {
		vl.limit = &limit
	}
}

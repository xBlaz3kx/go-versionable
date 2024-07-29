package go_versionable

// Option is a functional option for VersionList
type Option[Data comparable] func(*VersionList[Data])

// WithLimit sets the limit of total versions stored in the list
func WithLimit[Data comparable](limit int) func(*VersionList[Data]) {
	return func(vl *VersionList[Data]) {
		vl.limit = &limit
	}
}

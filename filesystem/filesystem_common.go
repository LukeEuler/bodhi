package filesystem

// FileSystem ...
type FileSystem struct{}

const name = "filesystem"

// Name implement Collector's
func (f *FileSystem) Name() string {
	return name
}

// Collect implement Collector's
func (f *FileSystem) Collect() (result interface{}, err error) {
	result, err = getFileSystemInfo()
	return
}

package filesystem

type FileSystem struct{}

const name = "filesystem"

func (f *FileSystem) Name() string {
	return name
}

func (f *FileSystem) Collect() (result interface{}, err error) {
	result, err = getFileSystemInfo()
	return
}

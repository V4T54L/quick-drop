package repo

type FileRepo interface {
	AddFileMetadata(filename, filepath string) (id int, err error)
	GetFileMetadata(id int) (filename, filepath string, err error)
}

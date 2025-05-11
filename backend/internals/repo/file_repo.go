package repo

import "context"

type FileRepo interface {
	AddFileMetadata(ctx context.Context, filename string, outFilename string) (err error)
	GetFileMetadata(ctx context.Context, outFileName string) (filename string, err error)
}

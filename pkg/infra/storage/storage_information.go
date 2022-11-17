package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func (sp *StorageProvider) GetFileInformation(ctx context.Context, username string) ([]FileInformation, error) {
	var fileInformation []FileInformation
	it := sp.Handler.Objects(ctx, &storage.Query{
		Prefix: username,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		fileInformation = append(fileInformation,
			FileInformation{
				Name: attrs.Name,
				Size: attrs.Size,
			},
		)
	}

	return fileInformation, nil
}

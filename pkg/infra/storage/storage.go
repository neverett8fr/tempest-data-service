package storage

import (
	"context"
	"fmt"
	"io"

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

func (sp *StorageProvider) GetFileContent(ctx context.Context, username string, item string) ([]byte, error) {
	rdr, err := sp.Handler.Object(
		fmt.Sprintf("%s/%s", username, item),
	).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	byt, err := io.ReadAll(rdr)
	if err != nil {
		return nil, err
	}

	return byt, nil
}

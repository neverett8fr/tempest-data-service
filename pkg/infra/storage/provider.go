package storage

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

type StorageProvider struct {
	BucketName string
	Handler    *storage.BucketHandle
}

func InitialiseStorageProvider(ctx context.Context, bucketName string) (StorageProvider, error) {

	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return StorageProvider{}, err
	}

	return StorageProvider{
		Handler: client.Bucket(bucketName),
	}, nil
}

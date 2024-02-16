package storage

import (
	"errors"
	"mime/multipart"
	"mlogreport/app/config"

	"github.com/google/uuid"
	supabase "github.com/supabase-community/storage-go"
)

type storageConfig struct {
	sb *supabase.Client
}

type StorageInterface interface {
	Upload(image *multipart.FileHeader) (string, error)
}

func NewStorage(sb *supabase.Client) StorageInterface {
	return &storageConfig{
		sb: sb,
	}
}

func InitStorage(cfg *config.Config) *supabase.Client {
	storageClient := supabase.NewClient(cfg.STORAGE_URL, cfg.API_STORAGE, nil)
	return storageClient
}

var (
	contentType = "application/pdf"
	bucket      = "mlogreport"
)

func (sc *storageConfig) Upload(image *multipart.FileHeader) (string, error) {
	path := uuid.NewString()
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := sc.sb.UploadFile(bucket, path, file, supabase.FileOptions{
		ContentType: &contentType,
	})
	if err != nil {
		return "", errors.New(err.Error())
	}

	url := "https://cimxqffotlogzqvadisz.supabase.co/storage/v1/object/public/"+result.Key

	return url, nil
}

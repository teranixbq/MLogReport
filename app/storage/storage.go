package storage

import (
	"errors"
	"mime/multipart"
	"mlogreport/app/config"
	"mlogreport/utils/constanta"
	"mlogreport/utils/validation"

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
	upsert      = true
	datepath    = validation.DateAsia()
)

func (sc *storageConfig) Upload(image *multipart.FileHeader) (string, error) {

	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	result, err := sc.sb.UploadFile(bucket, image.Filename+"-"+datepath, file, supabase.FileOptions{
		ContentType: &contentType,
		Upsert:      &upsert,
	})
	if err != nil {
		return "", errors.New(err.Error())
	}
	url := constanta.URL_STORAGE+result.Key
	return url, nil
}

package storage

import (
	"errors"
	"log"
	"mime/multipart"
	"mlogreport/app/config"

	"github.com/google/uuid"
	supabase "github.com/supabase-community/storage-go"
)

type storageConfig struct {
	sb *supabase.Client
}

type StorageInterface interface {
	Upload(image *multipart.FileHeader)(string,error)
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
	path = uuid.NewString()
	bucket="mlogreport"
)

func (sc *storageConfig) Upload(image *multipart.FileHeader)(string,error) {
	file, err := image.Open()
    if err != nil {
        return "", errors.New("error: " + err.Error())
    }

    defer file.Close()
	result, err := sc.sb.UploadFile(bucket,path,file)
	if err != nil {
		return "",errors.New(err.Error())
	}

	log.Println(result)
	
	return result.Key,nil
}
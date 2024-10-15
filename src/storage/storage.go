package storage

import (
	"log"

	"github.com/TheAlpha16/phoros/src/config"
	"github.com/TheAlpha16/phoros/src/models"
)

var Store models.Storage

func InitStorage() error {
	switch config.OBJECT_STORE {
		case "native":
			Store = &models.NativeFS{}
		case "s3":
			Store = &models.S3FS{}
	}

	if err := Store.Authenticate(); err != nil {
		log.Println(err)
	}
	
	return nil
}
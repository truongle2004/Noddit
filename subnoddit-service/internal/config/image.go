package config

import (
	"log"
	"os"
	"subnoddit-service/internal/constant"
)

func createNewUploadFolder() {
	log.Println("Starting to create new folder upload image")
	if err := os.MkdirAll("upload", os.ModePerm); err != nil {
		log.Println("Error during creating new folder upload image")
	}
}

func ImageUploadConfig() error {
	// Make sure the dir upload exists
	if err := os.MkdirAll(constant.UploadImagePath, os.ModePerm); err != nil {
		// if the folder not exist, create new
		log.Println("Folder upload not exists, start to create a new one")
		createNewUploadFolder()
	}
	return nil
}

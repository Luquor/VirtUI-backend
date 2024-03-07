package models

import (
	"encoding/json"
	"log"
	"time"
	"virtui/api"
)

type Images struct {
	Metadata []string `json:"metadata"`
	api.StandardReturn
}

type Image struct {
	Metadata struct {
		Aliases []struct {
			Description string `json:"description"`
			Name        string `json:"name"`
		} `json:"aliases"`
		Architecture string    `json:"architecture"`
		AutoUpdate   bool      `json:"auto_update"`
		Cached       bool      `json:"cached"`
		CreatedAt    time.Time `json:"created_at"`
		ExpiresAt    time.Time `json:"expires_at"`
		Filename     string    `json:"filename"`
		Fingerprint  string    `json:"fingerprint"`
		LastUsedAt   time.Time `json:"last_used_at"`
		Profiles     []string  `json:"profiles"`
		Properties   struct {
			Os      string `json:"os"`
			Release string `json:"release"`
			Variant string `json:"variant"`
		} `json:"properties"`
		Public       bool   `json:"public"`
		Size         int    `json:"size"`
		Type         string `json:"type"`
		UpdateSource struct {
			Alias       string `json:"alias"`
			Certificate string `json:"certificate"`
			ImageType   string `json:"image_type"`
			Protocol    string `json:"protocol"`
			Server      string `json:"server"`
		} `json:"update_source"`
		UploadedAt time.Time `json:"uploaded_at"`
	} `json:"metadata"`
	api.StandardReturn
}

func GetImages() []Image {
	var images Images
	var imagesDetail []Image
	var imageDetail Image
	err := json.Unmarshal([]byte(api.Cli.Get("/1.0/images")), &images)
	for _, metadatum := range images.Metadata {
		err = json.Unmarshal([]byte(api.Cli.Get(metadatum)), &imageDetail)
		imagesDetail = append(imagesDetail, imageDetail)
	}
	if err != nil {
		log.Fatal(err)
	}
	return imagesDetail
}

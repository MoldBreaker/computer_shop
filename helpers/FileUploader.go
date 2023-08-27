package helpers

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	uploader2 "github.com/cloudinary/cloudinary-go/v2/api/uploader"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"mime/multipart"
	"os"
)

func UploadFiles(files []*multipart.FileHeader) ([]string, string, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("CLOUD_API_KEY"), os.Getenv("CLOUD_API_SECRET"))
	if err != nil {
		return nil, "error when setting cloudiary", err
	}
	var urls []string
	var ctx = context.Background()
	for i := 0; i < len(files); i++ {
		id, err := gonanoid.New()
		if err != nil {
			return nil, "error when generating id", err
		}
		src, err := files[i].Open()
		if err != nil {
			return nil, "error when opening file", err
		}
		defer src.Close()
		uploadResult, err := cld.Upload.Upload(ctx, src, uploader2.UploadParams{
			Folder:   os.Getenv("CLOUD_FOLDER"),
			PublicID: id,
		})
		if err != nil {
			return nil, "\"Failed to upload file, " + err.Error(), err
		}
		urls = append(urls, uploadResult.SecureURL)
	}
	return urls, "", nil
}

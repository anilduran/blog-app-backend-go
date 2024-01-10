package utils

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

var PresignerInstance *Presigner

type Presigner struct {
	PresignClient *s3.PresignClient
}

func InitializePresigner(s3Client *s3.Client) {

	PresignerInstance = &Presigner{
		PresignClient: s3.NewPresignClient(s3Client),
	}

}

func (presigner Presigner) PutObject(userId string) (string, string, error) {

	id := uuid.New()

	key := "images/" + userId + "/" + id.String() + ".jpg"

	request, err := presigner.PresignClient.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("anilduran-s3-bucket"),
		Key:    aws.String(key),
	}, func(opts *s3.PresignOptions) {
		// link will be expired in 15 minutes
		opts.Expires = time.Duration(60 * time.Second * 15)
	})

	if err != nil {
		return "", "", err
	}

	return request.URL, key, nil

}

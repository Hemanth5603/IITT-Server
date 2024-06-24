package helpers

import (
	"bytes"

	"github.com/Hemanth5603/IITT-Server/infrastructure"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFile(path string, binaryFile []byte) error {
	tfb := bytes.NewReader(binaryFile)

	toObject := s3.PutObjectInput{
		Bucket: aws.String("iittnif-bucket"),
		Key:    aws.String(path),
		Body:   tfb,
		ACL:    aws.String("public-read-write"),
	}

	_, uploadErr := infrastructure.S3_CLIENT.PutObject(&toObject)
	if uploadErr != nil {
		return uploadErr
	}
	return nil
}

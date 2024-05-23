package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
)

var S3_CLIENT *s3.S3
var S3_BUCKENT_NAME string

func InitializeSpaces() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	fmt.Printf("called init")
	accessKey := os.Getenv("SPACES_ACCESS_KEY")
	secretKey := os.Getenv("SPACES_SECRET_KEY")
	S3_BUCKENT_NAME = os.Getenv("S3_BUCKET_NAME")

	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:      aws.String("eu-north-1"),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		fmt.Printf("Failed to create s3 new session: %s", err)

		os.Exit(1)
	}

	S3_CLIENT = s3.New(newSession)
}

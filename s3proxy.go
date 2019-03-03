package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// const S3Endpoint = "http://localhost:4572"

func S3Put(localfile string) {
	s := session.Must(session.NewSession(&aws.Config{
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(endpoints.ApSoutheast2RegionID),
		// Endpoint:         aws.String(S3Endpoint),
	}))

	c := s3.New(s, &aws.Config{})

	rs, err := os.Open(localfile)
	if err != nil {
		panic(err)
	}
	defer rs.Close()

	p := s3.PutObjectInput{
		Bucket: aws.String("linustest123"),
		Key:    aws.String("./README.md"),
		ACL:    aws.String("public-read"),
		Body:   rs,
	}

	_, err = c.PutObject(&p)
	if err != nil {
		panic(err)
	}
}

func main() {
	S3Put("README.md")
}

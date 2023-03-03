package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"log"
	"testing"
)

func TestCreateBucket(t *testing.T) {
	_, err := client.CreateBucket(context.Background(), &object.CreateBucketInput{
		Bucket: &testBucket,
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func TestDeleteBucket(t *testing.T) {
	_, err := client.DeleteBucket(context.Background(), &object.DeleteBucketInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

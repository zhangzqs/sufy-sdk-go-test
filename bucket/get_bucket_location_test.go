package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestGetBucketLocation(t *testing.T) {
	bucket := "test-my-bucket"
	output, err := client.GetBucketLocation(context.Background(), &object.GetBucketLocationInput{
		Bucket: &bucket,
	})
	if err != nil {
		log.Fatalln(err)
	}
	spew.Println(output)
}

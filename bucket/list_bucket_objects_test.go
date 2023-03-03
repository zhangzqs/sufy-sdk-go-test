package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestListBucketObjects(t *testing.T) {
	output, err := client.ListObjects(context.Background(), &object.ListObjectsInput{
		Bucket: &testBucket,
	})
	if err != nil {
		log.Fatalln(err)
	}
	spew.Println(output)
}

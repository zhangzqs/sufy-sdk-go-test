package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"testing"
)

func TestHeadObject(t *testing.T) {
	key := "dir/1.txt"
	_, err := client.HeadObject(context.Background(), &object.HeadObjectInput{
		Bucket: &testBucket,
		Key:    &key,
	})
	if err != nil {
		panic(err)
	}
}

package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/service/object/types"
	"github.com/bachue/sufy-sdk-go/sufy"
	"testing"
)

func TestDeleteObject(t *testing.T) {
	key := "dir/1.txt"
	_, err := client.DeleteObject(context.Background(), &object.DeleteObjectInput{
		Bucket: &testBucket,
		Key:    &key,
	})
	if err != nil {
		panic(err)
	}
}

func TestDelete(t *testing.T) {
	_, err := client.DeleteObjects(context.Background(), &object.DeleteObjectsInput{
		Bucket: &testBucket,
		Delete: &types.Delete{
			Objects: []types.ObjectIdentifier{
				{
					Key: sufy.String("dir/1.txt"),
				},
				{
					Key: sufy.String("dir/2.txt"),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

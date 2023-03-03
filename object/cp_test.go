package service

import (
	"context"
	"fmt"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/service/object/types"
	"github.com/bachue/sufy-sdk-go/sufy"
	"testing"
)

func TestCopyObject(t *testing.T) {
	_, err := client.CopyObject(context.Background(), &object.CopyObjectInput{
		Bucket:       &testBucket,
		Key:          sufy.String(fmt.Sprintf("%s.copy", testKey)),
		CopySource:   sufy.String(fmt.Sprintf("%s/%s", testBucket, testKey)),
		StorageClass: types.StorageClassStandard,
		Metadata: map[string]string{
			"meta1": "1",
			"meta2": "2",
		},
	})
	if err != nil {
		panic(err)
	}
}

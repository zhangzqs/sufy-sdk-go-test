package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/service/object/types"
	"github.com/bachue/sufy-sdk-go/sufy"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

var contentType = "text/plain"
var testKey = "multi.txt"

func TestCreateMultipartUpload(t *testing.T) {
	output, err := client.CreateMultipartUpload(context.Background(), &object.CreateMultipartUploadInput{
		Bucket:       &testBucket,
		Key:          &testKey,
		ContentType:  &contentType,
		StorageClass: types.StorageClassStandard,
		Metadata: map[string]string{
			"testMate1": "1",
		},
	})
	if err != nil {
		t.Error(err)
	}

	if output.UploadId == nil {
		assert.NotNil(t, output)
	}
}

var uploadId = "63ff095f9bc82e0007f72de8region02z0"

func TestUploadPart(t *testing.T) {
	for i := 1; i <= 2; i++ {
		_, err := client.UploadPart(context.Background(), &object.UploadPartInput{
			Bucket:     &testBucket,
			Key:        &testKey,
			PartNumber: int32(i),
			UploadId:   &uploadId,
			Body:       strings.NewReader("test content " + strconv.Itoa(i) + " \n"),
		})
		if err != nil {
			panic(err)
		}
	}
}

func TestListMultipartUploads(t *testing.T) {
	_, err := client.ListMultipartUploads(context.Background(), &object.ListMultipartUploadsInput{
		Bucket: &testBucket,
		//UploadIdMarker: &uploadId,

	})
	if err != nil {
		panic(err)
	}
}

func TestCompleteMultipartUpload(t *testing.T) {
	_, err := client.CompleteMultipartUpload(context.Background(), &object.CompleteMultipartUploadInput{
		Bucket:   &testBucket,
		Key:      &testKey,
		UploadId: &uploadId,
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: []types.CompletedPart{
				{
					ETag:       sufy.String("a9e7b4ecedaa5e5abf37a96e9614c952"),
					PartNumber: int32(1),
				},
				{
					ETag:       sufy.String("15084d0b898d1b271c76b256500b3ec8"),
					PartNumber: int32(2),
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestAbortMultipartUpload(t *testing.T) {
	_, err := client.AbortMultipartUpload(context.Background(), &object.AbortMultipartUploadInput{
		Bucket:   &testBucket,
		UploadId: &uploadId,
		Key:      &testKey,
	})
	if err != nil {
		panic(err)
	}
}

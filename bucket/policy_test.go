package service

import (
	"context"
	"fmt"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/sufy"
	"testing"
)

func TestPutBucketPolicy(t *testing.T) {
	_, err := client.PutBucketPolicy(context.Background(), &object.PutBucketPolicyInput{
		Bucket: &testBucket,
		Policy: sufy.String(fmt.Sprintf(`{
			"Version": "sufy",
			"Id": "public",
			"Statement": [
			  {
				"Sid": "publicGet",
				"Effect": "Allow",
				"Principal": "*",
				"Action": ["miku:MOSGetObject"],
				"Resource": ["srn:miku:::%s/*"]
			  }
			]
		}`, testBucket)),
	})
	if err != nil {
		panic(err)
	}
}

func TestGetBucketPolicy(t *testing.T) {
	_, err := client.GetBucketPolicy(context.Background(), &object.GetBucketPolicyInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

func TestDeleteBucketPolicy(t *testing.T) {
	_, err := client.DeleteBucketPolicy(context.Background(), &object.DeleteBucketPolicyInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

func TestGetBucketPolicyStatus(t *testing.T) {
	_, err := client.GetBucketPolicyStatus(context.Background(), &object.GetBucketPolicyStatusInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

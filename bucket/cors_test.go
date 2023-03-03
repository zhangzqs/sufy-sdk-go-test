package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/service/object/types"
	"github.com/bachue/sufy-sdk-go/sufy"
	"testing"
)

func TestPutBucketCors(t *testing.T) {
	_, err := client.PutBucketCors(context.Background(), &object.PutBucketCorsInput{
		Bucket: &testBucket,
		CORSConfiguration: &types.CORSConfiguration{
			CORSRules: []types.CORSRule{
				{
					ID: sufy.String("corsRule1"),
					AllowedOrigins: []string{
						"x.com",
					},
					AllowedMethods: []string{
						"GET", "POST",
					},
					AllowedHeaders: []string{
						"Content-Type",
					},
					ExposeHeaders: []string{
						"x-sufy-server-side-encryption",
					},
					MaxAgeSeconds: 100,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestGetBucketCors(t *testing.T) {
	_, err := client.GetBucketCors(context.Background(), &object.GetBucketCorsInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

func TestDeleteBucketCors(t *testing.T) {
	_, err := client.DeleteBucketCors(context.Background(), &object.DeleteBucketCorsInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

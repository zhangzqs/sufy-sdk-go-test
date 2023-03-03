package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/service/object/types"
	"strings"
	"testing"
)

var (
	testErrorPage = "404.html"
	testIndexPage = "index.html"
)

func prepareErrorPage() {
	_, err := client.PutObject(context.Background(), &object.PutObjectInput{
		Bucket: &testBucket,
		Key:    &testErrorPage,
		Body:   strings.NewReader("404 page not found"),
	})
	if err != nil {
		panic(err)
	}
}

func prepareIndexPage() {
	_, err := client.PutObject(context.Background(), &object.PutObjectInput{
		Bucket: &testBucket,
		Key:    &testIndexPage,
		Body:   strings.NewReader("index page"),
	})
	if err != nil {
		panic(err)
	}
}

func TestPutBucketWebsite(t *testing.T) {
	prepareErrorPage()
	prepareIndexPage()
	_, err := client.PutBucketWebsite(context.Background(), &object.PutBucketWebsiteInput{
		Bucket: &testBucket,
		WebsiteConfiguration: &types.WebsiteConfiguration{
			ErrorDocument: &types.ErrorDocument{
				Key: &testErrorPage,
			},
			IndexDocument: &types.IndexDocument{
				Suffix: &testIndexPage,
			},
			RoutingRules: []types.RoutingRule{},
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestGetBucketWebsite(t *testing.T) {
	_, err := client.GetBucketWebsite(context.Background(), &object.GetBucketWebsiteInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

func TestDeleteBucketWebsite(t *testing.T) {
	_, err := client.DeleteBucketWebsite(context.Background(), &object.DeleteBucketWebsiteInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/davecgh/go-spew/spew"
	"log"
	"testing"
)

func TestListBuckets(t *testing.T) {
	output, err := client.ListBuckets(context.Background(), &object.ListBucketsInput{})
	if err != nil {
		log.Fatalln(err)
	}
	spew.Println(output)
}

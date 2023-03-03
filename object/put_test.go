package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/sufy"
	"log"
	"strings"
	"testing"
)

func TestPutObject(t *testing.T) {
	_, err := client.PutObject(context.Background(), &object.PutObjectInput{
		Bucket:      sufy.String(testBucket),
		Key:         sufy.String("dir/2.txt"),
		Body:        strings.NewReader("test content"),
		ContentType: sufy.String("text/plain"),
	})
	if err != nil {
		log.Fatalln(err)
	}
}

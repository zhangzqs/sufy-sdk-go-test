package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/sufy"
	"log"
	"testing"
)

func TestGetObject(t *testing.T) {
	_, err := client.GetObject(context.Background(), &object.GetObjectInput{
		Bucket: sufy.String(testBucket),
		Key:    sufy.String("dir/1.txt"),
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func TestPresignGetObjectURL(t *testing.T) {
	req, err := presignClient.PresignGetObject(context.Background(), &object.GetObjectInput{
		Bucket: sufy.String(testBucket),
		Key:    sufy.String("dir/1.txt"),
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(req)
}

func TestGet(t *testing.T) {

}

package service

import (
	"github.com/bachue/sufy-sdk-go-test/global"
	"github.com/bachue/sufy-sdk-go/service/object"
)

var client *object.Client
var testBucket string
var presignClient *object.PresignClient

func init() {
	client = global.Client
	testBucket = global.JsonConfigFile.TestBucket
	presignClient = global.PresignClient
}

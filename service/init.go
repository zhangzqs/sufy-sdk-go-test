package service

import (
	"github.com/bachue/sufy-sdk-go-test/global"
	"github.com/bachue/sufy-sdk-go/service/object"
)

var client *object.Client

func init() {
	client = global.Client
}

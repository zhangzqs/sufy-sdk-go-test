package global

import (
	"encoding/json"
	"github.com/aws/smithy-go/logging"
	"github.com/bachue/sufy-sdk-go/credentials"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/sufy"
	"github.com/bachue/sufy-sdk-go/sufy/policy/putpolicy"
	"github.com/bachue/sufy-sdk-go/sufy/signer/uptoken"
	"os"
	"time"
)

type JsonConfig struct {
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	Endpoint   string `json:"endpoint"`
	TestBucket string `json:"test_bucket"`
}

func LoadJsonConfigFile(filename string) *JsonConfig {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	var config JsonConfig
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return &config
}

var (
	Config         sufy.Config
	PutPolicy      putpolicy.PutPolicy
	Client         *object.Client
	JsonConfigFile *JsonConfig
	PresignClient  *object.PresignClient
)

func init() {
	JsonConfigFile = LoadJsonConfigFile("/Users/zzq/Desktop/test_sdk/sufy-example/config.json")
	Config = sufy.Config{
		Credentials: credentials.StaticCredentialsProvider{
			Value: sufy.Credentials{
				AccessKeyID:     JsonConfigFile.AccessKey,
				SecretAccessKey: JsonConfigFile.SecretKey,
			},
		},
		// Region:        "cn-east-1",
		Logger:        logging.NewStandardLogger(os.Stderr),
		ClientLogMode: sufy.LogSigning | sufy.LogRetries | sufy.LogRequestWithBody | sufy.LogResponseWithBody,
		EndpointResolverWithOptions: sufy.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (sufy.Endpoint, error) {
			return sufy.Endpoint{
				URL: JsonConfigFile.Endpoint,
			}, nil
		}),
	}

	PutPolicy = putpolicy.CreatePutPolicy(JsonConfigFile.TestBucket, time.Now().Add(1*time.Hour))
	Client = object.NewFromConfig(Config, func(options *object.Options) {
		options.UsePathStyle = true
		options.UpTokenRetriever = uptoken.NewSigner(PutPolicy, credentials.StaticCredentialsProvider{
			Value: sufy.Credentials{
				AccessKeyID:     JsonConfigFile.AccessKey,
				SecretAccessKey: JsonConfigFile.SecretKey,
			},
		})
	})
	PresignClient = object.NewPresignClient(Client)
}

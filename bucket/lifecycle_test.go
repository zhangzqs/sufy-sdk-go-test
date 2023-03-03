package service

import (
	"context"
	"github.com/bachue/sufy-sdk-go/service/object"
	"github.com/bachue/sufy-sdk-go/service/object/types"
	"github.com/bachue/sufy-sdk-go/sufy"
	"testing"
)

func TestPutBucketLifecycleConfiguration(t *testing.T) {
	_, err := client.PutBucketLifecycleConfiguration(context.Background(), &object.PutBucketLifecycleConfigurationInput{
		Bucket: &testBucket,
		LifecycleConfiguration: &types.BucketLifecycleConfiguration{
			Rules: []types.LifecycleRule{
				{
					ID: sufy.String("testRule"),
					Filter: &types.LifecycleRuleFilterMemberPrefix{
						Value: "dir/",
					},
					Transitions: []types.Transition{
						{
							Days:         10,
							StorageClass: types.TransitionStorageClassStandardIa,
						},
					},
					Expiration: &types.LifecycleExpiration{
						Days: 30,
					},
					Status: types.ExpirationStatusEnabled,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestGutBucketLifecycleConfiguration(t *testing.T) {
	_, err := client.GetBucketLifecycleConfiguration(context.Background(), &object.GetBucketLifecycleConfigurationInput{
		Bucket: &testBucket,
	})
	if err != nil {
		panic(err)
	}
}

func TestDeleteBucketLifecycle(t *testing.T) {
	client.DeleteBucketLifecycle(context.Background(), &object.DeleteBucketLifecycleInput{
		Bucket: &testBucket,
	})
}

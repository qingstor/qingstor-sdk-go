package main

import (
	"encoding/json"
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"

	qs "github.com/yunify/qingstor-sdk-go/service"
)

// BucketlifecycleFeatureContext provides feature context for bucket lifecycle.
func BucketLifecycleFeatureContext(s *godog.Suite) {
	s.Step(`^put bucket lifecycle:$`, putBucketLifecycle)
	s.Step(`^put bucket lifecycle status code is (\d+)$`, putBucketLifecycleStatusCodeIs)

	s.Step(`^get bucket lifecycle$`, getBucketLifecycle)
	s.Step(`^get bucket lifecycle status code is (\d+)$`, getBucketLifecycleStatusCodeIs)
	s.Step(`^get bucket lifecycle should have filter prefix "([^"]*)"$`, getBucketLifecycleShouldHaveFilterPrefix)

	s.Step(`^delete bucket lifecycle`, deleteBucketLifecycle)
	s.Step(`^delete bucket lifecycle status code is (\d+)$`, deleteBucketLifecycleStatusCodeIs)
}

// --------------------------------------------------------------------------

var putBucketLifecycleOutput *qs.PutBucketLifecycleOutput

func putBucketLifecycle(LifecycleJSONText *gherkin.DocString) error {
	putBucketLifecycleInput := &qs.PutBucketLifecycleInput{}
	err = json.Unmarshal([]byte(LifecycleJSONText.Content), putBucketLifecycleInput)
	if err != nil {
		return err
	}

	putBucketLifecycleOutput, err = bucket.PutLifecycle(putBucketLifecycleInput)
	return err
}

func putBucketLifecycleStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(putBucketLifecycleOutput.StatusCode), statusCode)
}

// --------------------------------------------------------------------------

var getBucketLifecycleOutput *qs.GetBucketLifecycleOutput

func getBucketLifecycle() error {
	getBucketLifecycleOutput, err = bucket.GetLifecycle()
	return err
}

func getBucketLifecycleStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(getBucketLifecycleOutput.StatusCode), statusCode)
}

func getBucketLifecycleShouldHaveFilterPrefix(prefix string) error {
	for _, Rule := range getBucketLifecycleOutput.Rule {
		if qs.StringValue(Rule.Filter.Prefix) == prefix {
			return nil
		}
	}

	return fmt.Errorf("filter prefix \"%s\" not found in bucket rules", prefix)
}

// --------------------------------------------------------------------------

var deleteBucketLifecycleOutput *qs.DeleteBucketLifecycleOutput

func deleteBucketLifecycle() error {
	deleteBucketLifecycleOutput, err = bucket.DeleteLifecycle()
	return err
}

func deleteBucketLifecycleStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(deleteBucketLifecycleOutput.StatusCode), statusCode)
}

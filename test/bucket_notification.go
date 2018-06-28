package main

import (
	"encoding/json"
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"

	//qs "github.com/yunify/qingstor-sdk-go/service"
  qs "github.com/qiaoqiaolaile/qingstor-sdk-go/service"
)

// BucketNotificationFeatureContext provides feature context for bucket notification.
func BucketNotificationFeatureContext(s *godog.Suite) {
	s.Step(`^put bucket notification:$`, putBucketNotification)
	s.Step(`^put bucket notification status code is (\d+)$`, putBucketNotificationStatusCodeIs)

	s.Step(`^get bucket notification$`, getBucketNotification)
	s.Step(`^get bucket notification status code is (\d+)$`, getBucketNotificationStatusCodeIs)
	s.Step(`^get bucket notification should have cloudfunc "([^"]*)"$`, getBucketNotificationShouldHavecloudfunc)

	s.Step(`^delete bucket notification`, deleteBucketNotification)
	s.Step(`^delete bucket notification status code is (\d+)$`, deleteBucketNotificationStatusCodeIs)
}

// --------------------------------------------------------------------------

var putBucketNotificationOutput *qs.PutBucketNotificationOutput

func putBucketNotification(NotificationJSONText *gherkin.DocString) error {
	putBucketNotificationInput := &qs.PutBucketNotificationInput{}
	err = json.Unmarshal([]byte(NotificationJSONText.Content), putBucketNotificationInput)
	if err != nil {
		return err
	}

	putBucketNotificationOutput, err = bucket.PutNotification(putBucketNotificationInput)
	return err
}

func putBucketNotificationStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(putBucketNotificationOutput.StatusCode), statusCode)
}

// --------------------------------------------------------------------------

var getBucketNotificationOutput *qs.GetBucketNotificationOutput

func getBucketNotification() error {
	getBucketNotificationOutput, err = bucket.GetNotification()
	return err
}

func getBucketNotificationStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(getBucketNotificationOutput.StatusCode), statusCode)
}

func getBucketNotificationShouldHavecloudfunc(cloudfunc string) error {
	for _, Notification := range getBucketNotificationOutput.Notifications {
		if qs.StringValue(Notification.Cloudfunc) == cloudfunc {
			return nil
		}
	}

	return fmt.Errorf("Cloudfunc \"%s\" not found in bucket Notifications", cloudfunc)
}

// --------------------------------------------------------------------------

var deleteBucketNotificationOutput *qs.DeleteBucketNotificationOutput

func deleteBucketNotification() error {
	deleteBucketNotificationOutput, err = bucket.DeleteNotification()
	return err
}

func deleteBucketNotificationStatusCodeIs(statusCode int) error {
	return checkEqual(qs.IntValue(deleteBucketNotificationOutput.StatusCode), statusCode)
}

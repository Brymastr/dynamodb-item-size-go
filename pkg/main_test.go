package main

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/go-playground/assert/v2"
)

func TestCalculateItemSize(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {

		item := map[string]types.AttributeValue{
			"sessionID":      &types.AttributeValueMemberS{Value: "sessionid"},  // 9 + 9
			"userID":         &types.AttributeValueMemberS{Value: "userid"},     // 6 + 6
			"createdAt":      &types.AttributeValueMemberS{Value: "datestring"}, // 9 + 10
			"sessionTimeout": &types.AttributeValueMemberN{Value: "1000"},       // 14 + 2 + 1
		}

		putItemInput := &dynamodb.PutItemInput{
			TableName: aws.String("table"),
			Item:      item,
		}

		actual := CalculateItemSize(putItemInput)

		assert.Equal(t, 114, actual)
	})
}

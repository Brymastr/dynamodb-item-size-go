package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brymastr/dynamodb-item-size-go/pkg"
)

type DBItem struct {
	AttributeA string   `dynamodbav:"attributeA"`
	AttributeB int      `dynamodbav:"attributeB"`
	AttributeC []byte   `dynamodbav:"attributeC"`
	AttributeD []string `dynamodbav:"attributeD,stringset"`
	AttributeE bool     `dynamodbav:"attributeE"`
}

func main() {
	dbItem := DBItem{
		AttributeA: "aaaaa",
		AttributeB: 12345,
		AttributeC: []byte("ccccc"),
		AttributeD: []string{"first string", "second string"},
		AttributeE: true,
	}

	avm, _ := attributevalue.MarshalMap(dbItem)

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String("example_table"),
		Item:      avm,
	}

	itemSize := pkg.CalculateItemSize(putItemInput)

	fmt.Println(itemSize) // => 138
}

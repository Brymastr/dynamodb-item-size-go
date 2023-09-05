package main

import (
	"math"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Calculate the approximate size of the item as it is stored
// in a DynamoDB table
// For reference: https://zaccharles.medium.com/calculating-a-dynamodb-items-size-and-consumed-capacity-d1728942eb7c
func CalculateItemSize(putInput *dynamodb.PutItemInput) int {
	totalKeySize := 0
	totalValueSize := 0
	for key, attributeValue := range putInput.Item {
		totalKeySize += len(key)
		totalValueSize += calculateAttributeValueSize(attributeValue)
	}
	return totalKeySize + totalValueSize + 48
}

func calculateAttributeValueSize(value types.AttributeValue) int {
	totalValueSize := 0
	switch v := value.(type) {
	case *types.AttributeValueMemberB:
		// binary is just the length of the value
		totalValueSize += len(v.Value)
	case *types.AttributeValueMemberBS:
		// similar to binary but the length of all values combined
		for _, val := range v.Value {
			totalValueSize += len(val)
		}
	case *types.AttributeValueMemberN:
		// numbers are tricky. 1 byte per 2 significant digits
		// plus 1 extra byte for good measure
		totalValueSize += int(math.Ceil(float64(len(v.Value))/2)) + 1
	case *types.AttributeValueMemberNS:
		// similar to numbers but the length of all values combined
		for _, num := range v.Value {
			totalValueSize += int(math.Ceil(float64(len(num))/2)) + 1
		}
	case *types.AttributeValueMemberS:
		// DyamoDB stores strings as UTF-8 binary encoding so
		// we must calculate the number of hex bytes
		totalValueSize += len([]byte(v.Value))
	case *types.AttributeValueMemberSS:
		// similar to strings but the length of all values combined
		for _, str := range v.Value {
			totalValueSize += len([]byte(str))
		}
	case *types.AttributeValueMemberNULL:
		// null consumes 1 byte
		totalValueSize += 1
	case *types.AttributeValueMemberBOOL:
		// boolean consumes 1 byte
		totalValueSize += 1
	case *types.AttributeValueMemberL:
		// minimum 3 bytes plus 1 per item plus the size of each item
		totalValueSize += 3
		for _, val := range v.Value {
			totalValueSize += calculateAttributeValueSize(val)
			totalValueSize += 1
		}
	case *types.AttributeValueMemberM:
		// minimum 3 bytes plus 1 per item plus the size of each item
		totalValueSize += 3
		for _, val := range v.Value {
			totalValueSize += calculateAttributeValueSize(val)
			totalValueSize += 1
		}
	}
	return totalValueSize
}

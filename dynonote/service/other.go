package service

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/nqbao/learn-go/dynonote/model"
)

func ResetDatabase() error {
	client := dynamodb.New(newSession())
	writeErr := 0

	for {
		output, err := client.Scan(&dynamodb.ScanInput{
			TableName: &tableName,
			Limit:     aws.Int64(25),
		})

		if err != nil {
			return err
		}

		if len(output.Items) == 0 {
			break
		}

		// batch write max is 25
		writeRequests := []*dynamodb.WriteRequest{}
		for _, item := range output.Items {
			note := &model.Note{}
			dynamodbattribute.UnmarshalMap(item, note)

			writeRequest := &dynamodb.WriteRequest{
				DeleteRequest: &dynamodb.DeleteRequest{
					Key: map[string]*dynamodb.AttributeValue{
						"uid": &dynamodb.AttributeValue{S: &note.UserKey},
						"nid": &dynamodb.AttributeValue{S: &note.ULID},
					},
				},
			}

			writeRequests = append(writeRequests, writeRequest)
		}

		fmt.Printf("Deleting %v items\n", len(writeRequests))

		batchOutput, err := client.BatchWriteItem(&dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				tableName: writeRequests,
			},
		})

		if err != nil {
			writeErr++
			log.Printf("error: %v", err)

			if writeErr > 5 {
				return err
			}
		}

		fmt.Printf("Unprocessed items %v\n", len(batchOutput.UnprocessedItems))
	}

	return nil
}

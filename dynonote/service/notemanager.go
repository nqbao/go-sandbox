package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/nqbao/learn-go/dynonote/model"
	"github.com/oklog/ulid"
)

var (
	region    = "us-east-1"
	tableName = "test.notes"
)

func newID() string {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

	id := ulid.MustNew(ulid.Now(), entropy)
	return id.String()
}

func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	return sess
}

func CreateNote(n *model.Note) {
	n.ULID = newID()
	putNote(n)
}

func UpdateNote(n *model.Note) {
	putNote(n)
}

// put item will replace item with same key
func putNote(n *model.Note) {
	client := dynamodb.New(newSession())

	av, err := dynamodbattribute.MarshalMap(n)

	input := &dynamodb.PutItemInput{
		TableName:              aws.String(tableName),
		Item:                   av,
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}
	_, err = client.PutItem(input)

	if err != nil {
		panic(err)
	}
	// fmt.Printf("consumed WCU: %v\n", output.ConsumedCapacity)
}

func GetNote(user string, id string) (note *model.Note) {
	client := dynamodb.New(newSession())

	output, err := client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"uid": &dynamodb.AttributeValue{
				S: aws.String(user),
			},
			"nid": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		panic(err)
	}

	if output.Item != nil {
		note = &model.Note{}
		dynamodbattribute.UnmarshalMap(output.Item, note)
	}

	return
}

func DeleteNote(user string, id string) {
	client := dynamodb.New(newSession())

	output, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"uid": &dynamodb.AttributeValue{
				S: aws.String(user),
			},
			"nid": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", output.ConsumedCapacity)
}

func GetUserNote(user string) (result []*model.Note) {
	result = nil

	client := dynamodb.New(newSession())

	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}

	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("uid").Equal(expression.Value(user)),
	).Build()

	if err != nil {
		panic(err)
	}

	input.SetKeyConditionExpression(*expr.KeyCondition())
	input.SetExpressionAttributeNames((expr.Names()))
	input.SetExpressionAttributeValues(expr.Values())

	err = client.QueryPages(input, func(output *dynamodb.QueryOutput, lastPage bool) bool {
		for _, item := range output.Items {
			note := &model.Note{}
			dynamodbattribute.UnmarshalMap(item, &note)

			result = append(result, note)
		}

		return true
	})

	if err != nil {
		panic(err)
	}

	return
}

func GetAllNotes() (result []*model.Note, err error) {
	client := dynamodb.New(newSession())
	err = client.ScanPages(&dynamodb.ScanInput{
		TableName:              aws.String(tableName),
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		for _, item := range page.Items {
			note := &model.Note{}
			dynamodbattribute.UnmarshalMap(item, note)

			result = append(result, note)
		}

		fmt.Printf("%v\n", page.ConsumedCapacity)

		return true
	})

	return
}

func StarNote(user string, id string, star bool) {
	client := dynamodb.New(newSession())

	expr, err := expression.NewBuilder().WithUpdate(
		expression.Set(expression.Name("star"), expression.Value(star)),
	).Build()

	if err != nil {
		panic(err)
	}

	out, err := client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"uid": &dynamodb.AttributeValue{
				S: aws.String(user),
			},
			"nid": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
		ReturnConsumedCapacity:    aws.String("TOTAL"),
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", out.ConsumedCapacity)
}

func IncrNoteViews(user string, id string, counter int) {
	client := dynamodb.New(newSession())

	expr, err := expression.NewBuilder().WithUpdate(
		expression.Add(expression.Name("views"), expression.Value(counter)),
	).Build()

	if err != nil {
		panic(err)
	}

	out, err := client.UpdateItem(&dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"uid": &dynamodb.AttributeValue{
				S: aws.String(user),
			},
			"nid": &dynamodb.AttributeValue{
				S: aws.String(id),
			},
		},
		ReturnConsumedCapacity:    aws.String("TOTAL"),
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", out.ConsumedCapacity)
}

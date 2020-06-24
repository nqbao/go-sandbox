package service

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/nqbao/learn-go/dynonote/model"
)

func StarNote(user string, id string, star int) {
	client := dynamodb.New(newSession())

	var expr expression.Expression
	var err error

	if star != 0 {
		expr, err = expression.NewBuilder().WithUpdate(
			expression.Set(expression.Name("star"), expression.Value(star)),
		).Build()

		if err != nil {
			panic(err)
		}
	} else {
		expr, err = expression.NewBuilder().WithUpdate(
			expression.Remove(expression.Name("star")),
		).Build()
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

func GetStarNotes(user string) (result []*model.Note) {
	result = nil

	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String("uid-star-index"),
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

	result, _, err = queryNotes(input)

	if err != nil {
		panic(err)
	}

	return
}

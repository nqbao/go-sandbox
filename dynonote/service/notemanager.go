package service

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aws/aws-dax-go/dax"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/nqbao/learn-go/dynonote/model"
	"github.com/oklog/ulid"
)

var (
	region      = "us-east-1"
	tableName   = "test.notes"
	daxEndpoint = "test2.8mmam2.clustercfg.dax.use1.cache.amazonaws.com:8111"
)

type NoteManager struct {
	session *session.Session
}

func NewNoteManager(creds *credentials.Credentials) *NoteManager {
	nm := &NoteManager{}

	if creds != nil {
		nm.session, _ = session.NewSession(&aws.Config{
			Region:      &region,
			Credentials: creds,
		})
	} else {
		nm.session = newSession()
	}

	return nm
}

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

func NewSession() *session.Session {
	return newSession()
}

func newDaxClient() dynamodbiface.DynamoDBAPI {
	cfg := dax.DefaultConfig()
	cfg.HostPorts = []string{daxEndpoint}
	cfg.Region = region
	cli, err := dax.New(cfg)

	if err != nil {
		panic(err)
	}

	return cli
}

func (nm *NoteManager) CreateNote(n *model.Note) error {
	n.Timestamp = time.Now().Unix()
	n.ULID = newID()
	return nm.putNote(n, false)
}

func (nm *NoteManager) UpdateNote(n *model.Note) error {
	return nm.putNote(n, true)
}

// put item will replace item with same key
func (nm *NoteManager) putNote(n *model.Note, check bool) error {
	client := dynamodb.New(nm.session)

	av, err := dynamodbattribute.MarshalMap(n)

	input := &dynamodb.PutItemInput{
		TableName:              aws.String(tableName),
		Item:                   av,
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}

	if check {
		expr, err := expression.NewBuilder().WithCondition(
			expression.And(
				expression.Name("user_id").Equal(expression.Value(n.UserKey)),
				expression.Name("timestamp").Equal(expression.Value(n.Timestamp)),
			),
		).Build()

		if err != nil {
			return err
		}

		input.SetConditionExpression(*expr.Condition())
		input.SetExpressionAttributeNames(expr.Names())
		input.SetExpressionAttributeValues(expr.Values())
	}

	_, err = client.PutItem(input)

	return err
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

func (nm *NoteManager) DeleteNote(user string, id int) error {
	client := dynamodb.New(nm.session)
	// client := newDaxClient()

	output, err := client.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": &dynamodb.AttributeValue{
				S: aws.String(user),
			},
			"timestamp": &dynamodb.AttributeValue{
				N: aws.String(strconv.Itoa(id)),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	})

	fmt.Printf("%v\n", output.ConsumedCapacity)
	return err
}

func (nm *NoteManager) GetUserNote(user string, limit int, startKey string) (result []*model.Note, err error) {
	result = nil

	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		ReturnConsumedCapacity: aws.String("TOTAL"),
		ScanIndexForward:       aws.Bool(false),
	}

	expr, err := expression.NewBuilder().WithKeyCondition(
		expression.Key("user_id").Equal(expression.Value(user)),
	).Build()

	if err != nil {
		return
	}

	input.SetKeyConditionExpression(*expr.KeyCondition())
	input.SetExpressionAttributeNames((expr.Names()))
	input.SetExpressionAttributeValues(expr.Values())

	if limit > 0 {
		input.SetLimit(int64(limit))
	}

	if startKey != "" {
		input.SetExclusiveStartKey(map[string]*dynamodb.AttributeValue{
			"uid": &dynamodb.AttributeValue{S: &user},
			"nid": &dynamodb.AttributeValue{S: &startKey},
		})
	}

	result, err = queryNotes(input)

	return
}

func queryNotes(input *dynamodb.QueryInput) (result []*model.Note, err error) {
	client := dynamodb.New(newSession())
	// client := newDaxClient()

	if *input.Limit > 0 {
		output, err := client.Query(input)

		if err != nil {
			return nil, err
		}

		for _, item := range output.Items {
			note := &model.Note{}
			dynamodbattribute.UnmarshalMap(item, &note)

			result = append(result, note)
		}

		fmt.Printf("%v\n", output.LastEvaluatedKey)
	} else {
		err = client.QueryPages(input, func(output *dynamodb.QueryOutput, lastPage bool) bool {
			for _, item := range output.Items {
				note := &model.Note{}
				dynamodbattribute.UnmarshalMap(item, &note)

				result = append(result, note)
			}

			return true
		})
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

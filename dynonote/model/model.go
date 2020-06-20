package model

type Note struct {
	UserKey     string `dynamodbav:"uid"`
	ULID        string `dynamodbav:"nid"`
	Title       string
	Content     string
	CategoryKey string
}

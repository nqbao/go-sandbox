package model

type Note struct {
	UserKey     string `dynamodbav:"user_id" json:"user_id"`     // hashkey
	Timestamp   int64  `dynamodbav:"timestamp" json:"timestamp"` // sortkey
	UserName    string `dynamodbav:"username" json:"username"`
	ULID        string `dynamodbav:"note_id" json:"note_id"`
	Title       string `dynamodbav:"title" json:"title"`
	Content     string `dynamodbav:"content" json:"content"`
	CategoryKey string `dynamodbav:"cat,omitempty" json:"cat,omitempty"`
	Star        int    `dynamodb:"star,omitempty" json:"star,omitempty"`
}

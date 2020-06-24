# Sample application for using DynamoDB in Go

Adapted from [Udemy dynamodb course](udemy.com/course/dynamodb)

Sample query:

  * Login / logout: aws cognito ✅
  * Read/Write new note per user: UID,ULID ✅
  * Get all notes: Scan ✅
  * Query notes of a single user: UID ✅
  * Add note to category: UID, ULID
  * Query note per category: GSI
  * Star/unstar notes: GSI ✅

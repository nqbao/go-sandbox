# Sample application for using DynamoDB in Go

Adapted from [Udemy dynamodb course](https://udemy.com/course/dynamodb)

Sample query:

  * Login / logout: aws cognito ✅
  * Read/Write new note per user: UID,ULID ✅
  * Get all notes: Scan ✅ (Not in UI)
  * Query notes of a single user: UID ✅
  * Add note to category: Update using UID, ULID (not in UI)
  * Query note per category: GSI (not in UI)
  * Star/unstar notes: GSI ✅ (not in UI)

# Sample application for using DynamoDB in GO

Sample query:

  * Login / logout: a table user (don't do that for now)
  * Read/Write new note per user: UID,ULID
  * Get all notes: Scan
  * Query notes of a single user: UID
  * Add note to category: UID, ULID
  * Query note per category: GSI
  * Star/unstar notes: GSI

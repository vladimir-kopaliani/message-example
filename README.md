# Message Example

This project was made for job interview.

## How to test

1. Open http://localhost:3000/playground

2. Paste:

```graphql
mutation writeMessage {
writeMessage(
  input: {
    id: "1"
    text: "hello"
  }
){
  id
  text
  createdAt
}
}

query getMessages {
getMessages(
  input: {
    startDate: "2006-01-02T00:00:05Z"
    endDate: "2021-01-02T00:00:05Z"
  }
){
  id
  text
  createdAt
}
}
```

3. On separate tab paste and launch:

```graphql
subscription {
  newMessagesSubscription{
    id
    text
    createdAt
  }
}
```

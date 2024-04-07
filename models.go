package main

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id       primitive.ObjectID
	username string
	email    string
	thoughts []Thought
	friends  []primitive.ObjectID
}

type Thought struct {
	id primitive.ObjectID
	thoughtText string
	createdAt string
	username string
	userId primitive.ObjectID
	reactions []Reaction
}

type Reaction struct {
	reactionId primitive.ObjectID
	reactionBody string
	username string
	userId primitive.ObjectID
	createdAt string
}
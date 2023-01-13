package types

import (
	"time"
)

type Validator int

const (
	Validator_STRING Validator = iota
	Validator_NUMBER Validator = iota
)

type Policy struct {
	Id         string
	Name       string
        Description string
	Created    time.Time
	Resources  []Resource
	Actors     []Actor
	Attributes map[string]string
}

type Actor struct {
	Name       string
	Validators []Validator
}

type Resource struct {
	Name        string
	Relations   []Relation
	Permissions []Permission
}

type Relation struct {
	Name  string
	//Kinds []string
}

type Permission struct {
	Name       string
	Expression string
}

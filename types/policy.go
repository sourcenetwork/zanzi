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
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	CreatedAt   time.Time         `json:"createdAt"`
	Resources   []Resource        `json:"resources"`
	Actors      []Actor           `json:"actors"`
	Attributes  map[string]string `json:"attributes"`
}

type Actor struct {
	Name       string `json:"name"`
	Validators []Validator
}

type Resource struct {
	Name        string       `json:"name"`
	Relations   []Relation   `json:"relations"`
	Permissions []Permission `json:"permissions"`
}

type Relation struct {
	Name string `json:"name"`
}

type Permission struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

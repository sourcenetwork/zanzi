package types

import (
    "time"
)


type Policy struct {
    Id string
    Name string
    Created time.Time
    Resources []Resource
    Actors []Actor
    Attributes map[string]string
}

type Actor struct {
    Name string
    Kinds []string
}

type Resource struct {
    Name string
    Relations []Relation
    Permissions []Permission
}

type Relation struct {
    Name string
    Kinds []string
}

type Permission struct {
    Name string
    RelationExpression string
}

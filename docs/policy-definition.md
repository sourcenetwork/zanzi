---
title: Policy YAML Reference
date: 2023-05-01
---

The follow is a reference for Zanzi's YAML Policy definition.

# Example

```yaml
version: 0.1

name: Decentralized Pastebin

description: Policy models a decentralized pastebin app where users can create snippets and choose who they share the snippet with.

resources:

  snippet: 
    doc: A text snippet
    relations:
      author:
        types:
          - did
      reader:
        doc: Explicit snippet reader. Can be either a group or a specific user
        types:
          - did
          - group

    permissions:
      # Defra permissions
      read:
        doc: Grants read permision to read snippet
        expr: (reader + author)
      update: 
        expr: (author)
      delete:
        expr: (author)

      # Domain permision
      can_comment: 
        doc: expresses the permissions to comment on a snippet. Any reader should be able to comment
        expr: (read)

  group:
    doc: Represents a set of Users
    relations:
      owner:
      member: 

  comment:
    doc: A comment in a snippet
    relations:
      author:
    permissions:
      delete:
        expr: (author)
      edit: 
        expr: (author)

actors:
  employee:
    doc: App users

  colaborators:

attributes:
  foo: bar
  key: value
```

# Policy Structure Reference

## version [required]

Specifies the version of the Policy schema being defined.
Currently supported values are:

- `0.1`

## name [required]

Gives a user defined friendly identifier for the Policy.
`name` can be any utf-8 compliant string.


## description [optional]

A verbose string used to describe the purpose of the Policy.
`description` can be any utf-8 compliant string.


## resources Definition Reference [required]

The `resources` section is a map from string to resource definition, the keys are used as the resource names in the Policy.

A resource definition represents an entity which is operated upon by external actors.
Resources are often tied to application resources / objects, such as documents, groups or files.


### doc [optional]

A string containing documentation regarding the resource, such as what it represents or notes of caution.

### relations Reference [required]

The `relations` section is a map from string to relations definitions, keys are the relation names.

A relation represents a potential relationship between an object and a system subject.


#### doc [optional]

A string containing documentation for a relation.


#### types [optional]

types is a list of resource names a relation can point to.

### permissions Reference [optional]

The `permissions` section is a map from string to permission definitions, where keys are used as permission names.

A permission definition models a operation over an object.
The permission object contains an expression which defines how a permission should be computed with respect to relations.
The permission expression is a key feature which allows the Policy system to model complex access control models, see the `expr` section for more information.

#### doc [optional]

A string containing documentation for what the permission represents.


#### expr [required]

A permission expression.

The Permission Relation Expression is a mini-language used to define how a permission is evaluated with respect to resource relations.

See permission expression mini-language reference for more information.

## actors Definition Reference [required]

The `actors` section is a map from strings to actor definitions, where each key specify the actor name. 

### doc [optional]

A string describing the actor type.

## attributes [optional]

A map of key-value string pairs for any user defined metadata.

Note that `attributes` are used as part of the Policy's iding algorithm, meaning that two identitical policies with different attributes will have different ids.

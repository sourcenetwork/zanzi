---
date: 2023-09-08
title: Relation Expression Mini Language
---

The Relation Expression Mini Language defines a simple grammar which describes a Relation Expression Tree.
This doument outlines the grammar and gives some examples on how users can use it.

# Background

A Relation Expression Tree defines a set of rules which are used to "evaluate a Relation" while traversing the Relation Graph.
These Rules are used to dynamically compute the set of Sucessors to a Relation Node in the Relation Graph.

The leaves in a Relation Expression Tree are the Rewrite Rules defined by Zanzibar (ie. This, Tuple to Userset, Computed Userset).
Internal nodes are set operations (Union, Intersection, Difference) which combines the set of "users" reachable through the resolved Relation Node.
The Relation Expression Mini Language represents those trees.

For a more in depth introduction to this topic, see [Grooking Zanzibar's Access Control](../grokking-zanzibar-relbac/readme.md)

# Language overview

The Mini Language was designed to resemble an algebraic expression.
The set operators are given by the tokens:

- `&` -> Set Intersection 
- `+` -> Set Union
- `-` -> Set Difference

Rewrite Rules are expressed by either an identifiers, two identifiers chainer or a keywords.
Example:

- `_this`: This Rewrite Rule
- `owner`: Computed Userset for `owner` relation
- `parent->owner`: Tuple to Userset rule, where `parent` is the Tupleset Filter and `owner` is the Computed Userset

An identifier is any valid utf8 alphanumeric character and some special characters such as `_`.
See the reference for full definition.

Grouping and operation precedence can be achieve using `()`.
For instance the expression `a + (b - c)`, would perform the union of set `a` with the difference of sets `b` and `c`.

Note: all operations in the mini language have the same precedence, meaning they will be evaluated from left to right.
Users should be mindful of that because not every set operation is associative.

# Examples

The following are explained examples of expressions written with the mini language:


The following expression states that the current relation should be evaluated as the union of a `_this` rule (meaning, fetch my sucessors as defined per the relationships) and the Computed Userset given by the `owner` relation.

```
_this + owner
```

The following expression defines a tree which is the Union of a `_this` rule with a Tuple to Userset rule where `parent` is the tupleset filter and `owner` is the Computed Userset relation.

```
_this + parent->owner
```

The following expression defines a tree which is the difference of the computed userset `reader` with the union of a `_this` rule with a Computed Userset for `owner`.
That is, take all explicit sucessors and all owners and take the difference from the set of all readers.

```
_this + owner - reader
```

# Language Reference

Currently the mini language is defined as:

```
expr := term | term tail+
tail := ( op , term )
op := union | diff | intersection
term := rule | subexpr
rule = cu | ttu
cu := identifier
ttu := identifier + arrow + identifier
subexpr := ( expr )
arrow := "->"
identifier := alphanum | "_"
union := "+"
intersection := "&"
difference := "-"
```

See [relation_expression_parser](../internal/relation_expression_parser/doc.go) package for full implementation

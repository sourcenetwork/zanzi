# Data Model

The classic zanzibar data model is represented by "relation tuples".
Relation tuples effectively represent a graph edge, where the node ids are the nodes themselves

`relation_tuple = (namespace, object_id, relation) -> (namespace, object_id, relation)

The implemented data model makes the node explicit.
An AuthNode essentially represents a Zanzibar userset.
Relationship represents an abstract relation tuple.

AuthNode = (namespace, object_id, relation) 
Relationship = GetObject() -> AuthNode, GetSubject() -> AuthNode

Relationships are an interface as it means users can freely store additional attributes with a relationship.

# Modules

repository
graph
tree
authorizer

authorizer uses repository , graph, tree
graph uses repository
tree uses repository

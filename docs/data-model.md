Data Model
==========

The current data model is heavily inspired by Zanzibar's original spec.

At it's heart, the data model defined consists of a 3-tuple identifying an object, a relation and a user/subject.
The flexibility of this data model comes from the fact that the user can represent a group of users in the shape of an (object, relation) pair, which is called an userset.

The indirection provided by usersets allows expressing an object relation graph.
A natural way to think about this graph is by taking each tuple to represent an edge in the relation graph.
Graph Nodes are implictly defined through tuples.
The existance of an edge between two nodes imply the existance of said Node.


Relation Graph
==============

As mentioned, Zanzibar's tuples define relationships in a graph.
A point of attention however is that there is a distinction between the graph defined by the tuples and the effective relation graph.

Due to "Userset Rewrites", the stored graph is actually a subset of all nodes and edges in the total graph.
Userset Rewrite rules are able to both create new nodes and edges in the total graph.
Nodes are created using the "Computed Userset" rule, while edges are created through the "Tuple To Userset" rule.

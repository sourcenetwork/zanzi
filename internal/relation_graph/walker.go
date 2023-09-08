package relation_graph

import (
     "context"

     "github.com/sourcenetwork/zanzi/pkg/domain"
    "github.com/sourcenetwork/zanzi/pkg/types"
)

func newWalker(repository NodeRepository, logger types.Logger) walker {
    return walker{
        logger: logger,
        repository: repository,
    }
}

// walker traversers through the explicit nodes within the relation graph and builds a forest of trees
// explicit nodes are derived from Policy Relationships
type walker struct {
    logger types.Logger
    repository NodeRepository
}

func (w *walker) Walk(ctx context.Context, policy *domain.Policy) (*domain.RelationTree, error) {
    edges, err := w.repository.ListEdges(ctx, policy.Id)
    if err != nil {
        return nil, err //TODO
    }

    trees := make(map[string]*domain.RelationTree)
    orphans := make(map[string]struct{})

    for _, edge := range edges {
        orphans[edge.First().Id()] = struct{}{}
    }

    for _, edge := range edges {
        parent := w.getTree(edge.First(), trees)
        child := w.getTree(edge.Second(), trees)
        parent.Children = append(parent.Children, child)
        delete(orphans, edge.Second().Id())
    }

    rootChildren := make([]*domain.RelationTree, 0, len(orphans))
    for orphan := range orphans {
        rootChildren = append(rootChildren, trees[orphan])
    }

    return &domain.RelationTree{
        Node: nil,
        Children: rootChildren,
    }, nil
}

func (w *walker) getTree(node *domain.RelationNode, trees map[string]*domain.RelationTree) *domain.RelationTree {
    tree, ok := trees[node.Id()]
    if !ok {
        tree = &domain.RelationTree{
            Node: node,
        }
        trees[node.Id()] = tree
    }
    return tree
}

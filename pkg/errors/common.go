package errors

const (
	AttrPolicy   string = "policy"
	AttrResource        = "resource"
	AttrRelation        = "relation"
	AttrMethod          = "method"
	AttrUserset         = "userset"
)

var ErrInvalidVariant = New("invalid type variant", Internal)
var ErrEntityExists = New("entity already exists", BadInput)

func ErrPolicyNotFound(polId string) error {
	return Wrap("policy not found", NotFound, Pair(AttrPolicy, polId))
}

// package relation_expression_parser implements a parser for the relation expression mini language
//
// Currently the parser recognizes the following language:
//
// expr := term | term tail+
// tail := ( op , term )
// op := union | diff | intersection
// term := rule | subexpr
// rule = cu | ttu
// cu := identifier
// ttu := identifier + arrow + identifier
// subexpr := ( expr )
// arrow := "->"
// identifier := alphanum | "_"
// union := "+"
// intersection := "&"
// difference := "-"
package relation_expression_parser

// Generated by GoGLL.
package ast

import (
	"g1/token"
)

type G0 interface{}
type E1 interface{}
type T1 string

// G0 : E1 ;
func G00(p0 interface{}) (interface{}, error) {
	return p0.(string), nil
}

// E1 : E1 + T1 ;
func E10(p0, p1, p2 interface{}) (interface{}, error) {
	e := p0.([]string)
	e = append(e, "+", p2.(string))
	return e, nil
}

// E1 : T1 ;
func E11(p0 interface{}) ([]string, error) {

	return []string{p0.(string)}, nil
}

// T1 : a ;
func T10(p0 interface{}) (string, error) {
	return p0.(*token.Token).LiteralString(), nil
}

package test_utils

import (
    "reflect"
    "testing"
    "strings"
)

// TestSuite represents a type which groups several tests.
// Ideally TestSuite is used in combination with the RunSuite function,
// which introspects the implementing type for all methods that start with "Test".
//
// A typical implementation of the Run method would be:
// ```
// func (s *MyTestSuite) Run(t *testing.T) {
//     s.setup()
//     RunSuite(s, t)
// }
// ```
type TestSuite interface {
    Run(*testing.T)
}

// Run a test suite type
// The suite type must have methods prefixed with "Test" and these methods must receive a
// *testing.T argument
func RunSuite(suite TestSuite, t *testing.T) {
    suiteVal := reflect.ValueOf(suite)
    suiteT := reflect.TypeOf(suite)

    var tests []reflect.Method
    for i := 0; i  < suiteT.NumMethod(); i++ {
        method := suiteT.Method(i)
        if strings.HasPrefix(method.Name, "Test") {
            tests = append(tests, method)
        }
    }

    for _, test := range tests {
        f := test.Func
        t.Run(test.Name, func(t *testing.T) {
            tVal := reflect.ValueOf(t)
            in := []reflect.Value {suiteVal, tVal}
            f.Call(in)
        })
    }
}

//  Copyright 2019 Marius Ackerman
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

// Package symbols generates a Go parser symbols package
package symbols

import (
	"bytes"
	"text/template"

	"github.com/goccmack/gogll/ast"
	"github.com/goccmack/goutil/ioutil"
)

type Data struct {
	NonTerminals []string
	Terminals    []string
}

func Gen(fname string, g *ast.GoGLL) {
	tmpl, err := template.New("symbols").Parse(src)
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, getData(g)); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(fname, buf.Bytes()); err != nil {
		panic(err)
	}
}

func getData(g *ast.GoGLL) *Data {
	return &Data{
		NonTerminals: g.NonTerminals.ElementsSorted(),
		Terminals:    g.Terminals.ElementsSorted(),
	}
}

const src = `
// Package symbols is generated by gogll. Do not edit.
package symbols

type Symbol interface{
	isSymbol()
	IsNonTerminal() bool
	String() string
}

func (NT) isSymbol() {}
func (T) isSymbol() {}

// NT is the type of non-terminals symbols
type NT int
const( {{range $i, $nt := .NonTerminals}}
	NT_{{$nt}} {{if not $i}}NT = iota{{end}}{{end}}
)

// T is the type of terminals symbols
type T int
const( {{range $i, $t := .Terminals}}
	T_{{$i}} {{if not $i}}T = iota{{end}} // {{$t}} {{end}}
)

type Symbols []Symbol

func (ss Symbols) Strings() []string {
	strs := make([]string, len(ss))
	for i, s := range ss {
		strs[i] = s.String()
	}
	return strs
}

func (NT) IsNonTerminal() bool {
	return true
}

func (T) IsNonTerminal() bool {
	return false
}

func (nt NT) String() string {
	return ntToString[nt]
}

func (t T) String() string {
	return tToString[t]
}

var ntToString = []string { {{range $nt := .NonTerminals}}
	"{{$nt}}", /* NT_{{$nt}} */{{end}} 
}

var tToString = []string { {{range $i, $t := .Terminals}}
	"{{$t}}", /* T_{{$i}} */{{end}} 
}

var stringNT = map[string]NT{ {{range $i, $sym := .NonTerminals}}
	"{{$sym}}":NT_{{$sym}},{{end}}
}
`
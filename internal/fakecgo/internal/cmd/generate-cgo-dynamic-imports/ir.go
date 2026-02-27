package main

import "text/template"

type SO struct {
	Name string
	Syms []Sym
}

type Sym interface {
	Name() string
	isSym()
}

type FuncSym struct {
	name   string
	Params []GoType
	Return GoType
}

func (s *FuncSym) isSym() {}

func (s *FuncSym) Name() string {
	return s.name
}

type GoType struct {
	Package *string
	Name    string
}

func Do(buildTags *string, sos []*SO) ([]OutItem, error) {
	var gob 
}

type OutItem struct {
	Type GoCodeType
	Data []byte
}

type GoCodeType int

const (
	GoCodeTypeGo GoCodeType = iota
	GoCodeTypeAsm
)

package main

//go:generate echo tokens-stringer
//go:generate stringer -type=Tokens
type Tokens int

const (
	Erc20 Tokens = iota
	Erc721
	Erc777
	Erc1155
)

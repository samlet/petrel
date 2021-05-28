package main

type Status int

const (
	Offline Status = iota
	Online
	Disable
	Deleted
)

package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
///go:generate go run entgo.io/contrib/entproto/cmd/entproto -path ./schema
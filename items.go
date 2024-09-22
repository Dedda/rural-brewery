package main

type Item interface {
	Id() int
	Name() string
	MaxStackSize() int
}

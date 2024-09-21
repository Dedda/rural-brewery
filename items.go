package main

type Item interface {
	Name() string
	MaxStackSize() int
}

package main

type Transaction struct {
	id           int
	locksManager LocksManager
	database     Database
}

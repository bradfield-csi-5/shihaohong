package main

type Database struct {
	tables map[string]Table
}

type Table struct {
	rows []Row
}

type Row struct {
	// assume a row is just one element for now
	val int
}

package main

import "fmt"

type (
	Database interface {
		Insert()
	}

	PostgresDb struct{}
	MockDb struct{}
) 

func NewPostgresDb() Database {
	return &PostgresDb{}
}

func (p *PostgresDb) Insert() {
	fmt.Println("Insert real db.")
}

func NewMockDb() Database {
	return &MockDb{}
}

func (p *MockDb) Insert() {
	fmt.Println("Insert mock db.")
}

func InsertPlayerItem(db Database){
	db.Insert()
}

func main(){
	pos := NewPostgresDb()
	mock := NewMockDb()
	InsertPlayerItem(pos)
	InsertPlayerItem(mock)
}
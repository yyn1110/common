package common

import "testing"
import _ "ac-common-go/mysql"

func TestOpenDatabase(t *testing.T) {
	opts := DBOptions{
		Addr:        "localhost:3306",
		User:        "root",
		Password:    "123456",
		Database:    "test",
		MaxOpenConn: 50,
		MaxIdleConn: 10,
	}
	db, err := OpenDatabase("mysql", opts)
	if err != nil {
		t.Fatalf("open data base: %v", err)
	}
	db.Close()
}

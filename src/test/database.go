package test

import "project/src/db"

var database = db.GetDBInstance()

func ClearTable() {
	database.DB.Exec("DELETE FROM customers")
	database.DB.Exec("ALTER SEQUENCE customers_id_seq RESTART WITH 1")
}

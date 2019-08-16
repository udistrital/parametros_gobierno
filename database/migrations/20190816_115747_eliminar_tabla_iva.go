package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarTablaIva_20190816_115747 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarTablaIva_20190816_115747{}
	m.Created = "20190816_115747"

	migration.Register("EliminarTablaIva_20190816_115747", m)
}

// Run the migrations
func (m *EliminarTablaIva_20190816_115747) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("DROP SEQUENCE IF EXISTS parametros_gobierno.iva_id_seq;")
	m.SQL("DROP TABLE IF EXISTS parametros_gobierno.iva;")
}

// Reverse the migrations
func (m *EliminarTablaIva_20190816_115747) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

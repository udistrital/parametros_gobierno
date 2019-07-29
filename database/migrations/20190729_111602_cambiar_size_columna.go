package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambiarSizeColumna_20190729_111602 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambiarSizeColumna_20190729_111602{}
	m.Created = "20190729_111602"

	migration.Register("CambiarSizeColumna_20190729_111602", m)
}

// Run the migrations
func (m *CambiarSizeColumna_20190729_111602) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("ALTER TABLE parametros_gobierno.rup_especialidad ALTER COLUMN nombre type character varying(255);")
}

// Reverse the migrations
func (m *CambiarSizeColumna_20190729_111602) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("ALTER TABLE parametros_gobierno.rup_especialidad ALTER COLUMN nombre type character varying(100);")
}
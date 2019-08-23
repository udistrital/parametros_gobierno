package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CambiarSizeNombreActividadEconomica_20190823_105612 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CambiarSizeNombreActividadEconomica_20190823_105612{}
	m.Created = "20190823_105612"

	migration.Register("CambiarSizeNombreActividadEconomica_20190823_105612", m)
}

// Run the migrations
func (m *CambiarSizeNombreActividadEconomica_20190823_105612) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE parametros_gobierno.actividad_economica ALTER COLUMN nombre type character varying(255);")
}

// Reverse the migrations
func (m *CambiarSizeNombreActividadEconomica_20190823_105612) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE parametros_gobierno.actividad_economica ALTER COLUMN nombre type character varying(100);")
}

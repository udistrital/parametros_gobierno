package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type EliminarConstraint_20190729_111601 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &EliminarConstraint_20190729_111601{}
	m.Created = "20190729_111601"

	migration.Register("EliminarConstraint_20190729_111601", m)
}

// Run the migrations
func (m *EliminarConstraint_20190729_111601) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo DROP CONSTRAINT fk_rup_tipo_especialidad_rup_grupo;")
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo RENAME rup_tipo_especialidad_id TO rup_especialidad_id;")
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo ADD CONSTRAINT fk_rup_especialidad_rup_grupo FOREIGN KEY (rup_especialidad_id) REFERENCES parametros_gobierno.rup_especialidad (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *EliminarConstraint_20190729_111601) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo DROP CONSTRAINT IF EXISTS fk_rup_especialidad_rup_grupo CASCADE;")
}
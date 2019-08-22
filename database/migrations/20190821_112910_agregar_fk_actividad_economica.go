package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AgregarFkActividadEconomica_20190821_112910 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AgregarFkActividadEconomica_20190821_112910{}
	m.Created = "20190821_112910"

	migration.Register("AgregarFkActividadEconomica_20190821_112910", m)
}

// Run the migrations
func (m *AgregarFkActividadEconomica_20190821_112910) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE parametros_gobierno.actividad_economica ADD CONSTRAINT fk_clasificacion_ciiu_actividad_economica FOREIGN KEY (clasificacion_ciiu_id) REFERENCES parametros_gobierno.clasificacion_ciiu (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *AgregarFkActividadEconomica_20190821_112910) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE parametros_gobierno.actividad_economica DROP CONSTRAINT fk_clasificacion_ciiu_actividad_economica FOREIGN KEY (clasificacion_ciiu_id) REFERENCES parametros_gobierno.clasificacion_ciiu (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
	
}

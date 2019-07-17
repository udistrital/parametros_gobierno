package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearLlavesForaneas_20190716_101007 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearLlavesForaneas_20190716_101007{}
	m.Created = "20190716_101007"

	migration.Register("CrearLlavesForaneas_20190716_101007", m)
}

// Run the migrations
func (m *CrearLlavesForaneas_20190716_101007) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("ALTER TABLE parametros_gobierno.nucleo_basico_conocimiento ADD CONSTRAINT fk_area_conocimiento_nucleo_basico_conocimiento FOREIGN KEY (area_conocimiento_id) REFERENCES parametros_gobierno.area_conocimiento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
m.SQL("ALTER TABLE parametros_gobierno.rup_especialidad ADD CONSTRAINT fk_rup_tipo_especialidad_rup_especialidad FOREIGN KEY (rup_tipo_especialidad_id) REFERENCES parametros_gobierno.rup_tipo_especialidad (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo ADD CONSTRAINT fk_rup_tipo_especialidad_rup_grupo FOREIGN KEY (rup_tipo_especialidad_id) REFERENCES parametros_gobierno.rup_tipo_especialidad (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;")
}

// Reverse the migrations
func (m *CrearLlavesForaneas_20190716_101007) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("ALTER TABLE parametros_gobierno.nucleo_basico_conocimiento DROP CONSTRAINT IF EXISTS fk_area_conocimiento_nucleo_basico_conocimiento CASCADE;")
m.SQL("ALTER TABLE parametros_gobierno.rup_especialidad DROP CONSTRAINT IF EXISTS fk_rup_tipo_especialidad_rup_especialidad CASCADE;")
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo DROP CONSTRAINT IF EXISTS fk_rup_tipo_especialidad_rup_grupo CASCADE;")
}
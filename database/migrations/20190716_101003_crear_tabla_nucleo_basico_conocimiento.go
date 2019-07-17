package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaNucleoBasicoConocimiento_20190716_101003 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaNucleoBasicoConocimiento_20190716_101003{}
	m.Created = "20190716_101003"

	migration.Register("CrearTablaNucleoBasicoConocimiento_20190716_101003", m)
}

// Run the migrations
func (m *CrearTablaNucleoBasicoConocimiento_20190716_101003) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE parametros_gobierno.nucleo_basico_conocimiento(id serial NOT NULL, nombre varchar(100) NOT NULL, descripcion varchar(100), codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, area_conocimiento_id integer NOT NULL, CONSTRAINT pk_area_conocimiento PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.nucleo_basico_conocimiento OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaNucleoBasicoConocimiento_20190716_101003) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.nucleo_basico_conocimiento")
}
package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaAreaConocimiento_20190716_101002 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaAreaConocimiento_20190716_101002{}
	m.Created = "20190716_101002"

	migration.Register("CrearTablaAreaConocimiento_20190716_101002", m)
}

// Run the migrations
func (m *CrearTablaAreaConocimiento_20190716_101002) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.area_conocimiento(id serial NOT NULL, nombre varchar(100) NOT NULL, descripcion varchar(100), codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_area_conocimiento PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.area_conocimiento OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaAreaConocimiento_20190716_101002) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.area_conocimiento")
}
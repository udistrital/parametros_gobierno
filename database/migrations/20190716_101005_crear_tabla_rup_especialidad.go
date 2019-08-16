package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRupEspecialidad_20190716_101005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRupEspecialidad_20190716_101005{}
	m.Created = "20190716_101005"

	migration.Register("CrearTablaRupEspecialidad_20190716_101005", m)
}

// Run the migrations
func (m *CrearTablaRupEspecialidad_20190716_101005) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.rup_especialidad(id serial NOT NULL, nombre varchar(100) NOT NULL, descripcion varchar(100), codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, rup_tipo_especialidad_id INTEGER NOT NULL, CONSTRAINT pk_rup_especialidad PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.rup_especialidad OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaRupEspecialidad_20190716_101005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.rup_especialidad")
}
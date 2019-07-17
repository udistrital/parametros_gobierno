package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRupTipoEspecialidad_20190716_101004 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRupTipoEspecialidad_20190716_101004{}
	m.Created = "20190716_101004"

	migration.Register("CrearTablaRupTipoEspecialidad_20190716_101004", m)
}

// Run the migrations
func (m *CrearTablaRupTipoEspecialidad_20190716_101004) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE parametros_gobierno.rup_tipo_especialidad(id serial NOT NULL, nombre varchar(100) NOT NULL, descripcion varchar(100), codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_rup_tipo_especialidad PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.rup_tipo_especialidad OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaRupTipoEspecialidad_20190716_101004) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.rup_tipo_especialidad")
}
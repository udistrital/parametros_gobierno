package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaRupGrupo_20190716_101006 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaRupGrupo_20190716_101006{}
	m.Created = "20190716_101006"

	migration.Register("CrearTablaRupGrupo_20190716_101006", m)
}

// Run the migrations
func (m *CrearTablaRupGrupo_20190716_101006) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE parametros_gobierno.rup_grupo(id serial NOT NULL, nombre varchar(100) NOT NULL, descripcion varchar(100), codigo_abreviacion varchar(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, rup_tipo_especialidad_id INTEGER NOT NULL, CONSTRAINT pk_rup_grupo PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.rup_grupo OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaRupGrupo_20190716_101006) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.rup_grupo")
}
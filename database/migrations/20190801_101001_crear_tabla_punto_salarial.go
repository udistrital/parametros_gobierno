package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaPuntoSalarial_20190801_101001 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaPuntoSalarial_20190801_101001{}
	m.Created = "20190801_101001"

	migration.Register("CrearTablaPuntoSalarial_20190801_101001", m)
}

// Run the migrations
func (m *CrearTablaPuntoSalarial_20190801_101001) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.punto_salarial(id serial NOT NULL, vigencia numeric(4,0) NOT NULL, valor numeric(10,2) NOT NULL, decreto varchar(255), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_punto_salarial PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.punto_salarial OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaPuntoSalarial_20190801_101001) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.punto_salarial")
}
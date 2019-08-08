package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaSalarioMinimo_20190801_101002 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaSalarioMinimo_20190801_101002{}
	m.Created = "20190801_101002"

	migration.Register("CrearTablaSalarioMinimo_20190801_101002", m)
}

// Run the migrations
func (m *CrearTablaSalarioMinimo_20190801_101002) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE TABLE parametros_gobierno.salario_minimo(id serial NOT NULL, vigencia numeric(4,0) NOT NULL, valor numeric(10,2) NOT NULL, decreto varchar(255), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_salario_minimo PRIMARY KEY (id));")
m.SQL("ALTER TABLE parametros_gobierno.salario_minimo OWNER TO desarrollooas;")


}

// Reverse the migrations
func (m *CrearTablaSalarioMinimo_20190801_101002) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
 m.SQL("DROP TABLE IF EXISTS parametros_gobierno.salario_minimo")
}
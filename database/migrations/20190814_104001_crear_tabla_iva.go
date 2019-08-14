package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaIva_20190814_104001 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaIva_20190814_104001{}
	m.Created = "20190814_104001"

	migration.Register("CrearTablaIva_20190814_104001", m)
}

// Run the migrations
func (m *CrearTablaIva_20190814_104001) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE parametros_gobierno.iva(id integer NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), vigencia numeric(4,0) NOT NULL, valor_porcentaje numeric(5,4) NOT NULL, decreto character varying(255), activo boolean NOT NULL, fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_iva PRIMARY KEY (id));")
    m.SQL("ALTER TABLE parametros_gobierno.iva OWNER TO desarrollooas;")
}

// Reverse the migrations
func (m *CrearTablaIva_20190814_104001) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_gobierno.iva")
}

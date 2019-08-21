package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaImpuesto_20190821_110534 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaImpuesto_20190821_110534{}
	m.Created = "20190821_110534"

	migration.Register("CrearTablaImpuesto_20190821_110534", m)
}

// Run the migrations
func (m *CrearTablaImpuesto_20190821_110534) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.impuesto( id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, tipo_impuesto_id integer NOT NULL, CONSTRAINT pk_impuesto PRIMARY KEY (id), CONSTRAINT fk_tipo_impuesto_impuesto FOREIGN KEY (tipo_impuesto_id) REFERENCES parametros_gobierno.tipo_impuesto(id) );")
	m.SQL("ALTER TABLE parametros_gobierno.impuesto OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_gobierno.impuesto IS 'Tabla que parametriza los diferentes subtipos de impuestos, cada impuesto debe estar asociado a un tipo de impuesto.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.id IS 'Identificador unico de la tabla impuesto.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de impuesto.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.impuesto.tipo_impuesto_id IS 'Referencia foranea a la tabla tipo_impuesto que almacena los diferentes impuestos.';")
	
}

// Reverse the migrations
func (m *CrearTablaImpuesto_20190821_110534) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_gobierno.impuesto")
}

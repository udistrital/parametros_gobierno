package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaVigenciaImpuesto_20190821_110938 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaVigenciaImpuesto_20190821_110938{}
	m.Created = "20190821_110938"

	migration.Register("CrearTablaVigenciaImpuesto_20190821_110938", m)
}

// Run the migrations
func (m *CrearTablaVigenciaImpuesto_20190821_110938) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.vigencia_impuesto( id serial NOT NULL, activo boolean NOT NULL, tarifa numeric(5,2) NOT NULL, porcentaje_aplicacion numeric(5,2), base_uvt INTEGER, base_pesos numeric(20,7), inicio_vigencia TIMESTAMP NOT NULL, fin_vigencia TIMESTAMP NOT NULL, decreto character varying(255), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, impuesto_id integer NOT NULL, CONSTRAINT pk_vigencia_impuesto PRIMARY KEY (id), CONSTRAINT fk_impuesto_vigencia_impuesto FOREIGN KEY (impuesto_id) REFERENCES parametros_gobierno.impuesto(id) );")
	m.SQL("ALTER TABLE parametros_gobierno.vigencia_impuesto OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_gobierno.vigencia_impuesto IS 'Tabla que permite registrar los diferentes valores aplicables para los impuestos segun periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.id IS 'Identificador unico de la tabla vigencia_impuesto.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.tarifa IS 'Porcentaje que aplica para el impuesto que se está definiendo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.porcentaje_aplicacion IS 'Porcentaje sobre el valor total.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.base_uvt IS 'Valor de impuesto expresado en numero de Unidades de Valor Tributario (UVT).';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.base_pesos IS 'Valor de impuesto expresado en pesos';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.inicio_vigencia IS 'Fecha en la que empieza a regir esta tarifa.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.fin_vigencia IS 'Fecha en la que termina de regir esta tarifa.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.decreto IS 'Decreto o ley nacional que establece el valor para esta tarifa, en determinado periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.impuesto_id IS 'Referencia foranea a la tabla impuesto que almacena el nombre y tipo, para el cual corresponde el valor y la vigencia.';")
	
}

// Reverse the migrations
func (m *CrearTablaVigenciaImpuesto_20190821_110938) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_gobierno.vigencia_impuesto")
}

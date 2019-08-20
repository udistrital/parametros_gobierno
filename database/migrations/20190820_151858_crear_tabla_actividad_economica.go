package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaActividadEconomica_20190820_151858 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaActividadEconomica_20190820_151858{}
	m.Created = "20190820_151858"

	migration.Register("CrearTablaActividadEconomica_20190820_151858", m)
}

// Run the migrations
func (m *CrearTablaActividadEconomica_20190820_151858) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.actividad_economica( id serial NOT NULL, nombre character varying(100) NOT NULL, codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, actividad_economica_padre integer, clasificacion_ciiu_id integer NOT NULL, CONSTRAINT pk_actividad_economica PRIMARY KEY (id), CONSTRAINT fk_actividad_economica_actividad_economica FOREIGN KEY (actividad_economica_padre) REFERENCES parametros_gobierno.actividad_economica (id) );")
	m.SQL("ALTER TABLE parametros_gobierno.actividad_economica OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_gobierno.actividad_economica IS 'Tabla que parametriza las actividades economicas del CIIU (CLASIFICACIÓN INDUSTRIAL INTERNACIONAL UNIFORME DE TODAS LAS ACTIVIDADES ECONÓMICAS).';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.id IS 'Identificador unico de la tabla actividad_economica.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.actividad_economica_padre IS 'Referencia foranea a la misma tabla, que establece la relación jerarquica de los registros.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.actividad_economica.clasificacion_ciiu_id IS 'Referencia foranea a la tabla clasificacion_ciiu que almacena las clasificaciones de la DIAN.';")
	
}

// Reverse the migrations
func (m *CrearTablaActividadEconomica_20190820_151858) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_gobierno.actividad_economica")
}

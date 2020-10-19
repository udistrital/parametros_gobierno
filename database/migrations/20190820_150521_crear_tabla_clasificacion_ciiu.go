package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaClasificacionCiiu_20190820_150521 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaClasificacionCiiu_20190820_150521{}
	m.Created = "20190820_150521"

	migration.Register("CrearTablaClasificacionCiiu_20190820_150521", m)
}

// Run the migrations
func (m *CrearTablaClasificacionCiiu_20190820_150521) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE IF NOT EXISTS parametros_gobierno.clasificacion_ciiu(id serial NOT NULL, nombre character varying(100) NOT NULL, descripcion character varying(100), codigo_abreviacion character varying(20), activo boolean NOT NULL, numero_orden numeric(5,2), fecha_creacion TIMESTAMP NOT NULL, fecha_modificacion TIMESTAMP NOT NULL, CONSTRAINT pk_clasificacion_ciiu PRIMARY KEY (id));")
	m.SQL("ALTER TABLE parametros_gobierno.clasificacion_ciiu OWNER TO desarrollooas;")
	m.SQL("COMMENT ON TABLE parametros_gobierno.clasificacion_ciiu IS 'Tabla que parametriza la clasificacion general de los codigos de clase del CIIU (CLASIFICACIÓN INDUSTRIAL INTERNACIONAL UNIFORME DE TODAS LAS ACTIVIDADES ECONÓMICAS).';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.id IS 'Identificador unico de la tabla clasificacion_ciiu.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de clasificacion_ciiu.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CrearTablaClasificacionCiiu_20190820_150521) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE IF EXISTS parametros_gobierno.clasificacion_ciiu")
}

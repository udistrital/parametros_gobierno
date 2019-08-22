package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CommentOnNucleoBasicoConocimiento_20190820_145651 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CommentOnNucleoBasicoConocimiento_20190820_145651{}
	m.Created = "20190820_145651"

	migration.Register("CommentOnNucleoBasicoConocimiento_20190820_145651", m)
}

// Run the migrations
func (m *CommentOnNucleoBasicoConocimiento_20190820_145651) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("COMMENT ON TABLE parametros_gobierno.nucleo_basico_conocimiento IS 'Tabla que parametriza el Nucleo Básico del Conocimiento NBC para las Profesiones segun el SNIES.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.id IS 'Identificador unico de la tabla nucleo_basico_conocimiento.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de nucleo_basico_conocimiento.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.area_conocimiento_id IS 'Referencia foranea a la tabla area_conocimiento que almacena las áreas del conocimiento según el SNIES.';")
	
}

// Reverse the migrations
func (m *CommentOnNucleoBasicoConocimiento_20190820_145651) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

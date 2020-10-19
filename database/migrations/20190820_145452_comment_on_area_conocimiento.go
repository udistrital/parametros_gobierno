package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CommentOnAreaConocimiento_20190820_145452 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CommentOnAreaConocimiento_20190820_145452{}
	m.Created = "20190820_145452"

	migration.Register("CommentOnAreaConocimiento_20190820_145452", m)
}

// Run the migrations
func (m *CommentOnAreaConocimiento_20190820_145452) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("COMMENT ON TABLE parametros_gobierno.area_conocimiento IS 'Tabla que parametriza las Áreas del Conocimiento según el SNIES.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.id IS 'Identificador unico de la tabla area_conocimiento.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de area_conocimiento.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.area_conocimiento.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
}

// Reverse the migrations
func (m *CommentOnAreaConocimiento_20190820_145452) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

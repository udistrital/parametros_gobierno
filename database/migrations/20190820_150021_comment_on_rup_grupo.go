package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CommentOnRupGrupo_20190820_150021 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CommentOnRupGrupo_20190820_150021{}
	m.Created = "20190820_150021"

	migration.Register("CommentOnRupGrupo_20190820_150021", m)
}

// Run the migrations
func (m *CommentOnRupGrupo_20190820_150021) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("COMMENT ON TABLE parametros_gobierno.rup_grupo IS 'Tabla que parametriza los Grupos del Registro Único de Proponentes RUP.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.id IS 'Identificador unico de la tabla rup_grupo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de rup_grupo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.rup_grupo.rup_especialidad_id IS 'Referencia foranea a la tabla rup_tipo_especialidad que almacena las especialidades del RUP.';")
	
}

// Reverse the migrations
func (m *CommentOnRupGrupo_20190820_150021) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

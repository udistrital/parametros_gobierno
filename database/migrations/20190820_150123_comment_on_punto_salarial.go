package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CommentOnPuntoSalarial_20190820_150123 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CommentOnPuntoSalarial_20190820_150123{}
	m.Created = "20190820_150123"

	migration.Register("CommentOnPuntoSalarial_20190820_150123", m)
}

// Run the migrations
func (m *CommentOnPuntoSalarial_20190820_150123) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("COMMENT ON TABLE parametros_gobierno.punto_salarial IS 'Tabla que parametriza los valores del punto salarial correspondientes a cada periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.id IS 'Identificador unico de la tabla punto_salarial.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.vigencia IS 'Indica la vigencia para la cual el salario minimo rige.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.valor IS 'Valor del punto salarial en Pesos.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.decreto IS 'Decreto o ley nacional que establece el valor para el salario minimo, en determinado periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.punto_salarial.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CommentOnPuntoSalarial_20190820_150123) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CommentOnSalarioMinimo_20190820_150217 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CommentOnSalarioMinimo_20190820_150217{}
	m.Created = "20190820_150217"

	migration.Register("CommentOnSalarioMinimo_20190820_150217", m)
}

// Run the migrations
func (m *CommentOnSalarioMinimo_20190820_150217) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("COMMENT ON TABLE parametros_gobierno.salario_minimo IS 'Tabla que parametriza los valores del salario minimo correspondientes a cada periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.id IS 'Identificador unico de la tabla salario_minimo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.vigencia IS 'Indica la vigencia para la cual el salario minimo rige.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.valor IS 'Valor del salario minimo en Pesos.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.decreto IS 'Decreto o ley nacional que establece el valor para el salario minimo, en determinado periodo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';")
	m.SQL("COMMENT ON COLUMN parametros_gobierno.salario_minimo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';")
	
}

// Reverse the migrations
func (m *CommentOnSalarioMinimo_20190820_150217) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}

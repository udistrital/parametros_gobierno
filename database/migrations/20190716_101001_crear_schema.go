package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearSchema_20190716_101001 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearSchema_20190716_101001{}
	m.Created = "20190716_101001"

	migration.Register("CrearSchema_20190716_101001", m)
}

// Run the migrations
func (m *CrearSchema_20190716_101001) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
m.SQL("CREATE SCHEMA IF NOT EXISTS parametros_gobierno;")
m.SQL("ALTER SCHEMA parametros_gobierno OWNER TO desarrollooas;")
m.SQL("SET search_path TO pg_catalog,public,parametros_gobierno;")

}

// Reverse the migrations
func (m *CrearSchema_20190716_101001) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
m.SQL("DROP SCHEMA IF EXISTS parametros_gobierno");
}
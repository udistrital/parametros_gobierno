CREATE SCHEMA parametros_gobierno;

-- Creación de secuencia y tabla area_conocimiento
CREATE SEQUENCE parametros_gobierno.area_conocimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.area_conocimiento(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.area_conocimiento_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_area_conocimiento PRIMARY KEY (id)

);

-- Creación de secuencia y tabla nucleo_basico_conocimiento

CREATE SEQUENCE parametros_gobierno.nucleo_basico_conocimiento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.nucleo_basico_conocimiento(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.nucleo_basico_conocimiento_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	area_conocimiento_id integer NOT NULL,
	CONSTRAINT pk_nucleo_basico_conocimiento PRIMARY KEY (id),
	CONSTRAINT fk_area_conocimiento_nucleo_basico_conocimiento FOREIGN KEY (area_conocimiento_id) REFERENCES parametros_gobierno.area_conocimiento(id)
);


-- Creación de secuencia y tabla rup_tipo

CREATE SEQUENCE parametros_gobierno.rup_tipo_especialidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.rup_tipo_especialidad(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.rup_tipo_especialidad_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_rup_tipo_especialidad PRIMARY KEY (id)
);



-- Creación de secuencia y tabla rup_especialidad
CREATE SEQUENCE parametros_gobierno.rup_especialidad_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.rup_especialidad(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.rup_especialidad_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	rup_tipo_especialidad_id INTEGER NOT NULL,
	CONSTRAINT pk_rup_especialidad PRIMARY KEY (id),
    CONSTRAINT fk_rup_tipo_especialidad_rup_especialidad FOREIGN KEY (rup_tipo_especialidad_id) REFERENCES parametros_gobierno.rup_tipo_especialidad(id)
);


-- Creación de secuencia y tabla rup_grupo

CREATE SEQUENCE parametros_gobierno.rup_grupo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.rup_grupo(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.rup_grupo_id_seq'::regclass),
	nombre character varying(100) NOT NULL,     -- 
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	rup_tipo_especialidad_id INTEGER NOT NULL,
	CONSTRAINT pk_rup_grupo PRIMARY KEY (id),
	CONSTRAINT fk_rup_tipo_especialidad_rup_grupo FOREIGN KEY (rup_tipo_especialidad_id) REFERENCES parametros_gobierno.rup_tipo_especialidad(id)
);

--######################################################
--Migracion: 20190729_111601_eliminar_constraint

ALTER TABLE parametros_gobierno.rup_grupo
DROP CONSTRAINT fk_rup_tipo_especialidad_rup_grupo;

ALTER TABLE parametros_gobierno.rup_grupo
RENAME rup_tipo_especialidad_id TO rup_especialidad_id;

ALTER TABLE parametros_gobierno.rup_grupo 
ADD CONSTRAINT fk_rup_especialidad_rup_grupo 
FOREIGN KEY (rup_especialidad_id) 
REFERENCES parametros_gobierno.rup_especialidad (id) 
MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;
--######################################################
--Migracion: 20190729_111602_cambiar_size_columna
ALTER TABLE parametros_gobierno.rup_especialidad 
ALTER COLUMN nombre type character varying(255);
--######################################################

-- Creación de secuencia y tabla punto salarial
--Migracion: 20190801_101001_crear_tabla_punto_salarial

CREATE SEQUENCE parametros_gobierno.punto_salarial_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.punto_salarial(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.punto_salarial_id_seq'::regclass),
	vigencia numeric(4,0) NOT NULL,
	valor numeric(10,2) NOT NULL,
	decreto character varying(255),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_punto_salarial PRIMARY KEY (id)
	--unique para el año?
);

-- Creación de secuencia y tabla salario minimo

CREATE SEQUENCE parametros_gobierno.salario_minimo_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.salario_minimo(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.salario_minimo_id_seq'::regclass),
	vigencia numeric(4,0) NOT NULL,
	valor numeric(10,2) NOT NULL,
	decreto character varying(255),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP,
	fecha_modificacion TIMESTAMP,
	CONSTRAINT pk_salario_minimo PRIMARY KEY (id)
	--unique para el año?
);

--######################################################
--######################################################

--Creacion tabla iva 
--Migracion: 20190814_104001_crear_tabla_iva
CREATE SEQUENCE parametros_gobierno.iva_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.iva(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.iva_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	vigencia numeric(4,0) NOT NULL,
	valor_porcentaje numeric(5,4) NOT NULL,
	decreto character varying(255),
	activo boolean NOT NULL,
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_iva PRIMARY KEY (id)
);

Migration: 20190816_115747_eliminar_tabla_iva

DROP SEQUENCE parametros_gobierno.iva_id_seq;
DROP TABLE parametros_gobierno.iva;

--Comentarios Tablas

--area_conocimiento 
COMMENT ON TABLE parametros_gobierno.area_conocimiento IS 'Tabla que parametriza las Áreas del Conocimiento según el SNIES';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.id IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.nombre IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.descripcion IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.codigo_abreviacion IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.activo IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.numero_orden IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.fecha_modificacion IS '';

--nucleo_basico_conocimiento
COMMENT ON TABLE parametros_gobierno.nucleo_basico_conocimiento IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.id IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.nombre IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.descripcion IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.codigo_abreviacion IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.activo IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.numero_orden IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.fecha_modificacion IS '';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.area_conocimiento_id IS '';
COMMENT ON CONSTRAINT pk_area_conocimiento_id ON parametros_gobierno.nucleo_basico_conocimiento IS '';

--rup_tipo_especialidad
COMMENT ON TABLE parametros_gobierno.rup_tipo_especialidad IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.id IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.nombre IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.descripcion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.codigo_abreviacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.activo IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.numero_orden IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.fecha_modificacion IS '';

--rup_especialidad
COMMENT ON TABLE parametros_gobierno.rup_especialidad IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.id IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.nombre IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.descripcion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.codigo_abreviacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.activo IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.numero_orden IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.fecha_modificacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.rup_tipo_especialidad_id IS '';
COMMENT ON CONSTRAINT pk_rup_tipo_especialidad_id ON parametros_gobierno.rup_especialidad IS '';

--rup_grupo
COMMENT ON TABLE parametros_gobierno.rup_grupo IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.id IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.nombre IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.descripcion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.codigo_abreviacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.activo IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.numero_orden IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_modificacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.rup_especialidad_id IS '';
COMMENT ON CONSTRAINT pk_rup_especialidad_id ON parametros_gobierno.rup_grupo IS '';

--punto_salarial
COMMENT ON TABLE parametros_gobierno.punto_salarial IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.id IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.vigencia IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.valor IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.decreto IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.activo IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.fecha_modificacion IS '';

--salario_minimo
COMMENT ON TABLE parametros_gobierno.salario_minimo IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.id IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.vigencia IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.valor IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.decreto IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.activo IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.fecha_modificacion IS '';


--######################################################
--######################################################
-- IMPUESTOS

--Creacion tabla tipo_impuesto 
--Migracion: 
CREATE SEQUENCE parametros_gobierno.tipo_impuesto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.tipo_impuesto(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.tipo_impuesto_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_tipo_impuesto PRIMARY KEY (id)
);
--Creacion tabla impuesto 
--Migracion: 
CREATE SEQUENCE parametros_gobierno.impuesto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.impuesto(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.impuesto_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	tipo_impuesto_id integer NOT NULL,
	CONSTRAINT pk_impuesto PRIMARY KEY (id),
	CONSTRAINT fk_tipo_impuesto_impuesto FOREIGN KEY (tipo_impuesto_id) REFERENCES parametros_gobierno.impuesto(id)
);
--Creacion tabla vigencia_impuesto 
--Migracion: 
CREATE SEQUENCE parametros_gobierno.impuesto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.impuesto(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.impuesto_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	tipo_impuesto_id integer NOT NULL,
	CONSTRAINT pk_impuesto PRIMARY KEY (id),
	CONSTRAINT fk_tipo_impuesto_impuesto FOREIGN KEY (tipo_impuesto_id) REFERENCES parametros_gobierno.impuesto(id)
);

--######################################################
--######################################################










--######################################################
--######################################################
### Stuff 
COMMENT ON COLUMN parametros_gobierno.rup_grupo.id IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.nombre IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.descripcion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.codigo_abreviacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.activo IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.numero_orden IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_creacion IS '';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_modificacion IS '';
###
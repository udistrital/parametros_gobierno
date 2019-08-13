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
	nombre character varying(100) NOT NULL,     -- 
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
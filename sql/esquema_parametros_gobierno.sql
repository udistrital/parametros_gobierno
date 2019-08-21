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

--######################################################################################################################################################
--######################################################################################################################################################
--######################################################################################################################################################
--######################################################################################################################################################
--######################################################################################################################################################
--Comentarios Tablas

--area_conocimiento 
--Migracion: 20190820_145452_comment_on_area_conocimiento
COMMENT ON TABLE parametros_gobierno.area_conocimiento IS 'Tabla que parametriza las Áreas del Conocimiento según el SNIES.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.id IS 'Identificador unico de la tabla area_conocimiento.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de area_conocimiento.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.area_conocimiento.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--nucleo_basico_conocimiento
--Migracion: 20190820_145651_comment_on_nucleo_basico_conocimiento
COMMENT ON TABLE parametros_gobierno.nucleo_basico_conocimiento IS 'Tabla que parametriza el Nucleo Básico del Conocimiento NBC para las Profesiones segun el SNIES.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.id IS 'Identificador unico de la tabla nucleo_basico_conocimiento.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de nucleo_basico_conocimiento.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.nucleo_basico_conocimiento.area_conocimiento_id IS 'Referencia foranea a la tabla area_conocimiento que almacena las áreas del conocimiento según el SNIES.';

--rup_tipo_especialidad
--Migracion: 20190820_145810_comment_on_rup_tipo_especialidad
COMMENT ON TABLE parametros_gobierno.rup_tipo_especialidad IS 'Tabla que parametriza los tipos de Especialidad del Registro Único de Proponentes RUP';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.id IS 'Identificador unico de la tabla rup_tipo_especialidad.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de rup_tipo_especialidad.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.rup_tipo_especialidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--rup_especialidad
--Migracion: 20190820_145920_comment_on_rup_especialidad
COMMENT ON TABLE parametros_gobierno.rup_especialidad IS 'Tabla que parametriza las Especialidades del Registro Único de Proponentes RUP.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.id IS 'Identificador unico de la tabla rup_especialidad.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de rup_especialidad.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.rup_especialidad.rup_tipo_especialidad_id IS 'Referencia foranea a la tabla rup_tipo_especialidad que almacena los tipos de especialidades del RUP.';

--rup_grupo
--Migracion: 20190820_150021_comment_on_rup_grupo
COMMENT ON TABLE parametros_gobierno.rup_grupo IS 'Tabla que parametriza los Grupos del Registro Único de Proponentes RUP.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.id IS 'Identificador unico de la tabla rup_grupo.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de rup_grupo.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.rup_grupo.rup_especialidad_id IS 'Referencia foranea a la tabla rup_tipo_especialidad que almacena las especialidades del RUP.';

--punto_salarial
--Migracion: 20190820_150123_comment_on_punto_salarial
COMMENT ON TABLE parametros_gobierno.punto_salarial IS 'Tabla que parametriza los valores del punto salarial correspondientes a cada periodo.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.id IS 'Identificador unico de la tabla punto_salarial.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.vigencia IS 'Indica la vigencia para la cual el salario minimo rige.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.valor IS 'Valor del punto salarial en Pesos.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.decreto IS 'Decreto o ley nacional que establece el valor para el salario minimo, en determinado periodo.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.punto_salarial.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--salario_minimo
--Migracion: 20190820_150217_comment_on_salario_minimo
COMMENT ON TABLE parametros_gobierno.salario_minimo IS 'Tabla que parametriza los valores del salario minimo correspondientes a cada periodo.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.id IS 'Identificador unico de la tabla salario_minimo.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.vigencia IS 'Indica la vigencia para la cual el salario minimo rige.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.valor IS 'Valor del salario minimo en Pesos.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.decreto IS 'Decreto o ley nacional que establece el valor para el salario minimo, en determinado periodo.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.salario_minimo.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

--######################################################################################################################################################
--######################################################################################################################################################

-- CIIU: CLASIFICACIÓN INDUSTRIAL INTERNACIONAL UNIFORME DE TODAS LAS ACTIVIDADES ECONÓMICAS

--Creacion tabla clasificacion_ciiu
--Migracion: 20190820_150521_crear_tabla_clasificacion_ciiu
CREATE SEQUENCE parametros_gobierno.clasificacion_ciiu_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.clasificacion_ciiu(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.clasificacion_ciiu_id_seq'::regclass),
	nombre character varying(100) NOT NULL,
	descripcion character varying(100),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	CONSTRAINT pk_clasificacion_ciiu PRIMARY KEY (id)
);
--clasificacion_ciiu
COMMENT ON TABLE parametros_gobierno.clasificacion_ciiu IS 'Tabla que parametriza la clasificacion general de los codigos de clase del CIIU (CLASIFICACIÓN INDUSTRIAL INTERNACIONAL UNIFORME DE TODAS LAS ACTIVIDADES ECONÓMICAS).';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.id IS 'Identificador unico de la tabla clasificacion_ciiu.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de clasificacion_ciiu.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.clasificacion_ciiu.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';


--Creacion tabla actividad_economica 
--Migracion: 20190820_151858_crear_tabla_actividad_economica
CREATE SEQUENCE parametros_gobierno.actividad_economica_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.actividad_economica(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.actividad_economica_id_seq'::regclass),
	nombre character varying(100) NOT NULL,	
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	actividad_economica_padre integer,
	clasificacion_ciiu_id integer NOT NULL,
	CONSTRAINT pk_actividad_economica PRIMARY KEY (id),
	CONSTRAINT fk_actividad_economica_actividad_economica FOREIGN KEY (actividad_economica_padre) REFERENCES parametros_gobierno.actividad_economica (id)
);
--actividad_economica
COMMENT ON TABLE parametros_gobierno.actividad_economica IS 'Tabla que parametriza las actividades economicas del CIIU (CLASIFICACIÓN INDUSTRIAL INTERNACIONAL UNIFORME DE TODAS LAS ACTIVIDADES ECONÓMICAS).';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.id IS 'Identificador unico de la tabla actividad_economica.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.actividad_economica_padre IS 'Referencia foranea a la misma tabla, que establece la relación jerarquica de los registros.';
COMMENT ON COLUMN parametros_gobierno.actividad_economica.clasificacion_ciiu_id IS 'Referencia foranea a la tabla clasificacion_ciiu que almacena las clasificaciones de la DIAN.';


--######################################################################################################################################################
--######################################################################################################################################################
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

--tipo_impuesto
COMMENT ON TABLE parametros_gobierno.tipo_impuesto IS 'Tabla que parametriza los diferentes tipos de impuestos como pueden ser IVA, Retefuentes entre otros.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.id IS 'Identificador unico de la tabla tipo_impuesto.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de tipo_impuesto.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.tipo_impuesto.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';

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
	CONSTRAINT fk_tipo_impuesto_impuesto FOREIGN KEY (tipo_impuesto_id) REFERENCES parametros_gobierno.tipo_impuesto(id)
);

--impuesto
COMMENT ON TABLE parametros_gobierno.impuesto IS 'Tabla que parametriza los diferentes subtipos de impuestos, cada impuesto debe estar asociado a un tipo de impuesto.';
COMMENT ON COLUMN parametros_gobierno.impuesto.id IS 'Identificador unico de la tabla impuesto.';
COMMENT ON COLUMN parametros_gobierno.impuesto.nombre IS 'Campo obligatorio de la tabla que indica el nombre del parámetro.';
COMMENT ON COLUMN parametros_gobierno.impuesto.descripcion IS 'Campo en el que se puede registrar una descripcion de la informacion de impuesto.';
COMMENT ON COLUMN parametros_gobierno.impuesto.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere.';
COMMENT ON COLUMN parametros_gobierno.impuesto.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.impuesto.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
COMMENT ON COLUMN parametros_gobierno.impuesto.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.impuesto.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.impuesto.tipo_impuesto_id IS 'Referencia foranea a la tabla tipo_impuesto que almacena los diferentes impuestos.';

--Creacion tabla vigencia_impuesto 
--Migracion: 
CREATE SEQUENCE parametros_gobierno.vigencia_impuesto_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

CREATE TABLE parametros_gobierno.vigencia_impuesto(
	id integer NOT NULL DEFAULT nextval('parametros_gobierno.vigencia_impuesto_id_seq'::regclass),
	activo boolean NOT NULL,
	tarifa numeric(5,2) NOT NULL, 
	porcentaje_aplicacion numeric(5,2),
	base_uvt INTEGER, 
	base_pesos numeric(20,7),
	inicio_vigencia TIMESTAMP NOT NULL,
	fin_vigencia TIMESTAMP NOT NULL,
	decreto character varying(255),
	fecha_creacion TIMESTAMP NOT NULL,
	fecha_modificacion TIMESTAMP NOT NULL,
	impuesto_id integer NOT NULL,
	CONSTRAINT pk_vigencia_impuesto PRIMARY KEY (id),
	CONSTRAINT fk_impuesto_vigencia_impuesto FOREIGN KEY (impuesto_id) REFERENCES parametros_gobierno.impuesto(id)
);
--vigencia_impuesto
COMMENT ON TABLE parametros_gobierno.vigencia_impuesto IS 'Tabla que permite registrar los diferentes valores aplicables para los impuestos segun periodo.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.id IS 'Identificador unico de la tabla vigencia_impuesto.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.activo IS 'Valor booleano para indicar si el registro esta activo o inactivo.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.tarifa IS 'Porcentaje que aplica para el impuesto que se está definiendo.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.porcentaje_aplicacion IS 'Porcentaje sobre el valor total.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.base_uvt IS 'Valor de impuesto expresado en numero de Unidades de Valor Tributario (UVT).';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.base_pesos IS 'Valor de impuesto expresado en pesos';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.inicio_vigencia IS 'Fecha en la que empieza a regir esta tarifa.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.fin_vigencia IS 'Fecha en la que termina de regir esta tarifa.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.decreto IS 'Decreto o ley nacional que establece el valor para esta tarifa, en determinado periodo.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.fecha_creacion IS 'Fecha y hora de la creación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.fecha_modificacion IS 'Fecha y hora de la ultima modificación del registro en la BD.';
COMMENT ON COLUMN parametros_gobierno.vigencia_impuesto.impuesto_id IS 'Referencia foranea a la tabla impuesto que almacena el nombre y tipo, para el cual corresponde el valor y la vigencia.';

--######################################################################################################################################################
--######################################################################################################################################################

--Añadir llave foranea faltante
--Migracion: 20190821_112910_agregar_fk_actividad_economica

ALTER TABLE parametros_gobierno.actividad_economica 
ADD CONSTRAINT fk_clasificacion_ciiu_actividad_economica 
FOREIGN KEY (clasificacion_ciiu_id) 
REFERENCES parametros_gobierno.clasificacion_ciiu (id) 
MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;





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
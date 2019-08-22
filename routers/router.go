// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/parametros_gobierno/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/nucleo_basico_conocimiento",
			beego.NSInclude(
				&controllers.NucleoBasicoConocimientoController{},
			),
		),

		beego.NSNamespace("/rup_tipo_especialidad",
			beego.NSInclude(
				&controllers.RupTipoEspecialidadController{},
			),
		),

		beego.NSNamespace("/rup_especialidad",
			beego.NSInclude(
				&controllers.RupEspecialidadController{},
			),
		),

		beego.NSNamespace("/rup_grupo",
			beego.NSInclude(
				&controllers.RupGrupoController{},
			),
		),

		beego.NSNamespace("/area_conocimiento",
			beego.NSInclude(
				&controllers.AreaConocimientoController{},
			),
		),

		beego.NSNamespace("/salario_minimo",
			beego.NSInclude(
				&controllers.SalarioMinimoController{},
			),
		),

		beego.NSNamespace("/punto_salarial",
			beego.NSInclude(
				&controllers.PuntoSalarialController{},
			),
		),

		beego.NSNamespace("/iva",
			beego.NSInclude(
				&controllers.PuntoSalarialController{},
			),
		),

		beego.NSNamespace("/clasificacion_ciiu",
			beego.NSInclude(
				&controllers.ClasificacionCiiuController{},
			),
		),

		beego.NSNamespace("/actividad_economica",
			beego.NSInclude(
				&controllers.ActividadEconomicaController{},
			),
		),

		beego.NSNamespace("/tipo_impuesto",
			beego.NSInclude(
				&controllers.TipoImpuestoController{},
			),
		),

		beego.NSNamespace("/impuesto",
			beego.NSInclude(
				&controllers.ImpuestoController{},
			),
		),

		beego.NSNamespace("/vigencia_impuesto",
			beego.NSInclude(
				&controllers.VigenciaImpuestoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

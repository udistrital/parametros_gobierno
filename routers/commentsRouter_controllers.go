package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:AreaConocimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:NucleoBasicoConocimientoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:PuntoSalarialController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupEspecialidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupGrupoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:RupTipoEspecialidadController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/parametros_gobierno/controllers:SalarioMinimoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

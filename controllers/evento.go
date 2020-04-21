package controllers

import (
	"encoding/json"

	"github.com/BOTOOM/sismos/models"
	"github.com/astaxie/beego"
)

// EventoController operations for Evento
type EventoController struct {
	beego.Controller
}

// URLMapping ...
func (c *EventoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Evento
// @Param	body		body 	{}	true		"body for Evento content"
// @Success 201 {}
// @Failure 403 body is empty
// @router / [post]
func (c *EventoController) Post() {
	var listaEventos []map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &listaEventos); err == nil {
		recorridoEncontrado := models.ClasificacionV2(listaEventos)
		c.Data["json"] = recorridoEncontrado
	} else {
		c.Ctx.Output.SetStatus(403)
	}
	c.ServeJSON()

}

// GetAll ...
// @Title GetAll
// @Description get Evento
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {}
// @Failure 403
// @router / [get]
func (c *EventoController) GetAll() {
	// logs.Info("entro al all")
	recorridoEncontrado := models.Clasificacion()
	c.Data["json"] = recorridoEncontrado
	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the Evento
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	{}	true		"body for Evento content"
// @Success 200 {}
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EventoController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Evento
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EventoController) Delete() {

}

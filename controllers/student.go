package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type StudentController struct {
	beego.Controller
}

func (c *StudentController) URLMapping() {
	c.Mapping("GetStudents", c.GetStudents)
	c.Mapping("GetStudent", c.GetStudent)
	c.Mapping("CreateStudent", c.CreateStudent)
	c.Mapping("UpdateStudent", c.UpdateStudent)
	c.Mapping("DeleteStudent", c.DeleteStudent)
}

func (c *StudentController) GetStudents() {
	// ctx := c.Ctx.Request.Context()
	// w := c.Ctx.ResponseWriter
	// id := c.Ctx.Input.Param(auditTrailIdPathParam)
	// audit, err := auditServices.GetAuditTrail(ctx, id)
	// if err != nil {
	// 	log.Error(ctx, err, "Get Audit Fail")
	// 	server.SetResponse(w, err)
	// }
	// server.SetResponse(w, response.Success)
	// err = c.Ctx.JSONResp(&audit)
	// if err != nil {
	// 	log.Error(ctx, err, "Write json response body payload failed")
	// 	server.SetResponse(w, err)
	// }
}

func (c *StudentController) GetStudent()    {}
func (c *StudentController) CreateStudent() {}
func (c *StudentController) UpdateStudent() {}
func (c *StudentController) DeleteStudent() {}

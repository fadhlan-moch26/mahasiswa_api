package controllers

import (
	studentServices "mahasiswa/service"

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
	ctx := c.Ctx.Request.Context()
	students, err := studentServices.GetStudents(ctx)
	if err != nil {
		return
	}
	c.JSONResp(&students)
}

func (c *StudentController) GetStudent()    {}
func (c *StudentController) CreateStudent() {}
func (c *StudentController) UpdateStudent() {}
func (c *StudentController) DeleteStudent() {}

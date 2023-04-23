package controllers

import (
	"fmt"
	"mahasiswa/models/codec"
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
		fmt.Println(err)
		return
	}
	c.JSONResp(&students)
}

func (c *StudentController) GetStudent() {
	ctx := c.Ctx.Request.Context()
	w := c.Ctx.ResponseWriter
	id := c.Ctx.Input.Param(studentIdPathParam)
	student, err := studentServices.GetStudent(ctx, id)
	if err != nil {
		fmt.Print(w)
		// log.Error(ctx, err, "Get Audit Fail")
		// server.SetResponse(w, err)
	}
	// server.SetResponse(w, response.Success)
	err = c.Ctx.JSONResp(&student)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func (c *StudentController) CreateStudent() {}

func (c *StudentController) UpdateStudent() {
	ctx := c.Ctx.Request.Context()
	w := c.Ctx.ResponseWriter
	req := codec.Student{}
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Print(w)
		return
	}
	id, err := studentServices.UpdateStudent(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.Ctx.JSONResp(id)
}

func (c *StudentController) DeleteStudent() {
	ctx := c.Ctx.Request.Context()
	w := c.Ctx.ResponseWriter
	req := codec.Student{}
	err := c.BindJSON(&req)
	if err != nil {
		fmt.Print(w)
		return
	}
	id, err := studentServices.DeleteStudent(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 	server.SetResponse(w, response.Success)
	c.Ctx.JSONResp(id)
}

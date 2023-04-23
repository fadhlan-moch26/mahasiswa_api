// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"mahasiswa/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/v1/students", &controllers.StudentController{}, "get:GetStudents")
	beego.Router("/v1/student/:id", &controllers.StudentController{}, "get:GetStudent")
	beego.Router("/v1/student", &controllers.StudentController{}, "post:CreateStudent")
	beego.Router("/v1/student/:id", &controllers.StudentController{}, "put:UpdateStudent")
	beego.Router("/v1/student/:id", &controllers.StudentController{}, "delete:DeleteStudent")
	beego.Router("/v1/major", &controllers.StudentController{}, "post:CreateMajor")
	beego.Router("/v1/hobby", &controllers.StudentController{}, "post:CreateHobby")
}

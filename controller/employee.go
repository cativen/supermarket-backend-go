package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
	"superMarket-backend/dto"
	"superMarket-backend/model"
	"superMarket-backend/response"
	"superMarket-backend/service"
	"superMarket-backend/utils"
)

func EmployeeList(c *gin.Context) {
	var dto dto.EmpQueryDTO
	err := c.ShouldBind(&dto)
	if err != nil {
		response.Fail(c, "fail", "操作失败")
	}
	if dto.PageSize == 0 {
		dto.SetDefaultPageSize()
	}
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.PageEmployeeByQo(c, dto)
}

func EmployeeDetail(c *gin.Context) {
	uid := c.Query("uid")
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		panic(err)
	}
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	detail := employeeService.Detail(id)
	response.Success(c, detail, "操作成功")
}

func EmployeeCreate(c *gin.Context) {

}

func UploadImg(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	for _, file := range files {
		log.Println(file.Filename)
		dst, newName := utils.UploadUrl(file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		var imgMap = make(map[string]interface{})
		imgMap["uploaded"] = 1 //成功
		imgMap["url"] = "/static/img" + newName
		c.JSON(200, imgMap)
		//response.Success(c, imgMap, "操作成功")
	}
}

func ViewImg(c *gin.Context) {
	filename := c.Param("filename")
	filePath := "./static/img/" + filename

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	// 获取文件的内容类型
	buffer := make([]byte, 512) // 创建一个512字节的缓冲区
	_, err = file.Read(buffer)  // 读取文件的前512字节
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	contentType := http.DetectContentType(buffer) // 检测内容类型

	// 设置文件的内容类型
	c.Header("Content-Type", contentType)

	// 将文件内容写入响应体
	c.File(filePath)

}

func UpdateEmployee(c *gin.Context) {
	var employee model.Employee
	token := c.GetHeader("token")
	c.ShouldBind(&employee)
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.UpdateEmployee(c, employee, token)
	response.Success(c, nil, "操作成功")
}

func EditEmployeeBtn(c *gin.Context) {
	uid := c.Query("uid")
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.GetEmpById(c, uid)
}

func DeactivateEmp(c *gin.Context) {
	var employee model.Employee
	c.ShouldBind(&employee)
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.DeactivateEmp(c, employee.ID)
	response.Success(c, nil, "操作成功")
}

func ResetEmpPwd(c *gin.Context) {
	eid, _ := c.GetPostForm("eid")
	code, _ := c.GetPostForm("code")
	var employeeService service.IEmployeeService = &service.EmployeeServiceImpl{}
	employeeService.ResetEmpPwd(c, eid, code)

}

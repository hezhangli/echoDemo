package controller

import (
	"echoDemo/user/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"net/http"
	"os"
)

type controller struct {
	uc usecase.Usecase
}

func NewController(uc usecase.Usecase) *controller {
	return &controller{uc}
}

type UserRequest struct {
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender bool   `json:"gender"`
}

func (ctrl *controller) CreateUser(c echo.Context) error {
	var req UserRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "error parsing request body")
	}

	err := ctrl.uc.CreateUser(c.Request().Context(), req.Name, req.Age, req.Gender)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error parsing request body")
	}

	return c.JSON(http.StatusOK, "")
}

// 文件上传
// 处理上传文件的控制器
func (ctrl *controller) Upload(c echo.Context) error {
	// 通过FormFile函数获取客户端上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	//打开用户上传的文件
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建目标文件，就是我们打算把用户上传的文件保存到什么地方
	// file.Filename 参数指的是我们以用户上传的文件名，作为目标文件名，也就是服务端保存的文件名跟用户上传的文件名一样
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// 这里将用户上传的文件复制到服务端的目标文件
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>文件上传成功: %s</p>", file.Filename))
}

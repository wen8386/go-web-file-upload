package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wen8386/go-web-file-upload/common"
	"net/http"
	"path/filepath"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK,"index.html", nil)
}

func upload(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	filename := filepath.Base(file.Filename)
	fileLocate := filepath.Join("upload", filename)

	if err := c.SaveUploadedFile(file, fileLocate); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	filetype := filepath.Ext(file.Filename)
	if filetype == ".jpg" || filetype == ".png" || filetype == ".gif" || filetype == ".jpeg" {
		common.ShowPic("upload/"+filename)
	}
	//fmt.Println(filetype)

	c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully ", file.Filename))
}

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", index)
	r.POST("/upload",upload)
	r.StaticFS("upload",http.Dir("view"))

	resIp := common.GetLocalIpAddress()

	r.StaticFS("/view", http.Dir("upload/"))

	common.CreateQR(string("http://"+resIp+":8080"))

	r.Run(":8080")
}


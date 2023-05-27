package main

import (
	"bytes"
	"fmt"

	// "io"
	"net/http"

	"github.com/andybalholm/brotli"
	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	// shell "github.com/ipfs/go-ipfs-api"
)

// 从ipfs获取文件
func GetFile() {

}

func AddFile(c *gin.Context) {
	//TODO 权限
	//先注释掉因为要测试
	// session := sessions.Default(c)
	// if session.Get("is_login") != true {
	// 	c.AbortWithStatus(http.StatusUnauthorized) //返回401
	// 	return
	// }

	f, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest) //400
		return
	}
	ff, _ := f.Open()
	defer ff.Close()

	buf := bytes.Buffer{} //输出的缓冲区
	b2 := bytes.Buffer{}  //输入的缓冲区,因为multipart.File未实现io.Writer
	b2.ReadFrom(ff)
	cl := brotli.NewWriter(&buf)
	cl.Write(b2.Bytes())
	// sh := shell.NewLocalShell() //需要挂着ipfs daemon
	// sh.Add(f)
	cl.Close()
	fmt.Println(len(buf.Bytes()))
	// fmt.Println(buf.String())
	fmt.Println(len(b2.Bytes()))
}
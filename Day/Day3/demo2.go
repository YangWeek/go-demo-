package Day3

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"` // 表单 binding
	Password string `form:"password" json:"password" binding:"required"`
}

// gin 中间件 通常以闭包的形式定义
func Mindleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before request")
		c.Next() // 继续处理请求 处理自己写的逻辑请求
		fmt.Println("after request")
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func Test4() {
	r := SetupRouter()
	r.LoadHTMLGlob("templates/**/*.html")
	r.Use(Mindleware()) // 使用中间件

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "首页",
		})
	})

	// gin 获取querystring参数
	r.GET("/user", func(context *gin.Context) {
		name := context.Query("name")
		age := context.Query("age") // 都是字符串
		context.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 获取请求体的数据
	/// 通用的绑定方法，它会根据 Content-Type 自动选择合适的绑定策略 支持绑定 JSON、XML、表单数据
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			// 数据传输失败
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("D:/temp/%s", file.Filename)
		// 上传文件到指定的目录
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			fmt.Printf("save file error:%s", err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	r.GET("/json", func(context *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		context.AsciiJSON(http.StatusOK, data)
	})

	// josnp
	//
	r.GET("/jsonp", func(context *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		context.JSONP(http.StatusOK, data)
	})

	// 表单
	// 处理 application/x-www-form-urlencoded数据
	r.POST("/submit", func(context *gin.Context) {
		name := context.PostForm("name")
		age := context.PostForm("age")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	//该方法专门用于绑定 JSON 格式的数据到结构体。它只处理 Content-Type 为 application/json 的请求
	r.POST("/submit2", func(context *gin.Context) {
		if err := context.ShouldBindJSON(&Login{}); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"name": context.PostForm("name"),
			"age":  context.PostForm("age"),
		})
	})

	r.POST("/submit3", func(context *gin.Context) {
		if err := context.ShouldBindQuery(&Login{}); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"name": context.PostForm("name"),
			"age":  context.PostForm("age"),
		})
	})

	r.Run(":9090")

	// 优雅的关闭web服务
	//quit := make(chan os.Signal, 1)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//
	//fmt.Println("Shutting down server...")
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//
	//if err := server.Shutdown(ctx); err != nil {
	//	fmt.Printf("Error: %v\n", err)
	//}
	//
	//fmt.Println("Server closed.")
}

package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  "golang.org/x/crypto/bcrypt"
)

// Binding from JSON
type Register struct {
  Username string `json:"username" binding:"required"`
  Password string `json:"password" binding:"required"`
  Avatar string `json:"avatar" binding:"required"`
  Fullname string `json:"fullname" binding:"required"`
}

// model Tbl_user
type Tbl_user struct {
	gorm.Model
	Username string 
  Password string
  Avatar string 
  Fullname string 
}

func main() {

	// connect mysql
	dsn := "root:7890@tcp(127.0.0.1:3306)/flutter_backend_go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// migrate database
	db.AutoMigrate(&Tbl_user{})

  r := gin.Default()

	// post
	r.POST("/register", func(c *gin.Context) {

		var json Register
    if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

    // if json.Username != "manu" || json.Password != "123" {
    //   c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
    //   return
    // }

    // c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10);

    // set str data
    user := Tbl_user{
      Username: json.Username,
      Password: string(encryptedPassword),
      Avatar: json.Avatar,
      Fullname: json.Fullname,
    }


    // create data
    db.Create(&user);

// check insert data to structure
    if(user.ID > 0) {
      c.JSON(http.StatusOK, gin.H{
        "message": "Finish insert data",
        "status code" : 201
      })
    }else {
      c.JSON(http.StatusOK, gin.H{
        "message": "Server Error",
        "status code" : 500
      })
    }

   
  })

	// get
	r.GET("/register", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
      "message": "register page",
    })
  })


	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
      "message": "start migrate pages",
    })
  })

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"time"
)
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(20);not null"`
	Telephone  string `gorm:"type:varchar(110);not null; unique"`
	Password  string `gorm:"size(255);not null"`
}

func  main()  {
		db := InitDB()
		//defer db.Close 可能最新版的gorm 不需要关闭了
		if err != nil {
			panic("failed to connect database")
		}
		r := gin.Default()
		r.POST("/api/auth/register", func(ctx *gin.Context) {
			// 获取参数
			name := ctx.PostForm("name")
			telephone := ctx.PostForm("telephone")
			password := ctx.PostForm("password")
			// 数据验证
			if len(telephone) != 11 {

				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422 ,"msg":"手机号必须为11位"})
				return
			}
			if len(password)<6 {
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422 ,"msg":"密码不能少于6位"})
				return
			}
			if len(name) == 0{
				name = RandomString(6)
			}

			// 注册用户
			newUser := User{
				Name : name,
				Telephone: telephone,
				Password : password,
			}
			db.Create(&newUser)
			// 返回结果
			ctx.JSON(200,gin.H{
				"mag":"注册成功",
			})
		})
		r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func RandomString(n int) string {

	var letters = []byte("qwertyuipasdgfhjklzxcbvnmQWERYTUIOPASDFHGJKLZXDCVCVNBM")
	result := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:= range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "qdgoweb"
	username := "root"
	password := "root"
 	charset := "utf-8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
		)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error() )
	}
	db.AutoMigrate(&User{})
	return db
}
func isTelephoneExist(db *gorm.DB, telephone string) bool{
	var user User

	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"qdgo/goweb/common"
	"qdgo/goweb/model"
	"qdgo/goweb/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDb()
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
		name = util.RandomString(6)
	}
	// 判断手机号是否为空
	if isTelephoneExist(DB, telephone){
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422 ,"msg":"用户已经存在"})
		return
	}

	// 注册用户
	newUser := model.User{
		Name : name,
		Telephone: telephone,
		Password : password,
	}
	DB.Create(&newUser)
	// 返回结果
	ctx.JSON(200,gin.H{
		"mag":"注册成功",
	})
}



func isTelephoneExist(db *gorm.DB, telephone string) bool{
	var user model.User

	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

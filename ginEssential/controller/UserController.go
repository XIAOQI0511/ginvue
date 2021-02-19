package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"qi.xiao/ginessential/common"
	"qi.xiao/ginessential/dto"
	"qi.xiao/ginessential/model"
	"qi.xiao/ginessential/response"
	"qi.xiao/ginessential/util"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//使用map获取请求的参数
	// var requestMap = make(map[string]string)
	// json.NewDecoder(ctx.Request.Body).Decode(&requestMap)//解析json到map
	//使用结构体获取参数
	var requestUser = model.User{}
	// json.NewDecoder(ctx.Request.Body).Decode(&requestUser)
	//使用gin框架bind方法获取参数
	ctx.Bind(&requestUser)
	//获取参数
	// name := ctx.PostForm("name")
	// telephone := ctx.PostForm("telephone")
	// password := ctx.PostForm("password")
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号为11位！")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位！")
		return
	}
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, password, telephone)

	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在！")
		return
	}
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	//使用gin框架bind方法获取参数
	ctx.Bind(&requestUser)
	//获取参数
	telephone := requestUser.Telephone
	password := requestUser.Password
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号为11位！")
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位！")
		return
	}
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
	// ctx.JSON(200, gin.H{
	// 	"code": 200,
	// 	"data": gin.H{"token": token},
	// 	"msg":  "登录成功",
	// })
}

func Info(ctx *gin.Context) {
	//用户通过验证，可从上下文中获取用户信息
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}}) //接口断言
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

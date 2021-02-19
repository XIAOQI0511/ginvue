package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"qi.xiao/ginessential/common"
	"qi.xiao/ginessential/model"
	"qi.xiao/ginessential/response"
)

//定义接口 借助编辑器自动生成常用代码
type ICategoryController interface {
	RestController
}

//为了方法的复用，定义结构体方法
type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	//创建新CategoryController实例
	return CategoryController{DB: db}

}
func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory = model.Category{}
	ctx.Bind(&requestCategory)
	fmt.Println("aaaaa" + requestCategory.Name)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return

	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "添加成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	//接收path和body的参数
	//绑定body中的参数
	var requestCategory = model.Category{}
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "数据验证错误，分类名称必填")
		return

	}
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	//更新分类
	//可传三种类型，map，struct，name value
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "删除失败，请重试")
		return
	}
	response.Success(ctx, nil, "")
}

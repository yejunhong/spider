package controller

import (
	"github.com/gin-gonic/gin"
	"spider/model"
	"strconv"
)

type Controller struct{
	Model *model.Model
}

/**
 *
 * 返回分页参数
 * (当前页数，显示条数，开始条数)
 */
func (controller *Controller) Page(c *gin.Context) (int64, int64, int64){

	var page string = c.DefaultQuery("page", "1")
	var pagesize string = c.DefaultQuery("pagesize", "10")
	
	pageInt64, _ := strconv.ParseInt(page, 10, 64)
	pagesizeInt64, _ := strconv.ParseInt(pagesize, 10, 64)

	var startNum int64 = ( pageInt64 - 1 ) * pagesizeInt64

	return pageInt64, pagesizeInt64, startNum
}

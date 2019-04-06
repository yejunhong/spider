package controller

import (
	"github.com/satori/go.uuid"
	"spider/lib"
	"spider/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"os"
)

/**
 *
 * 书籍资源
 *
 */
func (controller *Controller) CartoonResourceInfo(c *gin.Context) {
	var id string = c.Param("id")
	idInt64, _ := strconv.ParseInt(id, 10, 64)
	var res model.CartoonResource = controller.Model.GetCartoonById(idInt64)
	contents, _ := ioutil.ReadFile("./node/src/config/" + res.ConfigName + ".ts")
	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"info": res,
			"config": string(contents),
		},
	})
}

type FormCartoonResource struct {
	model.CartoonResource
	ConfigText string
}
/**
 *
 * 设置资源信息
 * @param c *gin.Context
 * @gin.JSON json
 *
 */
func (controller *Controller) SetCartoonResource(c *gin.Context) {
	var FromConfig FormCartoonResource
	if err := c.ShouldBindJSON(&FromConfig); err != nil {
		c.JSON(200, gin.H{"error": 100001, "msg": err.Error()})
		return
	}

	var file_name string = uuid.NewV4().String()

	var res model.CartoonResource = controller.Model.GetCartoonByResourceNo(FromConfig.ResourceNo)

	if res.Id > 0 {
		file_name = res.ConfigName
	}

	controller.Model.SetCartoonResource(model.CartoonResource{
		ResourceNo: FromConfig.ResourceNo,
		ResourceUrl: FromConfig.ResourceUrl,
		ResourceName: FromConfig.ResourceName,
		ConfigName: file_name,
		BookType: FromConfig.BookType,
	})
	file, err := os.OpenFile("./node/src/config/" + file_name + ".ts", os.O_RDWR|os.O_CREATE, 0766);
	if err != nil { // 打开文件失败
		c.JSON(200, gin.H{"error": 100002, "msg": err})
		return
	}
	i, err := file.Write([]byte(FromConfig.ConfigText))
	if err != nil { // 写入文件失败
		c.JSON(200, gin.H{"error": 100002, "msg": err})
		return
	}
	file.Close();
	c.JSON(200, gin.H{"error": 0, "msg": i})
}

func (controller *Controller) CartoonResource(c *gin.Context){
	var page, size, num int64 = controller.Page(c)
	var search string = c.Query("search")

	var list []model.CartoonResource
	var count int64
	list, count = controller.Model.GetCartoonResources(search, size, num)

	var No []interface{}
	for _, v := range list {
		No = append(No, v.ResourceNo)
	}

	var bookCount []model.BookCount = controller.Model.GetBookByResourceNoCount(No)
	var resBookCount map[string]int64 = map[string]int64{}
	
	for _, v := range bookCount {
		resBookCount[v.ResourceNo] = v.Number
	}

	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"page": page,
			"pagesize": size,
			"list": list,
			"book_count": resBookCount,
			"count": count,
		},
	})
}

/**
 *
 * 书籍列表
 *
 */
func (controller *Controller) CartoonList(c *gin.Context){
	var page, size, num int64 = controller.Page(c)
	var resource_no string = c.Query("resource_no")
	var search string = c.Query("search")

	var list []model.ResCartoonList
	var count int64

	list, count = controller.Model.GetCartoons(resource_no, search, size, num)
	
	var uniqueId []interface{}
	for k, v := range list {
		uniqueId = append(uniqueId, v.UniqueId)
		list[k].CdateText = lib.DateTime(v.Cdate)
	}

	var chaptersCount []model.ChaptersCount = controller.Model.GetChaptersFindByListUniqueIdCount(uniqueId)
	var resChaptersCount map[string]interface{} = map[string]interface{}{}
	
	for _, v := range chaptersCount {
		resChaptersCount[v.ListUniqueId] = v
	}

	c.JSON(200, gin.H{
		"error": 0,
		"msg": gin.H{
			"page": page,
			"pagesize": size,
			"list": list,
			"chapters_count": resChaptersCount,
			"count": count,
		},
	})
}

/**
 *
 * 书籍章节
 *
 */
func (controller *Controller) CartoonChapter(c *gin.Context){
	var list_unique_id string = c.Query("list_unique_id")
	var list []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(list_unique_id, -1)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": list,
	})
}

/**
 *
 * 书籍章节内容
 *
 */
func (controller *Controller) CartoonChapterContent(c *gin.Context){
	var chapter_unique_id string = c.Query("chapter_unique_id")
	var list []model.CartoonChapterContent = controller.Model.GetContentsFindByChapterUniqueId(chapter_unique_id)
	c.JSON(200, gin.H{
		"error": 0,
		"msg": list,
	})
}
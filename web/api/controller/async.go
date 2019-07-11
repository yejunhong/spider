package controller

import (
	"spider/model"
	"spider/lib"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"strings"
	"os"
	"sync"
  "encoding/json"
)

/**
 *
 * 书籍同步到生产库
 *
 */
 func (controller *Controller) AsyncBookProduce(c *gin.Context){

	var bookId string = c.Query("bookId")
	bookIdInt64, _ := strconv.ParseInt(bookId, 10, 64) 
	var bookInfo model.CartoonList = controller.Model.GetCartoonInfoById(bookIdInt64)
	var chapterList []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(bookInfo.UniqueId, 1)

	var resource = controller.Model.GetCartoonByResourceNo(bookInfo.ResourceNo)

	if resource.BookType == 1 { // 漫画
		go func() {
			fmt.Println("同步开始")
			var category map[string]CmfPortalCategory  = controller.ManhuaCategoryList()
			var chapterLists []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(bookInfo.UniqueId, 1)
			var portalBook CmfPortalPost = controller.ayncManhuaPortalPost(bookInfo)
			controller.ayncManhuaPortalChapter(bookInfo, chapterLists, portalBook)
			controller.ayncManhuaCategory(bookInfo, portalBook.Id, category)
			fmt.Println("同步结束")
			/*
			fmt.Println(len(chapterList))
			var bookList = controller.Model.GetSqlCartoonListByNo(bookInfo.ResourceNo)
			fmt.Println("需要同步：", len(bookList))
			var category map[string]CmfPortalCategory  = controller.ManhuaCategoryList()

			var post_title []interface{}
			for _, v := range bookList {
				post_title = append(post_title, v.ResourceName)
			}

			var postTitles = controller.GetManhuaListByPostTitle(post_title);
			var wait sync.WaitGroup
			var next chan int = make(chan int, 1) // 并发5

			for k, v := range bookList {

				if post, ok := postTitles[v.ResourceName]; ok {
					if post.UniqueId != v.UniqueId {
						continue
					}
				}

				wait.Add(1)
				next <- 1
				go func(info model.CartoonList) {
					var chapterLists []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(info.UniqueId, 1)
					var portalBook CmfPortalPost = controller.ayncManhuaPortalPost(info)
					controller.ayncManhuaPortalChapter(info, chapterLists, portalBook)
					controller.ayncManhuaCategory(info, portalBook.Id, category)
					<-next
					wait.Done()
				}(v)
				fmt.Printf("\r已同步：%d/%d", k + 1, len(bookList))
				os.Stdout.Sync()
				
			}
			wait.Wait()
			fmt.Println("同步完毕")
			*/
		}()

	} else {
		go func(){
			fmt.Println(len(chapterList))
			// controller.ayncXiaoShuo(bookInfo, chapterList)
			var bookList = controller.Model.GetSqlCartoonListByNo(bookInfo.ResourceNo)
			fmt.Println("需要同步：", len(bookList))
			var category map[string]CmfPortalCategory  = controller.CategoryList()

			var wait sync.WaitGroup
			var next chan int = make(chan int, 10) // 并发5
			for k, v := range bookList {
				wait.Add(1)
				go func(info model.CartoonList) {
					var chapterLists []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(info.UniqueId, 1)
					var portalBook CmfPortalPost = controller.ayncPortalPost(info) // 同步书籍
					controller.ayncPortalChapter(info, chapterLists, portalBook) // 同步章节
					controller.ayncPortalCategory(info, portalBook.Id, category) // 同步分类
					<-next
					wait.Done()
				}(v)
				fmt.Printf("\r已同步：%d", k)
				os.Stdout.Sync()
				next <- 1
			}
			wait.Wait()
			fmt.Println("同步完毕")
		}()
	}

	c.JSON(200, gin.H{
		"error": 0,
		"msg": "上传同步成功",
	})
}

type CmfPortalPost struct {
	Id int64
	UniqueId string
	PostTitle string
}

var src = "./static/"
// 小说文章管理
func(controller *Controller) ayncPortalPost(book model.CartoonList) CmfPortalPost {

	var path = "upload/bookcover/" + book.UniqueId + ".jpg"
	lib.DonwloadFile(src + path, book.ResourceImgUrl)

	var post_source = "完结"
	if book.IsEnd == 0 {
		post_source = ""
	}

	var bookData []map[string]interface{} = []map[string]interface{}{
		map[string]interface{}{
			"parent_id": 0, // '父级id',
			"post_type": 1, // '类型,1:文章;2:页面',
			"post_format": 1, // '内容格式;1:html;2:md',
			"user_id": 2, // '发表者用户id',
			"post_status": 0, // '状态;1:已发布;0:未发布;',
			"comment_status": 0, // '评论状态;1:允许;0:不允许',
			"is_top": 0, // '是否置顶;1:置顶;0:不置顶',
			"recommended": 0, //'是否推荐;1:推荐;0:不推荐',
			"post_buy": 0,
			"post_bookcase": 0, // '书柜量',
			"create_time": lib.Time(), // '创建时间',
			"update_time": lib.Time(), // '更新时间',
			"chapter_update_time": lib.Time(), //'章节更新时间',
			"post_title": book.ResourceName, // 'post标题',
			"post_excerpt": book.Detail,// 'post摘要',
			"post_source": post_source, // varchar(150) NOT NULL DEFAULT '' COMMENT '更新章节数',
			"more": `{"thumbnail":"` + path + `"}`,// '扩展属性,如缩略图;格式为json',
			"isfinish": book.IsEnd, // '写作进度是否完成 0连载中 1已完成',
			"isfree": 0, // '是否免费 1免费 0收费',
			"post_tag": 2, // '文章标识：1、漫画，2、小说',
			"adult": 1, // 18X--1：是，0：否
			// "file_path": "", // '小说文本存放的位置',
			"unique_id": book.UniqueId, // '数据同步唯一标识',
		},
	}
	model.DbBatchInsert(controller.Model.Db61, "cmf_portal_post", bookData, []string{"more", "post_source", "post_title", "post_excerpt", "isfinish"})

	// 修改同步信息
	controller.Model.UpdateCartoonListById(book.Id, map[string]interface{}{"is_async": 1})

	var bookInfo CmfPortalPost
	controller.Model.Db61.Where("unique_id = ?", book.UniqueId).Find(&bookInfo)
	return bookInfo
}

// 同步小说
// cmf_portal_category
// cmf_portal_category_post
func(controller *Controller) ayncPortalChapter(book model.CartoonList, chapter []model.CartoonChapter, portalBook CmfPortalPost) {
	
		var data []map[string]interface{}
		var chapter_price int = 0
		var ids []int64

    for _, v := range chapter {
			ids = append(ids, v.Id)
			var path = "upload/book/" + strconv.FormatInt(portalBook.Id, 10) + "/" + v.UniqueId + ".txt"
			var sort int = lib.InterceptStrNumberToInt(v.ResourceName)
			if sort > 5 {
				chapter_price = 25
			}

			data = append(data, map[string]interface{}{
				"status": 1, // '状态;1:显示;0:不显示',
				"price": chapter_price,// '价格 、观看金币。0为免费',
				"list_order": sort, // '排序',
				"chapter_excerpt":  book.Detail, // '摘要',
				"chapter_keywords": book.Detail,
				"chapter_content": path,
				"create_time": lib.Time(),
				"update_time": lib.Time(),
				"name": v.ResourceName, // '章节名称',
				"published_time": lib.Time(), // '发布时间',
				"pid": portalBook.Id, // '对应的上级ID',
				"unique_id": v.UniqueId, // '数据同步唯一标识',
			})
			lib.WriteFile(src + path, v.Content)
    }
    if len(data) > 0 {
			model.DbBatchInsert(controller.Model.Db61, "cmf_portal_chapter", data, []string{"name", "price", "chapter_excerpt", "chapter_content", "chapter_keywords", "list_order"})
		}
		controller.Model.UpdateCartoonChapterByIds(ids, map[string]interface{}{"is_async": 1})
}

type CmfPortalCategory struct {
	Id int64
	Name string
	Status int64
}

/**
 *
 * 分类列表
 *
 */
func (controller *Controller) CategoryList() map[string]CmfPortalCategory {
	var category []CmfPortalCategory
	controller.Model.Db61.Find(&category)
	var categoryMap map[string]CmfPortalCategory = map[string]CmfPortalCategory{}
	for _, v := range category {
		categoryMap[v.Name] = v
	}
	return categoryMap
}

type CmfPortalCategoryPost struct {
	Id int64
	PostId int64
}

/**
 * 同步小说分类数据
 */
func (controller *Controller) ayncPortalCategory(book model.CartoonList, pid int64, categoryMap map[string]CmfPortalCategory) {
	// Tags
	var tags []string = strings.Split(book.Tags, ",")
	var data []map[string]interface{}
    for _, v := range tags {
		if chapter, ok := categoryMap[v]; ok {
			data = append(data, map[string]interface{}{
				"category_id": chapter.Id,
				"status": chapter.Status, // 状态,1:发布;0:不发布
				"post_id": pid,// 
			})
		}
	}
	controller.Model.Db61.Where("post_id = ?", pid).Delete(CmfPortalCategoryPost{})
	if len(data) > 0 {
		model.DbBatchInsert(controller.Model.Db61, "cmf_portal_category_post", data, []string{})
	}
}

// 漫画书籍 =============================================================
// 小说文章管理
func(controller *Controller) ayncManhuaPortalPost(book model.CartoonList) CmfPortalPost {

	var post_source = "完结"
	if book.IsEnd == 0 {
		post_source = ""
	}

	var bookData []map[string]interface{} = []map[string]interface{}{
		map[string]interface{}{
			"parent_id": 0, // '父级id',
			"post_type": 1, // '类型,1:文章;2:页面',
			"post_format": 1, // '内容格式;1:html;2:md',
			"user_id": 2, // '发表者用户id',
			"post_status": 0, // '状态;1:已发布;0:未发布;',
			"comment_status": 0, // '评论状态;1:允许;0:不允许',
			"is_top": 0, // '是否置顶;1:置顶;0:不置顶',
			"recommended": 0, //'是否推荐;1:推荐;0:不推荐',
			"post_bookcase": 0, // '书柜量',
			"create_time": lib.Time(), // '创建时间',
			"update_time": lib.Time(), // '更新时间',
			"chapter_update_time": lib.Time(), //'章节更新时间',
			"post_buy": 0,
			"post_title": book.ResourceName, // 'post标题',
			"post_excerpt": book.Detail,// 'post摘要',
			"post_source": post_source, // varchar(150) NOT NULL DEFAULT '' COMMENT '更新章节数',
			"more": `{"thumbnail":"` + book.DownloadImgUrl + `"}`,// '扩展属性,如缩略图;格式为json',
			"isfinish": book.IsEnd, // '写作进度是否完成 0连载中 1已完成',
			"isfree": 0, // '是否免费 1免费 0收费',
			"post_tag": 1, // '文章标识：1、漫画，2、小说',
			// "adult": 1, // 18X--1：是，0：否
			"unique_id": book.UniqueId, // '数据同步唯一标识',
		},
	}
	model.DbBatchInsert(controller.Model.DbManhua, "cmf_portal_post", bookData, []string{"more", "post_source", "post_title", "post_excerpt", "isfinish", "chapter_update_time"})
	// 修改同步信息
	controller.Model.UpdateCartoonListById(book.Id, map[string]interface{}{"is_async": 1})
	var bookInfo CmfPortalPost
	controller.Model.DbManhua.Where("unique_id = ?", book.UniqueId).Find(&bookInfo)
	return bookInfo
}

// 同步漫画
func(controller *Controller) ayncManhuaPortalChapter(book model.CartoonList, chapter []model.CartoonChapter, portalBook CmfPortalPost) {
	
	var data []map[string]interface{}
	var ids []int64

	var img = map[string][]model.CartoonChapterContent{}
	var contents = controller.Model.GetContentsFindByChapterListUniqueId(book.UniqueId)

	for _, v := range contents {
		img[v.ChapterUniqueId] = append(img[v.ChapterUniqueId], v)
	}
	
	fmt.Println("获取数据所有内容：", book.ResourceName, len(img))
	var l = len(chapter)
	var str string = "";
    for k, v := range chapter {
		var chapter_price int = 0
		ids = append(ids, v.Id)
		if v.Sort > 2 {
			chapter_price = 48
		}
		var more = map[string]interface{}{
			"thumbnail": v.DownloadImgUrl,
			"files": []map[string]string{},
		}
		var photos = []map[string]string{}
		for k, img := range img[v.UniqueId] {
			photos = append(photos, map[string]string{
				"url": img.DownloadImgUrl,
				"name": strconv.Itoa((k + 1)) + ".jpg",
			})
		}
		more["photos"] = photos
		moreString, _ := json.Marshal(more)

		data = append(data, map[string]interface{}{
			"status": 1, // '状态;1:显示;0:不显示',
			"price": chapter_price,// '价格 、观看金币。0为免费',
			"list_order": v.Sort, // '排序',
			"chapter_excerpt":  book.Detail, // '摘要',
			"chapter_keywords": book.Detail,
			"chapter_content": book.Detail,
			"more": string(moreString),
			"create_time": lib.Time(),
			"update_time": lib.Time(),
			"name": v.ResourceName, // '章节名称',
			"published_time": lib.Time(), // '发布时间',
			"pid": portalBook.Id, // '对应的上级ID',
			"unique_id": v.UniqueId, // '数据同步唯一标识',
		})
		if (len(data) > 10) {
			model.DbBatchInsert(controller.Model.DbManhua, "cmf_portal_chapter", data, []string{"pid", "more", "name", "price", "chapter_excerpt", "chapter_content", "chapter_keywords", "list_order"})
			data = []map[string]interface{}{}
		}
		str = v.ResourceName
		fmt.Printf("\r已同步%d章节：%d/%d", portalBook.Id, k + 1, l)
		os.Stdout.Sync()
	}
	if len(data) > 0 {
		model.DbBatchInsert(controller.Model.DbManhua, "cmf_portal_chapter", data, []string{"more", "name", "price", "chapter_excerpt", "chapter_content", "chapter_keywords", "list_order"})
	}
	controller.Model.UpdateCartoonChapterByIds(ids, map[string]interface{}{"is_async": 1})
	if book.IsEnd == 0 {
		controller.Model.DbManhua.Model(&CmfPortalPost{}).Where("id = ?", portalBook.Id).Update("post_source", str)
	}
	fmt.Println()
}

/**
 *
 * 漫画分类列表
 *
 */
 func (controller *Controller) ManhuaCategoryList() map[string]CmfPortalCategory {
	var category []CmfPortalCategory
	controller.Model.DbManhua.Find(&category)
	var categoryMap map[string]CmfPortalCategory = map[string]CmfPortalCategory{}
	for _, v := range category {
		categoryMap[v.Name] = v
	}
	return categoryMap
}

/**
 * 同步漫画分类数据
 */
 func (controller *Controller) ayncManhuaCategory(book model.CartoonList, pid int64, categoryMap map[string]CmfPortalCategory) {
	// Tags
	var tags []string = strings.Split(book.Tags, ",")
	var data []map[string]interface{}
    for _, v := range tags {
		if chapter, ok := categoryMap[v]; ok {
			data = append(data, map[string]interface{}{
				"category_id": chapter.Id,
				"status": chapter.Status, // 状态,1:发布;0:不发布
				"post_id": pid,// 
			})
		}
	}
	controller.Model.DbManhua.Where("post_id = ?", pid).Delete(CmfPortalCategoryPost{})
	if len(data) > 0 {
		model.DbBatchInsert(controller.Model.DbManhua, "cmf_portal_category_post", data, []string{})
	}
}


/**
 *
 * 通过标题 获取漫画资源书籍
 * @return map[string]CmfPortalPost
 *
 */
 func (controller *Controller) GetManhuaListByPostTitle(post_title []interface{}) map[string]CmfPortalPost{
	var cartoon_list []CmfPortalPost = []CmfPortalPost{}
	controller.Model.DbManhua.Where("post_title IN (?)", post_title).Find(&cartoon_list)
	var res map[string]CmfPortalPost = map[string]CmfPortalPost{}
	for _, v := range cartoon_list {
		res[v.PostTitle] = v
	}
	return res
}
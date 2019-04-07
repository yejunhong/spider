package controller

import (
	"spider/model"
	"spider/lib"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
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
		var chapterContents []model.CartoonChapterContent =  controller.Model.GetContentsFindByChapterListUniqueId(bookInfo.UniqueId)
		fmt.Println(chapterContents)
	} else {
		go controller.ayncXiaoShuo(bookInfo, chapterList)
		go func(){
			var bookList = controller.Model.GetSqlCartoonListByNo(bookInfo.ResourceNo)
			fmt.Println(len(bookList))
			for _, v := range bookList {
				var chapterLists []model.CartoonChapter = controller.Model.GetChaptersFindByListUniqueId(v.UniqueId, 1)
				controller.ayncXiaoShuo(v, chapterLists)
			}
		}()
	}

	c.JSON(200, gin.H{
		"error": 0,
		"msg": "上传同步成功",
	})
}

type CmfSpiderPost struct {
	Id int64
	UniqueId string
}
// 同步小说
func(controller *Controller) ayncXiaoShuo(book model.CartoonList, chapter []model.CartoonChapter) {
	// CmfSpiderPost
	var bookData []map[string]interface{} = []map[string]interface{}{
		map[string]interface{}{
			"title": book.ResourceName, // 书籍名称
			"author": book.Author, // 作者
			"update_ok": 0, // '更新到正式环境状态;1:已更新,0:未更新',
			"summary": book.Detail, // '简介',
			"linkurl": book.ResourceUrl, // '源文章链接',
			"imgurl": book.ResourceImgUrl, // '封面图链接',
			"src_pid": 0, // '原站点对应的分类ID',
			// "update_num" int(10) NOT NULL DEFAULT '0' COMMENT '源站更新到第几则',
			// "src_name" varchar(20) DEFAULT NULL,
			"update_status": 0, // '状态;1:完结,0:更新中',
			"status": 0, // '状态;1:显示,0:不显示',
			"list_order": 0, // '排序',
			"remark": "", // '备注',
			"create_time": lib.Time(), // '创建时间',
			"update_time": lib.Time(), // '更新时间',
			"web_tag": book.ResourceNo, // '站点标识',
			"pid": 0, // 'post表ID',
			// "addition_count" int(11) DEFAULT NULL COMMENT '附加章节数，特殊需求字段',
			// "adult" int(11) DEFAULT NULL COMMENT '18X--1：是，0：否',
			"unique_id": book.UniqueId, // '数据同步唯一标识',
		},
	}
	model.DbBatchInsert(controller.Model.Db61, "cmf_spider_post", bookData, 
			[]string{"title", "author", "summary", "imgurl", "linkurl", "update_status"})

	var asyncBook CmfSpiderPost
	controller.Model.Db61.Where("unique_id = ?", book.UniqueId).Find(&asyncBook)

	var data []map[string]interface{}
	fmt.Println(len(chapter))
    for _, v := range chapter {
        data = append(data, map[string]interface{}{
            "pid": asyncBook.Id, // '对应的上级ID',
			"name": v.ResourceName, // '' COMMENT '章节名称',
			"update_ok": 0, // '更新到正式环境状态;1:已更新,0:未更新',
			"summary": book.Detail, //'摘要',
			"linkurl": v.ResourceUrl, // '' COMMENT '源文章链接',
			// "imgurl" varchar(250) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '封面图链接',
			// "src_pid" int(10) NOT NULL DEFAULT '0' COMMENT '原站点对应的分类ID',
			// "src_chapterid" int(10) NOT NULL DEFAULT '0' COMMENT '原站点对应的话ID',
			// "src_url" text COMMENT '源站图片连接',
			// "src_more" text,
			"status": 0, // '状态;1:显示,0:不显示',
			"more": v.Content, // '内容',
			"list_order": 0, // '排序',
			// "remark": lib.Time(), // varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
			"create_time": lib.Time(), // '创建时间',
			"update_time": lib.Time(), // '更新时间',
			"web_tag": v.ResourceNo, // '站点标识',
			"unique_id": v.UniqueId, // '数据同步唯一标识',
        })
    }
    if len(data) > 0 {
		model.DbBatchInsert(controller.Model.Db61, "cmf_spider_chapter", data, 
			[]string{"name", "linkurl", "summary", "more", "update_time"})
	}
	fmt.Println("同步结束")
}
/*

UPDATE cmf_spider_post p, (
	SELECT pid FROM cmf_spider_chapter where more like '%（完）%' GROUP BY pid
) s set p.update_status = 1 where p.id = s.pid
*/
type CmfPortalPost struct {
	Id int64
	UniqueId string
}
// 同步小说
func(controller *Controller) ayncPortalXiaoShuo(book model.CartoonList, chapter []model.CartoonChapter) {
	// CmfSpiderPost
	var bookData []map[string]interface{} = []map[string]interface{}{
		map[string]interface{}{
			`post_type`: 1, // '类型,1:文章;2:页面',
			`post_format`: 1, // '内容格式;1:html;2:md',
			`post_status`: 0, // '状态;1:已发布;0:未发布;',
			`comment_status`: 0, // '评论状态;1:允许;0:不允许',
			`create_time`: lib.Time(), // '创建时间',
			`update_time`: lib.Time(), // '更新时间',
			`chapter_update_time`: lib.Time(), // '章节更新时间',
			`post_title`: book.ResourceName, // 'post标题',
			`post_keywords`: "", // 'seo keywords',
			`post_excerpt`: book.Detail, // 'post摘要',
			`post_source`: "", //  '更新章节数',
			`post_content`: "", // '文章内容',
			`post_content_filtered`: "", // '处理过的文章内容',
			`more`: "", // '扩展属性,如缩略图;格式为json',
			`isfinish`: 0, // '写作进度是否完成 0连载中 1已完成',
			`isfree`: 0, // '是否免费 1免费 0收费',
			`post_tag`: 2, // '文章标识：1、漫画，2、小说',
			
			"title": book.ResourceName, // 书籍名称
			"author": book.Author, // 作者
			"update_ok": 0, // '更新到正式环境状态;1:已更新,0:未更新',
			"summary": book.Detail, // '简介',
			"linkurl": book.ResourceUrl, // '源文章链接',
			"imgurl": book.ResourceImgUrl, // '封面图链接',
			"src_pid": 0, // '原站点对应的分类ID',
			// "update_num" int(10) NOT NULL DEFAULT '0' COMMENT '源站更新到第几则',
			// "src_name" varchar(20) DEFAULT NULL,
			"update_status": 0, // '状态;1:完结,0:更新中',
			"status": 0, // '状态;1:显示,0:不显示',
			"list_order": 0, // '排序',
			"remark": "", // '备注',
			"create_time": lib.Time(), // '创建时间',
			"update_time": lib.Time(), // '更新时间',
			"web_tag": book.ResourceNo, // '站点标识',
			"pid": 0, // 'post表ID',
			// "addition_count" int(11) DEFAULT NULL COMMENT '附加章节数，特殊需求字段',
			// "adult" int(11) DEFAULT NULL COMMENT '18X--1：是，0：否',
			"unique_id": book.UniqueId, // '数据同步唯一标识',
		},
	}
	model.DbBatchInsert(controller.Model.Db61, "cmf_spider_post", bookData, 
			[]string{"title", "author", "summary", "imgurl", "linkurl", "update_status"})

	var asyncBook CmfSpiderPost
	controller.Model.Db61.Where("unique_id = ?", book.UniqueId).Find(&asyncBook)

	var data []map[string]interface{}
	fmt.Println(len(chapter))
    for _, v := range chapter {
        data = append(data, map[string]interface{}{
            "pid": asyncBook.Id, // '对应的上级ID',
			"name": v.ResourceName, // '' COMMENT '章节名称',
			"update_ok": 0, // '更新到正式环境状态;1:已更新,0:未更新',
			"summary": "", //'摘要',
			"linkurl": v.ResourceUrl, // '' COMMENT '源文章链接',
			// "imgurl" varchar(250) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '封面图链接',
			// "src_pid" int(10) NOT NULL DEFAULT '0' COMMENT '原站点对应的分类ID',
			// "src_chapterid" int(10) NOT NULL DEFAULT '0' COMMENT '原站点对应的话ID',
			// "src_url" text COMMENT '源站图片连接',
			// "src_more" text,
			"status": 0, // '状态;1:显示,0:不显示',
			"more": v.Content, // '内容',
			"list_order": 0, // '排序',
			// "remark": lib.Time(), // varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
			"create_time": lib.Time(), // '创建时间',
			"update_time": lib.Time(), // '更新时间',
			"web_tag": v.ResourceNo, // '站点标识',
			"unique_id": v.UniqueId, // '数据同步唯一标识',
        })
    }
    if len(data) > 0 {
		model.DbBatchInsert(controller.Model.Db61, "cmf_spider_chapter", data, 
			[]string{"name", "linkurl", "more", "update_time"})
	}
	fmt.Println("同步结束")
}
package main
 
import (
    "fmt"
    "spider/lib"
    "spider/model"
)

func main(){
    
    // lib.DonwloadFile("./1.png", "http://f2.kkmh.com/image/190305/0UX9Ax9Z0.webp-fe.w360.webp.m.i1")
    var config = lib.LoadConfig()
 
    var models *model.Model = &model.Model{
        Db: model.InitDb(config.Db_caiji.Host, config.Db_caiji.User, config.Db_caiji.Pass, config.Db_caiji.Name),
        // Db61: model.InitDb(config.Db_xiaoshuo.Host, config.Db_xiaoshuo.User, config.Db_xiaoshuo.Pass, config.Db_xiaoshuo.Name),
        // DbManhua: model.InitDb(config.Db_manhua.Host, config.Db_manhua.User, config.Db_manhua.Pass, config.Db_manhua.Name),
    }

    var no string = "C009"
    for i := 0 ; i < 4; i ++ {
        var book = models.GetCartoonListByNoStatus(no, -1)
        var bookName = Transfer(len(book) + 24)
        var list = models.GetChapterListByNoNum("C009", 350)
        var title string = "第" + bookName + "本"
        var ids []int64
        for _, v := range list {
            ids = append(ids, v.Id)
        }
        if len(ids) >= 350 {
            var uniqueId = lib.MD5(no + title)
            models.Db.Create(&model.CartoonList{
                ResourceNo: no,
                UniqueId: uniqueId,
                Tags: "激情",
                Author: "",
                Detail: "",
                Status: 0,
                ResourceUrl: "",
                ResourceName: title,
                ResourceImgUrl: "",
                DownloadImgUrl: "",
                BookType: 2,
                IsFree: 1,
                IsEnd: 1,
                Cdate: lib.Time(),
            })
            models.UpdateCartoonChapterByIds(ids, map[string]interface{}{"list_unique_id": uniqueId})
            fmt.Println(title, "：", uniqueId)
        }

    }
   
}

func Transfer(num int) string{
    chineseMap:=[]string{"","十", "百", "千", "万", "十", "百", "千", "亿", "十", "百", "千"}
    chineseNum:=[]string{"零", "一", "二","三","四","五","六","七","八","九"}
    listNum := []int{}
    for ;num >0;num = num/10{
        listNum= append(listNum, num%10)
    }
    n :=len(listNum)
    chinese :=""
    //注意这里是倒序的
    for i:=n-1; i>=0 ;i-- {
        chinese = fmt.Sprintf("%s%s%s", chinese, chineseNum[listNum[i]], chineseMap[i])
    }

    return chinese
}
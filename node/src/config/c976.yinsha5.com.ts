import url from 'url';
// print 打印
module.exports.Page = {
  name: '芒果书坊-小说',
  mobile: "",
  cookie: [
    {name: "_novelOpenid", value:"oFFzA514uRnqPLc908Y1Zwn8sizc", domain: "c976.yinsha5.com"},
  ],
  user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN",
}

// 获取书籍列表
// http://c976.yinsha5.com/index/ajax/getclassifybook.html?page=1&bookSex=0&bookStatus=0&pageSize=100&bookType='{"1":"奇幻·玄幻","2":"武侠·仙侠","4":"重生·穿越","5":"游戏·竞技","6":"科幻·灵异","8":"乡村·激情","9":"都市·言情","10":"历史·军事"}'
/*page	1
bookSex	0
bookStatus	0
bookType	
pageSize	10
<img src="http://img.phpyang.cn/uploads/5c0a31c267e02.jpg" alt="">
<div class="book-details">
<div class="book-title">他曾是一只哈士奇</div>
<div class="book-des">
    本是仙界大神，却不小心来到地球夺舍在了一条二十一世纪的哈士奇身上。
    繁华都市的灯红酒绿、
    诱人垂涎的美女主人、
    种种香艳、逗趣的经历交织在一起，上演一出美女与畜生的新奇特。
</div>
  <div class="book-other">
    <p class="author">
      <i class="iconfont icon-wode"></i>无楼
    </p>
    <p class="type">都市·言情</p>
  </div>
</div>
*/
module.exports.Book = {
  selector: 'div.content-top', // 列表选择器
  print: false,
  // selector: 'table.bookshelfListMain', // 列表选择器
  async handle(res: any): Promise<any> {
    const resdata = await res[0].$$eval('a', (e: any) => {
      let resdata: any = [];
      for(const v of e) {
        let author = v.querySelector('p.author').innerHTML;
        let tags = v.querySelector('p.type').innerHTML;
        tags = tags.split('·').join(",");
        author = author.split('i> ')[1];

        const resource_url = v.getAttribute("href").replace(/bookinfo/, "catalogue")
        resdata.push({
          resource_name: v.querySelector('div.book-title').innerHTML,
          resource_url: `http://c976.yinsha5.com${resource_url}`,
          resource_img_url: v.querySelector('img').getAttribute("src"),
          detail: v.querySelector('div.book-des').innerHTML,
          author: author,
          tags: tags,
        })
      }
      return resdata
    });
    return resdata
  },
  scroll: true, // 是否滚动页面
}

// 数据章节配置
/*
<a href="/index/book/bookcontent/chapeterid/934978.html?1554190185">
  <span class="chapter">第1章</span>
  <span class="is_free">
    <i class="iconfont free">免费</i>
  </span>
</a>
*/
module.exports.Chapter = {
  selector: 'div.catalog-list', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata = await res[0].$$eval('a', (e: any) => {
      let resdata: any = [];
      for(const v of e) {
        let is_free = "1";
        if(v.querySelector('span.is_free i').innerHTML == '免费'){
          is_free = "0";
        }
        resdata.push({
          is_free: is_free,
          resource_name: v.querySelector('span.chapter').innerHTML,
          resource_url: `http://c976.yinsha5.com${v.getAttribute("href")}`,
        })
      }
      return resdata
    });
    return resdata
  },
  scroll: true, // 是否滚动页面
}

module.exports.Content = {
  selector: 'body', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    resdata.push({
      resource_img_url: await res[0].$eval("div.chapterBox", e => e.innerHTML),
    })
    return resdata
  },
}
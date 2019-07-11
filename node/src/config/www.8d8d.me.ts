import url from 'url';
// print 打印
module.exports.Page = {
  name: '芒果书坊-小说',
  mobile: "",
}

// 获取书籍列表
module.exports.Book = {
  selector: '没有书籍概念', // 列表选择器
  print: false,
  // selector: 'table.bookshelfListMain', // 列表选择器
  async handle(res: any): Promise<any> {
    return [];
  },
  scroll: true, // 是否滚动页面
}

// 数据章节配置
/*
<div class="post-details">
  <div class="post-opts">
    <div class="post-date">
      <h1>13</h1>
      <h6>
      5月<br>
      2019 </h6>
    </div>
  </div>
  <h1><a href="/blog-detail.php?target=text&amp;no=10335" class="colr">遊泳池vs姐姐</a></h1>
  <div class="clear"></div>
<div class="post-content">
  <p>我是一個即將邁入高中的學生，而現在正是炎炎夏日，窮極無聊的暑假，由於要上新高中，所以根本沒有課業壓力，也很少人約我出門逛街、看電影之類的休閒活動，因此我把很多時間都拿來運動蜢，遊泳－就是我的娛樂，而這暑假每天遊泳，也開啟了我性愛的人生。</p><p>我家附近就有一間不算大的社區泳池，票價也是十分便宜，而我因為怕人多，所以都趁早上或晚上的時段去遊個幾小時。</p> </div>
  <div style="margin-top:20px">
    <a href="/blog-detail.php?target=text&amp;no=10335" class="readmore">繼續閱讀</a>
    <span style="float:right">分類：<a href="/index.php?target=text&amp;no=21">亂倫文學</a></span>
    <div class="clear"></div>
  </div>
</div>
*/
module.exports.Chapter = {
  selector: 'div#showblog_area', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata = await res[0].$$eval('div.post-details', (e: any) => {
      let resdata: any = [];
      for(const v of e) {
        const a = v.querySelector('h1 a');
        let detail = "";
        if (v.querySelector('div.post-content p')) {
          detail = v.querySelector('div.post-content p').innerHTML;
        }
        resdata.push({
          is_free: "0",
          detail: detail,
          resource_name: a.innerHTML,
          resource_url: `https://www.8d8d.me${a.getAttribute("href")}`,
        })
      }
      return resdata
    });
    // console.log(resdata)
    return resdata;
  },
  // 爬取 下一页数据
  next: {
    selector: 'body', // 列表选择器
    async handle(e: any, urlStr: string){
      const urls = url.parse(urlStr, true)
      urls.query.pageNum_BlogList = (parseInt(urls.query.pageNum_BlogList.toString()) + 1).toString()
      const nextUrl = url.format({
                      protocol: urls.protocol,
                      host: urls.host,
                      query: urls.query,
                      pathname: urls.pathname,
                    });
      return nextUrl
    }
  },
}

module.exports.Content = {
  selector: 'body', // 列表选择器
  print: false,
  // browser_request: false, // 是通过浏览器请求
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    resdata.push({
      resource_img_url: await res[0].$eval("div.post-content", (e: any) => {
        return e.innerHTML.replace(/<img(.*?)>/, '').replace(/<img(.*?)>/, '');
      }),
    })
    return resdata
  },
}
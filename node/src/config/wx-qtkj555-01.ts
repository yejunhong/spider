import url from 'url';

module.exports.Page = {
  name: '谦棠科技-漫画',
  mobile: "",
  cookie: [
    /*CNZZDATA1260292476=38577323-1553934597-%7C1553934597; 
    UM_distinctid=169ce00b0d412d-0fc3f772ab10708-7e145f62-4a574-169ce00b0d5b5; 
    ASPSESSIONIDQCRRDQCT=ILJMIPDDOMCGOOLGNFCCLHNN; 
    sid=6807933; 
    token=3300330093007300030083006300160057008600E6001600D600*/
    {name: "sid", value:"6807933", domain: "c1021.w406.s1388630.ririyue.cn"},
    {name: "token", value: "3300330093007300030083006300160057008600E6001600D600", domain: "c1021.w406.s1388630.ririyue.cn"},
    {name: "UM_distinctid", value: "169ce00b0d412d-0fc3f772ab10708-7e145f62-4a574-169ce00b0d5b5", domain: "c1021.w406.s1388630.ririyue.cn"}
  ],
  user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN",
}

// 获取书籍列表
module.exports.Book = {
  selector: 'body', // 列表选择器
  // selector: 'table.bookshelfListMain', // 列表选择器
  async jsonHandle(res: any): Promise<any> {
    // const json = await res.$eval()
    const resdata: any = [];
    res.data.map((v: any) => {
      resdata.push({
        resource_name: v.title,
        resource_url: `http://c1021.w406.s4694780.5fmj.com.cn/manhua/info.html?id=${v.bid}`,
        resource_img_url: v.litpicd,
        detail: v.description,
        author: v.author
      })
    })
    return resdata
  },
  // 爬取 下一页数据
  next: {
    selector: 'body', // 列表选择器
    async handle(e: any, urlStr: string){
      const urls = url.parse(urlStr, true)
      urls.query.page = (parseInt(urls.query.page.toString()) + 1).toString()
      const nextUrl = url.format({
                      protocol: urls.protocol,
                      host: urls.host,
                      query: urls.query,
                      pathname: urls.pathname,
                    });
      return nextUrl
    }
  },
  // scroll: true, // 是否滚动页面
}

// 数据章节配置
module.exports.chapter = {
  selector: 'div.titleDiv', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      let is_free = "1";
      if(await e.Attr('.needBuyDiv img', 'class') == 'd0'){
        is_free = "0";
      }
      resdata.push({
        is_free: is_free,
        resource_name: await e.Html('.title'),
        resource_url: "http://c1021.w406.s4694780.5fmj.com.cn/manhua/" + await e.Attr('a', 'href'),
      })
    }
    return resdata
  },
}

module.exports.content = {
  selector: 'div.readMain img', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        resource_img_url: await e.getAttribute('data-original'),
      })
    }
    return resdata
  },
}
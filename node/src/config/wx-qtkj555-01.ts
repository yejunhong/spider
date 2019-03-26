
module.exports.Page = {
  name: '谦棠科技-漫画',
  mobile: "",
  cookie: [
    {name: "sid", value:"6807933", domain: "c1021.w406.s4694780.5fmj.com.cn"},
    {name: "token", value: "3300330093007300030083006300160057008600E6001600D600", domain: "c1021.w406.s4694780.5fmj.com.cn"},
    {name: "UM_distinctid", value: "1691805e44257f-0d1ddf426e1d63-7e145f62-4a574-1691805e4434f4", domain: "c1021.w406.s4694780.5fmj.com.cn"}
  ],
  user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN",
}

// 获取书籍列表
module.exports.Book = {
  selector: 'tr.bookshelfListMainTr', // 列表选择器
  async handle (res, Element): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      const resource_url = v.getAttribute("onclick");
      resdata.push({
        resource_name: e.Html('.bookshelfListMainTitle'),
        resource_url: "http://c1021.w406.s4694780.5fmj.com.cn/manhua/" + resource_url.split('"')[1],
        resource_img_url: e.Attr('.bookshelfListMainImg img', 'src'),
        detail: e.Html('.bookshelfListMainDes'),
      })
    }
    return resdata
  },
  scroll: true, // 是否滚动页面
}

// 数据章节配置
module.exports.chapter = {
  selector: 'div.titleDiv', // 列表选择器
  async handle (res, Element): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      let is_free = "1";
      if(e.Attr('.needBuyDiv img', 'class') == 'd0'){
        is_free = "0";
      }
      resdata.push({
        is_free: is_free,
        resource_name: e.Html('.title'),
        resource_url: "http://c1021.w406.s4694780.5fmj.com.cn/manhua/" + e.Attr('a', 'href'),
      })
    }
    return resdata
  },
}

module.exports.content = {
  selector: 'div.readMain img', // 列表选择器
  async handle (res, Element): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        resource_img_url: e.getAttribute('data-original'),
      })
    }
    return resdata
  },
}
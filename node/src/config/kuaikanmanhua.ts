
module.exports.Page = {
  name: '快看漫画',
  mobile: "",
  cookie: "",
  user_agent: "",
}

module.exports.Login = {
  url: 'https://www.kuaikanmanhua.com/webs/loginh',
  user: {selector: 'div.inputPhone input', value: '18320777006'},
  pass: {selector: 'div.password div.inputBox input', value: 'junchun153259'},
  click: 'div.submit input.loginhBtn', // 
  async handle (res: any): Promise<any> { // 处理数据
    // loginhBtn//
  },
}

// 获取书籍列表
module.exports.Book = {
  islogin: false,
  selector: 'div.ItemSpecial', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        tags: await e.Html('span.itemTitle'),
        detail: await e.Html('span.itemTitle'),
        resource_name: await e.Html('span.itemTitle'),
        resource_url: await e.Attr('a', 'href'),
        resource_img_url: await e.Attr('.img', 'data-src'),
        author: await e.Html('p .author')
      })
    }
    return resdata
  },
  // 爬取 下一页数据
  next: {
    selector: 'ul.pagination', // 列表选择器
    async handle(e: any){
      let maxPage = await e.$$eval('li.itemBten a', e => e[e.length - 1].innerHTML)
      let page = await e.$eval('li.active a', e => e.innerHTML)
      if(parseInt(maxPage) > parseInt(page)){
        return `https://www.kuaikanmanhua.com/tag/0?state=1&page=${(parseInt(page) + 1)}`
      }
      return false
    }
  },
  scroll: false,
}

// 数据章节配置
module.exports.Chapter = {
  islogin: false,
  selector: 'div.article-list table tr', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        is_free: await e.IsExist('i.ico-lockoff'),
        resource_name: await e.Attr('.article-img', 'title'),
        resource_url: await e.Attr('.article-img', 'href'),
        resource_img_url: await e.Attr('.kk-sub-img', 'src'),
      })
    }
    return resdata
  },
}

module.exports.Content = {
  islogin: false,
  selector: 'div.imgList img', // 列表选择器
  async handle (res: any, Element: any): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        resource_img_url: await e.getAttribute('data-src'),
      })
    }
    return resdata
  },
}
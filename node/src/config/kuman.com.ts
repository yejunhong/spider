import { fchmod } from "fs";

module.exports.Page = {
  name: '酷漫网',
  mobile: "",
  cookie: "",
  user_agent: ""
}

// 获取书籍列表
module.exports.Book = {
  selector: 'div.list_mh ul', // 列表选择器
  async handle (res, Element): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      const tags = await v.$$eval('li.biaoqian a', res => {
        let tags: any = [];
        for (const e of res) {
          tags.push(e.innerHTML);
        }
        tags = tags.join(",");
        return tags
      });
      // console.log(tags)
      resdata.push({
        tags: tags,
        author: await e.Html('li.zuozhe'),
        resource_name: await e.Attr('li.title a', 'title'),
        resource_url: await e.Attr('li.title a', 'href'),
        resource_img_url: await e.Attr('li.pic a img', 'src'),
        detail: await e.Html('li.info'),
      })
    }
    return resdata
  },
  // 爬取 下一页数据
  next: {
    selector: 'div.page', // 列表选择器
    async handle(e: any, urlStr: string){
      let page = await e.$eval('a.mun', (e: any) => e.innerHTML)
      let maxPage = await e.$$eval('a', (ele: any) => {
        return ele[ele.length-1].innerHTML
      })
      if(parseInt(maxPage) > parseInt(page)){
        return `http://www.kuman.com/all/list---------${(parseInt(page) + 1)}-1/`
      }
      return false
    }
  },
}

// 数据章节配置
module.exports.Chapter = {
  selector: 'div#play_0 ul li', // 列表选择器
  async handle (res, Element): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      let is_free = "1";
      if(await e.IsExist('p em')){
        is_free = "0";
      }
      resdata.push({
        is_free: is_free,
        resource_name: await e.Attr('p a', 'title'),
        resource_url: await e.Attr('p a', 'href'),
      })
    }
    return resdata
  }
}

module.exports.Content = {
  selector: 'div.show_list ul li', // 列表选择器
  async handle (res, Element): Promise<any> { // 处理数据
    const resdata: any = [];
    for(const v of res){
      const e = new Element(v);
      resdata.push({
        resource_img_url: await e.Attr('span img', 'src'),
      })
    }
    return resdata
  },
}

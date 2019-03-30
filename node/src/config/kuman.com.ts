
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
      const tags = v.$$eval('li.biaoqian a', res => {
        let tags: any = [];
        for (const e of res) {
          tags.push(e.innerHTML);
        }
        tags = tags.join(",");
        return tags
      });
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
  }
}

// 数据章节配置
module.exports.chapter = {
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

module.exports.content = {
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

// print 打印
module.exports.Page = {
  name: '歪歪漫画',
  mobile: "",
  cookie: [
    {name: 'timestamp', value: '53749965702593295734', domain: "m.titi004.com"},
    {name: 'ciu_key', value: '541a7bd9-627e-45f4-af5c-ab1a90080eb0$120.229.74.175', domain: "m.titi004.com"},
    {name: 'JSESSIONID', value: 'm7t924j24u4g1tuaclmuzdp3q', domain: "m.titi004.com"},
    {name: "ticket", value:"4eab780d-cc2f-4c8f-9563-27bb3969fd18", domain: "m.titi004.com"},
  ],
}

// 获取书籍列表
module.exports.Book = {
  selector: 'pre', // 列表选择器
  // print: true,
  // selector: 'table.bookshelfListMain', // 列表选择器
  async jsonHandle(res: any): Promise<any> {
    // const json = await res.$eval()
    const resdata: any = [];
    res.content.list.map((v: any) => {
      resdata.push({
        tags: v.category,
        resource_name: v.name,
        resource_url: `https://m.titi004.com/query/book/directory?bookId=${v.id}`,
        resource_img_url: v.coverUrl,
        detail: v.description,
        author: v.author
      })
    })
    return resdata
  },
  // scroll: true, // 是否滚动页面
}

// 数据章节配置
module.exports.Chapter = {
  selector: 'body pre', // 列表选择器
  // print: true,
  async jsonHandle (res: any): Promise<any> { // 处理数据
    const resdata: any = [];
    let order: number = 0;
    res.content.map((v: any) => {
      let is_free = "1";
      // console.log(v.freeFlag)
      if(v.freeFlag){
        is_free = "0";
      }
      order = order + 1;
      resdata.push({
        sort: `${order}`,
        is_free: is_free,
        resource_name: v.title,
        resource_img_url: v.coverUrl,
        resource_url: `https://m.titi004.com/query/book/chapter?bookId=${v.bookId}&chapterId=${v.id}`,
      })
    })
    return resdata
  },
}

module.exports.Content = {
  selector: 'body pre', // 列表选择器
  async jsonHandle (res: any): Promise<any> { // 处理数据
    const resdata: any = [];
    res.content.imageList.map((v: any) => {
      resdata.push({
        resource_img_url: v.url,
      })
    })
    return resdata
  },
}
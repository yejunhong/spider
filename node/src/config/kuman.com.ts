
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
        author: e.Html('li.zuozhe'),
        resource_name: e.Attr('li.title a', 'title'),
        resource_url: e.Attr('li.title a', 'href'),
        resource_img_url: e.Attr('li.pic a img', 'src'),
        detail: e.Html('li.info'),
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
        resource_img_url: e.Attr('span img', 'src'),
      })
    }
    return resdata
  },
}
// http://c1021.w406.s236341.5fmj.com.cn/manhua/
module.exports = {
  name: '酷漫网',
  // cookie 信息 Referer: http://c1021.w406.s4694780.5fmj.com.cn/manhua/
  // 伪造浏览器
  // user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN",
  // jquery: false,
  cookie: false,
  login: false,
  list: { // 漫画列表
    selector: 'div.list_mh ul', // 列表选择器
    datas: 'GetListData', // 对应当前配置文件 function
    next: ['div.page a', 'Next'], // 下一页
    // scroll: true, // 滚动操作 异步加载信息
  }, 
  chapter: { // 漫画章节
    selector: 'div#play_0 ul li', // 列表选择器
    detail: ['div.xinxi span', 'GetChapterDetail'], // 获取章节详情
    datas: 'GetChapterData', // 对应当前配置文件 function
  },
}
/*
function Next(pages){
  let href = ""
  if(pages){
    let p = false
    for (const e of pages) {
      // console.log(e.getAttribute("class"))
      if(p){
        href = e.getAttribute("href")
        break
      }
      if(e.getAttribute("class") == 'mun'){
        p = true
      }
    }
  }
  return href
}

function GetChapterDetail(e){
  let is_end = "0";
  if(e[1].querySelector('font').innerHTML == '完结'){
    is_end = "1";
  }
  // console.log()
  return {
    is_end: is_end, //
  }
}
*/
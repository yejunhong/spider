
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
  chapter_content: {// 漫画章节-内容
    selector: 'div.show_list ul li', // 列表选择器
    datas: 'GetChapterContentData', // 对应当前配置文件 function
  }
}

function GetListData(e) {
  const t = e.querySelectorAll('li.biaoqian a');
  
  let tags = [];
  if(t){
    for (const e of t) {
      tags.push(e.innerHTML);
    }
    tags = tags.join(",");
  }
  return {
    tags: tags,
    author: e.querySelector('li.zuozhe').innerHTML,
    resource_name: e.querySelector('li.title a').getAttribute("title"),
    resource_url: e.querySelector('li.title a').getAttribute("href"),
    resource_img_url: e.querySelector('li.pic a img').getAttribute("src"),
    detail: e.querySelector('li.info').innerHTML,
  };
}

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

function GetChapterData(e) {
  let is_free = "1";
  if(e.querySelector('p em') == undefined){
    is_free = "0";
  }
  // console.log(e.querySelector('.title').innerHTML)
  return {
    is_free: is_free,
    resource_name: e.querySelector('p a').getAttribute('title'),
    resource_url: e.querySelector('p a').getAttribute('href'),
  };
}

/*
<span><img src="http://s2.static.kuman.com//001/00/56/61/16876/66885bfba317ad0d0.jpeg!km2" data-image_id="10178" id="image_10178"></span>
*/
function GetChapterContentData(e) {
  return {
    resource_img_url: e.querySelector('span img').getAttribute('src'),
  };
}


// http://c1021.w406.s236341.5fmj.com.cn/manhua/

module.exports = {
  name: '谦棠科技-漫画',
  console: true, // 是否开启 console.log() 控制台
  // cookie 信息
  cookie: [
    {name: "sid", value:"6807933", domain: "c1021.w406.s236341.5fmj.com.cn"},
    {name: "token", value: "3300330093007300030083006300160057008600E6001600D600", domain: "c1021.w406.s236341.5fmj.com.cn"},
    {name: "UM_distinctid", value: "1691805e44257f-0d1ddf426e1d63-7e145f62-4a574-1691805e4434f4", domain: "c1021.w406.s236341.5fmj.com.cn"}
  ],
  // 伪造浏览器
  user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN",
  // jquery: false,
  login: false,
  list: { // 漫画列表
    selector: 'div.similar div.item', // 列表选择器
    datas: 'get_list_data', // 对应当前配置文件 function
  },
  chapter: { // 漫画章节
    selector: 'div.article-list table tr', // 列表选择器
    datas: 'get_chapter_data', // 对应当前配置文件 function
  },
  chapter_content: {// 漫画章节-内容
    selector: 'div.ItemSpecial', // 列表选择器
    datas: 'get_chapter_content_data', // 对应当前配置文件 function
  }
}

function get_list_data(e) {
  return {
    resource_name: e.querySelector('div.txt h4').innerHTML,
    resource_url: e.querySelector('img').getAttribute("src"),
    detail: e.querySelector('div.txt span').innerHTML,
  };
}

function get_chapter_data(e) {
  const res = e.querySelector('.article-img');
  return {
    is_free: e.querySelector('i.ico-lockoff'),
    resource_name: res.getAttribute('title'),
    resource_url: res.getAttribute('href'),
    resource_img_url: e.querySelector('.kk-sub-img').getAttribute('src'),
  };
}


function get_chapter_content_data(e) {
  return {
    tags: e.querySelector('span.itemTitle').innerHTML,
    detail: e.querySelector('span.itemTitle').innerHTML,
    resource_name: e.querySelector('span.itemTitle').innerHTML,
    resource_url: e.querySelector('a').getAttribute('href'),
    resource_img_url: e.querySelector('.img').getAttribute('data-src'),
    author: e.querySelector('p .author').innerHTML
  };
}

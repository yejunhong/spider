
// http://c1021.w406.s236341.5fmj.com.cn/manhua/

module.exports = {
  name: '谦棠科技-漫画',
  console: true, // 是否开启 console.log() 控制台
  // cookie 信息 Referer: http://c1021.w406.s4694780.5fmj.com.cn/manhua/
  cookie: [
    {name: "sid", value:"6807933", domain: ""},
    {name: "token", value: "3300330093007300030083006300160057008600E6001600D600", domain: ""},
    {name: "UM_distinctid", value: "1691805e44257f-0d1ddf426e1d63-7e145f62-4a574-1691805e4434f4", domain: ""}
  ],
  // 伪造浏览器
  user_agent: "Mozilla/5.0 (Linux; Android 4.4.4; HM NOTE 1LTEW Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36 MicroMessenger/6.0.0.54_r849063.501 NetType/WIFI",
  // jquery: false,
  login: false,
  list: { // 漫画列表
    selector: 'tr.bookshelfListMainTr', // 列表选择器
    datas: 'get_list_data', // 对应当前配置文件 function
    scroll: true, // 滚动操作 异步加载信息
  },
  chapter: { // 漫画章节
    selector: 'div.titleDiv', // 列表选择器
    datas: 'get_chapter_data', // 对应当前配置文件 function
  },
  chapter_content: {// 漫画章节-内容
    selector: 'div.ItemSpecial', // 列表选择器
    datas: 'get_chapter_content_data', // 对应当前配置文件 function
  }
}

function get_list_data(e) {
  const resource_url = e.getAttribute("onclick");
  return {
    resource_name: e.querySelector('.bookshelfListMainTitle').innerHTML,
    resource_url: "http://c1021.w406.s4694780.5fmj.com.cn/manhua/" + resource_url.split('"')[1],
    resource_img_url: e.querySelector('.bookshelfListMainImg img').getAttribute("src"),
    detail: e.querySelector('.bookshelfListMainDes').innerHTML,
  };
}

/**
<div class="titleDiv" id='2922828'  num='5' >
  <a href='reader.html?chapter_id=2922828&bid=67894'>
  <div class='title already'>第5话</div>
  <div class='needBuyDiv' cid='2922828'>
    <img  src="images/needBuy.png"  class="d0">
  </div>
  </a>
</div>
 */
function get_chapter_data(e) {
  let is_free = 1;
  if(e.querySelector('.needBuyDiv img').getAttribute('class') == 'd0'){
    is_free = 0;
  }
  console.log(e)
  return {
    is_free: is_free,
    resource_name: e.querySelector('.title').innerHTML,
    resource_url: "http://c1021.w406.s4694780.5fmj.com.cn/manhua/" + e.querySelector('a').getAttribute('href'),
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


// http://c1021.w406.s236341.5fmj.com.cn/manhua/

module.exports = {
  name: '谦棠科技-漫画',
  // cookie 信息 Referer: http://c1021.w406.s4694780.5fmj.com.cn/manhua/books_type.html
  // 伪造浏览器
  user_agent: "Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN",
  // jquery: false,
  cookie: [
    {name: "sid", value:"6807933", domain: "c1021.w406.s4694780.5fmj.com.cn"},
    {name: "token", value: "3300330093007300030083006300160057008600E6001600D600", domain: "c1021.w406.s4694780.5fmj.com.cn"},
    {name: "UM_distinctid", value: "1691805e44257f-0d1ddf426e1d63-7e145f62-4a574-1691805e4434f4", domain: "c1021.w406.s4694780.5fmj.com.cn"}
  ],
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
    selector: 'div.readMain img', // 列表选择器
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
  let is_free = "1";
  if(e.querySelector('.needBuyDiv img').getAttribute('class') == 'd0'){
    is_free = "0";
  }
  // console.log(e.querySelector('.title').innerHTML)
  return {
    is_free: is_free,
    resource_name: e.querySelector('.title').innerHTML,
    resource_url: "http://c1021.w406.s4694780.5fmj.com.cn/manhua/" + e.querySelector('a').getAttribute('href'),
  };
}

function get_chapter_content_data(e) {
  return {
    resource_img_url: e.getAttribute('data-original'),
  };
}

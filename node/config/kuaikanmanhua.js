
module.exports = {
  name: '快看漫画',
  console: true, // 是否开启 console.log() 控制台
  jquery: false,
  list: { // 漫画列表
    url: 'https://www.kuaikanmanhua.com/tag/0',
    selector: 'div.ItemSpecial', // 列表选择器
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
    tags: e.querySelector('span.itemTitle').innerHTML,
    detail: e.querySelector('span.itemTitle').innerHTML,
    resource_name: e.querySelector('span.itemTitle').innerHTML,
    resource_url: e.querySelector('a').getAttribute('href'),
    resource_img_url: e.querySelector('.img').getAttribute('data-src'),
    author: e.querySelector('p .author').innerHTML
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

const puppeteer = require('puppeteer');

const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};


const kuaikanmanhua_config = {
  name: '快看漫画',
  list: {
    url: 'https://www.kuaikanmanhua.com/tag/0',
    class: 'div.ItemSpecial', // 通过class 获取列表内容
    data (e) => {
      return {
        tags: ['span.itemTitle', 'innerHTML'],
        detail: e.querySelector('span.itemTitle').innerHTML,
        resource_name: e.querySelector('span.itemTitle').innerHTML,
        resource_url: e.querySelector('a').getAttribute('href'),
        resource_img_url: e.querySelector('.img').getAttribute('data-src'),
        author: e.querySelector('p .author').innerHTML
      }
    }
  }
}


puppeteer.launch(options).then(async browser => {
  const page = await browser.newPage();
  await page.goto('https://www.kuaikanmanhua.com/tag/0');
  const list = await page.$$eval('div.ItemSpecial', (d) => {
    return d.map(e => {
      return {
        unique_id: '',
        tags: e.querySelector('span.itemTitle').innerHTML,
        detail: e.querySelector('span.itemTitle').innerHTML,
        resource_name: e.querySelector('span.itemTitle').innerHTML,
        resource_url: e.querySelector('a').getAttribute('href'),
        resource_img_url: e.querySelector('.img').getAttribute('data-src'),
        author: e.querySelector('p .author').innerHTML
      }
    });
  });
  console.log(list)
  // 其他操作...
  await browser.close();
});

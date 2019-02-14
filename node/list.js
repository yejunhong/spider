const puppeteer = require('puppeteer');
const crypto = require('crypto');

const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};

const kuaikanmanhua_config = {
  name: '快看漫画',
  list: {
    url: 'https://www.kuaikanmanhua.com/tag/0',
    selector: 'div.ItemSpecial', // 通过class 获取列表内容
    data: {
      tags: ['span.itemTitle', 'href'] // e.querySelector('a').getAttribute('href') 获取a选择器src属性
    }
  }
}

const config = kuaikanmanhua_config;

puppeteer.launch(options).then(async browser => {
  const page = await browser.newPage();
  await page.goto(config.list.url);
  page.on('console', msg => console.log(msg.text()));
  const list = await page.$$eval(config.list.selector, async (d, config) => {
    return d.map(e => {
      return {
        tags: e.querySelector('span.itemTitle').innerHTML,
        detail: e.querySelector('span.itemTitle').innerHTML,
        resource_name: e.querySelector('span.itemTitle').innerHTML,
        resource_url: e.querySelector('a').getAttribute('href'),
        resource_img_url: e.querySelector('.img').getAttribute('data-src'),
        author: e.querySelector('p .author').innerHTML
      }
    });
  }, config);

  console.log(list)
  // 其他操作...
  await browser.close();
});
/*
crypto.createHash('md5').update(text).digest('hex')
*/


const puppeteer = require('puppeteer');

const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};
const createPuppeteer = puppeteer.launch(options);

class ClassName {

  

}

puppeteer.launch(options).then(async browser => {
  const page = await browser.newPage(); // 打开一个新标签
  await page.goto('https://www.baidu.com');
  await page.screenshot({path: 'yqq.png'});
  // 其他操作...
  await browser.close();
});

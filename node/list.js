const puppeteer = require('puppeteer');
const crypto = require('crypto');

const arguments = process.argv.splice(2);
const config_name = arguments[0]
const config = require(`./config/${config_name}`);

console.log(arguments)

const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};

puppeteer.launch(options).then(async browser => {
  const page = await browser.newPage();
  await page.goto(config.list.url);
  page.on('console', msg => console.log(msg.text()));
  await page.exposeFunction('md5', text =>
    crypto.createHash('md5').update(text).digest('hex')
  );
  //await page.exposeFunction('res_data', e => 1);
  // 注入配置信息
  await page.addScriptTag({
    path: `config/${config_name}.js`,
  });

  const list = await page.$$eval(config.list.selector, async (d, config) => {
    let res = [];
    for (let e of d) {
      const r = eval(`${config.list.datas}(e)`)
      res.push(r)
    }
    return res
  }, config);

  console.log(list)
  // 其他操作...
  // await browser.close();
});
/*
crypto.createHash('md5').update(text).digest('hex')
*/

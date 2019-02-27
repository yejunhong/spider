const puppeteer = require('puppeteer');
const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};

(async function () {

  const browser = await puppeteer.launch(options);
  const page = await browser.newPage();

  // chapter_id=2921307&bid=67862
  await page.setCookie({name: "sid", value:"6807933", domain: "c1021.w406.s4694780.5fmj.com.cn"});     // 设置cookie
  await page.setCookie({name: "token", value: "3300330093007300030083006300160057008600E6001600D600", domain: "c1021.w406.s4694780.5fmj.com.cn"});     // 设置cookie
  await page.setCookie({name: "UM_distinctid", value: "1691805e44257f-0d1ddf426e1d63-7e145f62-4a574-1691805e4434f4", domain: "c1021.w406.s4694780.5fmj.com.cn"});     // 设置cookie

  await page.setUserAgent("Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN");
  
  await page.goto("http://c1021.w406.s4694780.5fmj.com.cn/manhua/reader.html?chapter_id=2903018&bid=67519");
  
  const res = await page.$$eval('div.similar div.item', async (d) => {
    return [{a: 1}, {a: 2}]; // 登录结果
  });
  await page.close(); // 关闭当前标签页
  console.log(res)
})();
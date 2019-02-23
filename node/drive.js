const puppeteer = require('puppeteer');
const crypto = require('crypto');
const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};

class Drive{

  /**
   * 创建浏览器
   */
  async CreateBrowser() {
    const browser = await puppeteer.launch(options);
  }

  /**
   * 爬取 列表
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  CrawlList(call, callback){
    const request = call.request
    const config = require(`./config/${request.config_name}`);
    await this.Login(); // 登录操作
    const resData = await this.OpenPage(request.url, request.config_name, config.list);
    callback(null, {data: resData, next: ''});
  }

  /**
   * 爬取 列表-章节
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  async CrawlChapter(call, callback){
    const request = call.request;
    const config = require(`./config/${request.config_name}`);
    await this.Login(); // 登录操作
    const resData = await this.OpenPage(request.url, request.config_name, config.chapter);
    callback(null, {data: resData, next: ''});
  }

  /**
   * 爬取 章节-内容
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  async CrawlChapterContent(call, callback){
    const request = call.request;
    const config = require(`./config/${request.config_name}`);
    await this.Login(); // 登录操作
    const resData = await this.OpenPage(request.url, request.config_name, config.chapter_content);
    callback(null, {data: resData, next: ''});
  }

  /**
   * 登录操作 
   * @param url // 访问http-url
   * @param config_name // 配置名称, 用于标签页面加载js配置文件
   * @param config // 配置信息
   */
  async Login(url, config_name, config){
    if(config.login == false){ // 是否进行登录操作
      return
    }
    const page = await this.browser.newPage();
    await page.goto(url);
    // 注入配置信息
    await page.addScriptTag({
      path: `config/${config_name}.js`,
    });
    // 待定过期
    await page.waitForNavigation({timeout: 1000});
    const res = await page.$$eval(config.selector, async (d, config) => {
      const res = eval(`${config.login}(e)`);
      return res; // 登录结果
    }, config);
    await page.close(); // 关闭当前标签页
    return res

  }

  /**
   * 打开一个标签页面
   * @param url // 访问http-url
   * @param config_name // 配置名称, 用于标签页面加载js配置文件
   * @param config // 配置信息
   */
  async OpenPage(url, config_name, config) {
    const page = await browser.newPage();
    
    if (config.cookie != false){
      for (let e of d) {
        await page.setCookie(e);     // 设置cookie
      }
    }

    // 伪造浏览器
    if (config.user_agent != false){
      await page.setUserAgent(config.user_agent);
    }
   
    await page.goto(url);
    page.on('console', msg => console.log(msg.text()));
    // 注入函数到浏览器
    await page.exposeFunction('md5', text =>
      crypto.createHash('md5').update(text).digest('hex')
    );
    //await page.exposeFunction('res_data', e => 1);
    // 注入配置信息
    await page.addScriptTag({
      path: `config/${config_name}.js`,
    });

    const list = await page.$$eval(config.selector, async (d, config) => {
      let res = {
        data: [], 
        next: eval(`${config.datas}(e)`) // 下一页url
      };
      for (let e of d) {
        const r = eval(`${config.datas}(e)`);
        res.data.push(r);
      }
      return res;
    }, config);

    await page.close(); // 关闭当前标签页

    return list;
  }

}
module.exports = Drive;
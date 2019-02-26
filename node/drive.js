const puppeteer = require('puppeteer');
const devices = require('puppeteer/DeviceDescriptors');
const crypto = require('crypto');
const options = {args: ['--no-sandbox', '--disable-setuid-sandbox']};

let browser = {};
const iPhone = devices['iPhone X'];

class Drive{
  
  /**
   * 创建浏览器
   */
  async CreateBrowser() {
    browser = await puppeteer.launch(options);
    return this;
  }

  /**
   * 爬取 列表
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  async CrawlList(call, callback){
    const request = call.request
    const config = require(`./config/${request.config_name}`);
    await this.Login(config); // 登录操作
    const resData = await this.OpenPage(request.url, request.config_name, config, config.list);
    callback(null, resData);
  }

  /**
   * 爬取 列表-章节
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  async CrawlChapter(call, callback){
    const request = call.request;
    const config = require(`./config/${request.config_name}`);
    await this.Login(config); // 登录操作
    const resData = await this.OpenPage(request.url, request.config_name, config, config.chapter);
    callback(null, resData);
  }

  /**
   * 爬取 章节-内容
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  async CrawlChapterContent(call, callback){
    const request = call.request;
    const config = require(`./config/${request.config_name}`);
    await this.Login(config); // 登录操作
    const resData = await this.OpenPage(request.url, request.config_name, config, config.chapter_content);
    callback(null, resData);
  }

  /**
   * 登录操作 
   * @param url // 访问http-url
   * @param config_name // 配置名称, 用于标签页面加载js配置文件
   * @param config // 配置信息
   */
  async Login(config){
    if(config.login == false){ // 是否进行登录操作
      return
    }
    const page = await browser.newPage();
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
  async OpenPage(url, config_name, config, config_eval) {
 
    const page = await browser.newPage();
    
    if (config.cookie != undefined && config.cookie != false){
      for (let e of config.cookie) {
        await page.setCookie(e);     // 设置cookie
      }
    }

    if (config_eval.cookie != undefined && config_eval.cookie != false){
      for (let e of config_eval.cookie) {
        await page.setCookie(e);     // 设置cookie
      }
    }

    // 伪造浏览器
    if (config.user_agent != false){
      await page.setUserAgent(config.user_agent);
    }
    // await page.emulate(iPhone);
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

    // 滚动操作
    if (config_eval.scroll != undefined && config_eval.scroll != false){
      let preScrollHeight = 0;
      let scrollHeight = -1;
      while(preScrollHeight !== scrollHeight) {
        // 详情信息是根据滚动异步加载，所以需要让页面滚动到屏幕最下方，通过延迟等待的方式进行多次滚动
        let scrollH1 = await page.evaluate(async () => {
          let h1 = document.body.scrollHeight;
          window.scrollTo(0, h1);
          return h1;
        });
        
        await page.waitFor(500);
        let scrollH2 = await page.evaluate(async () => {
          return document.body.scrollHeight;
        });

        console.log(`滚动页面高度：${scrollH2}`)

        let scrollResult = [scrollH1, scrollH2];
        preScrollHeight = scrollResult[0];
        scrollHeight = scrollResult[1];
      }
    }

    // console.log(await page.content())
    const list = await page.$$eval(config_eval.selector, async (d, config) => {
      let res = {
        data: [], 
        next: '' // 下一页url
      };
      for (let e of d) {
        const r = eval(`${config.datas}(e)`);
        if(r != false){
          res.data.push(r);
        }
      }
      return res;
    }, config_eval);

    
    await page.close(); // 关闭当前标签页
    return list;
  }

}
module.exports = Drive;
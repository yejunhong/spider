import puppeteer from 'puppeteer';
import devices from 'puppeteer/DeviceDescriptors';

// 配置
interface cfg {
  name?: string;
  url: string;
  mobile: string;
  cookie?: any;
  login?: string;
  user_agent?: string;
  list: any; // [] 列表
  detail?: any; // 详情
  next?: string; // 下一页
  scroll?: boolean;
}

class ChromeBrowser{
  
  public browser: any;

  /**
   * 创建浏览器
   */
  public async Create() {
    const options = {args: [
                      '--no-sandbox', 
                      '--disable-setuid-sandbox', 
                      '--process-per-tab', 
                      '--disable-images'
                    ]};
    this.browser = await puppeteer.launch(options);
  }

  /**
   * 打开一个标签页面
   * @param url // 访问http-url
   * @param config_name // 配置名称, 用于标签页面加载js配置文件
   * @param config // 配置信息
   */
  public async Page(url: string, config: cfg) {
    const page = await this.browser.newPage();
    if (config.cookie == true){
      for (let e of config.cookie) {
        await page.setCookie(e);     // 设置cookie
      }
    }
    // 伪造浏览器 
    if (config.user_agent != ""){
      await page.setUserAgent(config.user_agent);
    }
    // 是否启用手机模式
    if (config.mobile != ""){
      await page.emulate(config.mobile); // 手机浏览器模式
    }
    
    await page.goto(url);
    page.on('console', msg => console.log(msg.text()));
    // 注入函数到浏览器
    /*await page.exposeFunction('md5', text =>
      crypto.createHash('md5').update(text).digest('hex')
    );*/
    // 注入配置信息
    await page.addScriptTag({path: `${__dirname}/config/${config_name}.js`});
    // console.log(await page.content())
    await page.close(); // 关闭当前标签页
    return resData;
  }

  /**
   * get 内容
   * @param page 页面标签对象
   * @param config 配置{selector?: 选择器; func: 回调函数}
   */
  async QuerySelector(page: any, config: {selector?: string; func: any}): Promise<any> {
    const res = await page.$$eval(config.selector, async (e: any, func: any) => {
      return eval(`${func}(e)`)
    }, config.func);
    return res
  }

  /**
   * 自动滚动页面
   * @param page 页面标签对象
   * @param scroll 是否进行滚动
   */
  async PageScroll(page: any, scroll?: boolean) {
    if (scroll != true){
      return
    }
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

  /**
   * 爬取 列表
   * @param call // 客户端请求参数
   * @param callback // 推送信息到客户端
   */
  async CrawlList(call, callback){
    const request = call.request
    const config = require(`${__dirname}/config/${request.config_name}`);
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
    const config = require(`${__dirname}/config/${request.config_name}`);
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
    const config = require(`${__dirname}/config/${request.config_name}`);
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
    await page.goto(config.login.url);
    // 注入配置信息
    await page.addScriptTag({
      path: `${__dirname}/config/${config_name}.js`,
    });
    // 待定过期
    await page.waitForNavigation({timeout: 1000});
    const res = await page.$$eval(config.selector, async (d, config) => {
      const res = eval(`${config.login.operation}(e)`);
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
    if (config.user_agent != undefined && config.user_agent != false){
      await page.setUserAgent(config.user_agent);
    }
    // 是否启用手机模式
    if (config.mobile != undefined && config.mobile != false){
      await page.emulate(iPhone); // 手机浏览器模式
    }
    
    await page.goto(url);

    page.on('console', msg => console.log(msg.text()));
    // 注入函数到浏览器
    /*await page.exposeFunction('md5', text =>
      crypto.createHash('md5').update(text).digest('hex')
    );*/
    // 注入配置信息
    await page.addScriptTag({path: `${__dirname}/config/${config_name}.js`});
    // console.log(await page.content())
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

    const resData = {data: [], next: ""}; // 返回的对象
    // console.log(await page.content())
    if(config_eval.detail != undefined && config_eval.detail != false){
      resData.detail = await page.$$eval(config_eval.detail[0], async (e, func) => {
        const res = eval(`${func}(e)`);
        return res
      }, config_eval.detail[1]);
    }

    if(config_eval.next != undefined && config_eval.next != false){
      resData.next = await page.$$eval(config_eval.next[0], async (e, func) => {
        const res = eval(`${func}(e)`);
        return res
      }, config_eval.next[1]);
    }

    // console.log(await page.content())
    resData.data = await page.$$eval(config_eval.selector, async (d, config) => {
      let res = [];
      for (let e of d) {
        const r = eval(`${config.datas}(e)`);
        if(r != false){
          res.push(r);
        }
      }
      return res;
    }, config_eval);
    
    await page.close(); // 关闭当前标签页
    return resData;
  }

}
module.exports = ChromeBrowser;
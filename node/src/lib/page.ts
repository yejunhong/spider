import puppeteer from 'puppeteer';
import devices from 'puppeteer/DeviceDescriptors';

// 配置
interface cfg {
  name?: string;
  mobile?: string;
  cookie?: any;
  login?: string;
  user_agent?: string;
  list?: any; // [] 列表
  detail?: any; // 详情
  next?: string; // 下一页
  scroll?: boolean;
}

class Pages{
  
  public page: any;

  /**
   * 打开一个标签页面
   * @param url // 访问http-url
   * @param config // 配置信息
   */
  public async OpenTabPage(browser: any, config: cfg): Promise<any> {
    this.page = await browser.newPage();
    if ( config.cookie != undefined && config.cookie == true ){
      for (let e of config.cookie) {
        await this.page.setCookie(e);     // 设置cookie
      }
    }
    // 伪造浏览器 
    if ( config.user_agent != undefined && config.user_agent != "" ){
      await this.page.setUserAgent(config.user_agent);
    }
    // 是否启用手机模式
    if ( config.mobile != undefined && config.mobile != "" ){
      await this.page.emulate(config.mobile); // 手机浏览器模式
    }
    return this
  }
  
  /**
   * 打开一个标签页面
   * @param url // 访问http-url
   * @param config_name // 配置名称, 用于标签页面加载js配置文件
   * @param config // 配置信息
   */
  public async RequestUrl(url: string) {
    await this.page.goto(url);
    // page.on('console', msg => console.log(msg.text()));
    // 注入函数到浏览器
    /*await page.exposeFunction('md5', text =>
      crypto.createHash('md5').update(text).digest('hex')
    );*/
    // 注入配置信息
    // await page.addScriptTag({path: `${__dirname}/config/${config_name}.js`});
    // console.log(await page.content())
    // await page.close(); // 关闭当前标签页
  }

  /**
   * get 内容
   * @param page 页面标签对象
   * @param config 配置{selector?: 选择器; func: 回调函数}
   */
  public async QuerySelector(config: {selector?: string}): Promise<any> {
    const res = await this.page.$$(config.selector);
    return res
  }

  /**
   * 自动滚动页面
   * @param page 页面标签对象
   * @param scroll 是否进行滚动
   */
  public async PageScroll(scroll?: boolean) {
    if (scroll != true){
      return
    }
    let preScrollHeight = 0;
    let scrollHeight = -1;
/*
    const handle = await this.page.evaluateHandle(() => ({window, document}));
    const properties = await handle.getProperties();
    const windowHandle = properties.get('window');
    const documentHandle = properties.get('document');*/

    while(preScrollHeight !== scrollHeight) {
      // 详情信息是根据滚动异步加载，所以需要让页面滚动到屏幕最下方，通过延迟等待的方式进行多次滚动
      /*let scrollH1 = await this.page.evaluate(async () => {
        let h1 = document.body.scrollHeight;
        window.scrollTo(0, h1);
        return h1;
      });
      await this.page.waitFor(500);
      let scrollH2 = await this.page.evaluate(async () => {
        return document.body.scrollHeight;
      });
      console.log(`滚动页面高度：${scrollH2}`)
      let scrollResult = [scrollH1, scrollH2];
      preScrollHeight = scrollResult[0];
      scrollHeight = scrollResult[1];*/
    }
  }

  public async Html(res: any, selector: string): Promise<any> {
    return await res.$eval(selector, e => e.innerHTML);
  }

  public async AttrText(res: any, selector: string): Promise<any> {
    return await res.$eval(selector, e => e.innerHTML);
  }

  public async close(){
    await this.page.close();
  }

}
export default Pages;
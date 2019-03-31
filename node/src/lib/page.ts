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
    if ( config.cookie != undefined && config.cookie != "" ){
      console.log('设置cookie')
      for (let e of config.cookie) {
        await this.page.setCookie(e);     // 设置cookie
      }
    }
    // 伪造浏览器 
    if ( config.user_agent != undefined && config.user_agent != "" ){
      console.log('设置user_agent')
      await this.page.setUserAgent(config.user_agent);
    }
    // 是否启用手机模式
    if ( config.mobile != undefined && config.mobile != "" ){
      console.log('开启手机模式')
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
    return await this.page.content();
    // console.log(await this.page.content())
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
   * @param config 配置{selector?: 选择器; func: 回调函数}
   */
  public async QuerySelector(config: {selector?: string}): Promise<any> {
    const res = await this.page.$(config.selector);
    return res
  }

  /**
   * get json内容
   * @param config 配置{selector?: 选择器; func: 回调函数}
   */
  public async JsonContent(config: {selector?: string}): Promise<any> {
    const res = await this.page.$eval(config.selector, (e: any) => e.innerHTML);
    return res
  }

  /**
   * get 内容
   * @param config 配置{selector?: 选择器; func: 回调函数}
   */
  public async QuerySelectors(config: {selector?: string}): Promise<any> {
    const res = await this.page.$$(config.selector);
    return res
  }

  public GetUrl(): string {
    return this.page.url();
  }

  /**
   * 自动滚动页面
   * @param scroll 是否进行滚动
   */
  public async PageScroll(scroll?: boolean) {
    if (scroll != true){
      return
    }
    await this.page.waitFor(1000)
    console.log("滚动开始")
    let scrollEnable: boolean = true;
    let scrollStep: number = 300; //每次滚动的步长
    let scrollTop: number = 0;
    const window = await this.page.evaluateHandle('window'); // 'window'对象
    const document = await this.page.evaluateHandle('document'); // 'document'对象
    while(scrollEnable){
      var scrollHeight = await this.page.evaluate((document: any) => {
        return document.body.scrollHeight
      }, document)
      /*var scrollTop = await this.page.evaluate((document: any) => {
        return document.body.scrollTop
      }, document)*/
      scrollTop = scrollTop + scrollStep;
      console.log(`滚动高度：${scrollTop}, 最大高度：${scrollHeight}`)
      if (scrollTop < scrollHeight) {
        scrollTop = scrollTop + scrollStep
        await this.page.evaluate((window: any, scrollTop: number, scrollStep: number) => {
          window.scrollTo(scrollTop, scrollTop + scrollStep)
        }, window, scrollTop, scrollStep)
        await this.page.waitFor(800)
      }else {
        scrollEnable = false
      }
    }
    
    /*
    while (scrollEnable) {
      scrollEnable = await this.page.evaluate((body: any, scrollStep: number) => {
        let scrollTop = body.scrollTop;
        body.scrollTop = scrollTop + scrollStep;
        console.log(body.scrollHeight)
        console.log(scrollTop)
        window.scrollTo(0, 100)
        return body.scrollHeight > scrollTop + 100 ? true : false
      }, bodyHandle, scrollStep);
      
      await this.page.waitFor(600)
    }*/
    console.log("滚动结束")
  }

  public async close(){
    await this.page.close({runBeforeUnload: true});
    // console.log(this.page.isClosed())
  }

}
export default Pages;
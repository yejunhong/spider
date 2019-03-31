import Pages from './lib/page';
import Element from './lib/element';

class Spider {

  /**
   * 新建一个标签页
   * @param config 
   */
  public async newPage(browser, config: any){
    const page = new Pages();
    await page.OpenTabPage(browser, config);
    return page
  }

  public async LoginPage(browserPage: any, config: any, islogin: boolean) {
    if (islogin === true) {
      await browserPage.RequestUrl(config.url);
      await browserPage.page.waitFor(1000)
      console.log("检查是否需要登录");
      const e = await browserPage.page.$(config.click)
      if (e == null) {
        return;
      }
      console.log("输入账号");
      await browserPage.page.type(config.user.selector, config.user.value); // 立即输入账号
      console.log("输入密码");
      await browserPage.page.type(config.pass.selector, config.pass.value); // 立即输入密码
      await browserPage.page.waitFor(600)
      console.log("点击进行登录");
      await browserPage.page.tap(config.click); // 点击登录
      await browserPage.page.waitFor(10000)
    }
  }

  /**
   * 请求url
   * @param page
   * @param url
   * @param config
   */
  public async Request(page: any, url: string, config: any){
    const content = await page.RequestUrl(url);
    
    // console.log(res)
    let resdata: any = [];
    let next: any = false;

    if (config.jsonHandle != undefined) {
      // console.log(content)
      const res = await page.JsonContent({selector: config.selector});
      resdata = await config.jsonHandle(JSON.parse(res));
    }
    
    if (config.scroll != undefined) {
      await page.PageScroll(config.scroll);
    }
    if (config.handle != undefined) {
      const res = await page.QuerySelectors({selector: config.selector});
      resdata = await config.handle(res, Element);
    }

    // 有数据情况下尝试爬取下一页
    if (config.next != undefined && resdata.length > 0) {
      const resNext = await page.QuerySelector({selector: config.next.selector});
      next = await config.next.handle(resNext, page.GetUrl());
    }
    return {data: resdata, next: next}
  }

}
export default Spider;

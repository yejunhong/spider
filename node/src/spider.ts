import Pages from './lib/page';
import Element from './lib/element';
import url from 'url';

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
  public async Request(page: any, url: string, config: any, PageCfg: any){

    let content: string = '';
    if(config.browser_request == undefined || config.browser_request == true){
      content = await page.RequestUrl(url);
    } else {
      content = await page.SuperagentRequest(url, PageCfg)
    }

    if (config.print != undefined && config.print == true) {
      console.log(content)
    }
    await this.Event(page, config);
    // console.log(res)
    let resdata: any = {list: [], detail: {}};
    let next: any = false;

    if (config.jsonHandle != undefined) {
      const res = await page.JsonContent({selector: config.selector});
      const rData = await config.jsonHandle(JSON.parse(res));
      if (rData.list != undefined) {
        resdata.list = rData.list;
        resdata.detail = rData.detail;
      } else {
        resdata.list = rData;
        resdata.detail = {};
      }
    }
    
    if (config.scroll != undefined) {
      await page.PageScroll(config.scroll);
    }
    if (config.handle != undefined) {
      const res = await page.QuerySelectors({selector: config.selector});
      const rData = await config.handle(res, Element);
      if (rData.list != undefined) {
        resdata.list = rData.list;
        resdata.detail = rData.detail;
      } else {
        resdata.list = rData;
        resdata.detail = {};
      }
    }
    
    // 有数据情况下尝试爬取下一页
    if (config.next != undefined && resdata.list.length > 0) {
      const resNext = await page.QuerySelector({selector: config.next.selector});
      next = await config.next.handle(resNext, page.GetUrl());
    }
    return {data: resdata.list, detail: resdata.detail, next: next}
  }

  /**
   * 操作页面事件
   * @param page 页面标签对象
   * @param config 配置信息
   */
  public async Event(page: any, config: any) {
    if (typeof config.event == 'function') {
      await config.event(page.page);
    }
  }

}
export default Spider;

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

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
    await page.RequestUrl(url);
    const res = await page.QuerySelector({selector: config.selector});
    // console.log(res)
    let resdata: any = [];
    let next: any = '';
    if (config.handle != undefined) {
      resdata = await config.handle(res, Element);
    }
    if (config.next != undefined) {
      const resNext = await page.QuerySelector({selector: config.next.selector});
      next = await config.next.handle(resNext[0]);
    }
    return {data: resdata, next: next}
  }

}
export default Spider;

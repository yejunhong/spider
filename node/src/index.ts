import grpc from 'grpc';
import Browser from './lib/browser';
import Pages from './lib/page';
import Element from './lib/element';

const {Page, Book, Chapter, Content} = require('./config/kuaikanmanhua');
class Spider {
  public browser: any;
  public async newPage(config: any){
    const page = new Pages();
    await page.OpenTabPage(this.browser, config);
    return page
  }

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
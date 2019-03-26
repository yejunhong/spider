import grpc from 'grpc';
import Browser from './lib/browser';
import Pages from './lib/page';
import Element from './lib/element';
// const DriveClass = require('./drive');
// const node_rpc = grpc.load(`${__dirname}/../grpc/drive.proto`).grpc;

(async function () {
  
  const browser = await Browser.Create()

  const {Page, Book} = require('./config/kuaikanmanhua');

  const page = new Pages();
  await page.OpenTabPage(browser, Page);
  await page.RequestUrl("https://www.kuaikanmanhua.com/tag/0");
  const res = await page.QuerySelector({selector: Book.selector});
  // console.log(res)
  let resdata: any = [];
  if (Book.handle != undefined) {
    resdata = await Book.handle(res, Element);
  }
  console.log(resdata)
  await page.close();
  /*

  const drive_ = new DriveClass();
  await drive_.CreateBrowser(); // 创建浏览器
  
  const server = new grpc.Server();
  // testProto.browser.service .proto 设置的service
  // {ping: test} ping提供的函数 test回调的函数
  server.addProtoService(node_rpc.browser.service, {
      CrawlList: async (call, callback) => { await drive_.CrawlList(call, callback) },
      CrawlChapter: async (call, callback) => { await drive_.CrawlChapter(call, callback) },
      CrawlChapterContent: async (call, callback) => { await drive_.CrawlChapterContent(call, callback) },
    })
  server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
  server.start()*/
})();
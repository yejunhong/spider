import grpc from 'grpc';
import browser from './lib/browser';
import NewSpider from './spider';
/*
(async function(){

  const newBrowser = await browser.Create()
  const spider = new NewSpider();
  
  console.log(newBrowser)
  const page = await spider.newPage(newBrowser, Page);
  console.log(page)
})()*/
let newBrowser: any;

class Request{
  public async Write(steam: any, spider: any, page: any, url: string, config: any) {
    const res = await spider.Request(page, url, config);
    // console.log(res)
    console.log(`获取数量：${res.data.length}`)
    if(res.data.length > 0) {
      steam.write({data: res.data});
      if (res.next === false) {
        await page.close() // 关闭页面
        steam.end();
        return
      }
      console.log(`下一页：${res.next}`)
      await this.Write(steam, spider, page, res.next, config)
      return
    }
    await page.close() // 关闭页面
    steam.end();
  }
}

class GrpcServer {

  /**
   * 启动服务器
   */
  public async Run() {
    const server = new grpc.Server();
    const node_rpc = grpc.load(`${__dirname}/../../grpc/drive.proto`).grpc;
    newBrowser = await browser.Create()
    server.addProtoService(node_rpc.browser.service, {
      Book: this.Book,
      Chapter: this.Chapter,
      Content: this.Content,
    });
    server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure());
    server.start()
  }

  /**
   * 书籍列表
   * @param steam
   */
  public Book(steam: any) {
    steam.on('data', async (note: any) => {
      delete require.cache[require.resolve(`${__dirname}/config/${note.config_name}`)];
      const {Page, Login, Book} = require(`${__dirname}/config/${note.config_name}`);
      // console.log(note)
      const spider = new NewSpider();
      const page = await spider.newPage(newBrowser, Page);
      if (Book.islogin != undefined) {
        await spider.LoginPage(page, Login, Book.islogin);
      }
      const res = new Request()
      await res.Write(steam, spider, page, note.url, Book)
      
    });
    steam.on('end', () => {
      steam.end();
    });
  }

  /**
   * 书籍章节
   * @param steam  
   */
  public Chapter(steam: any){
    steam.on('data', async (note: any) => {
      const {Page, Chapter} = require(`${__dirname}/config/${note.config_name}`);
      const spider = new NewSpider();
      const page = await spider.newPage(newBrowser, Page);
      const res = new Request()
      await res.Write(steam, spider, page, note.url, Chapter)
    });
    steam.on('end', () => {
      steam.end();
    });
  }

  /**
   * 书籍章节内容
   * @param steam  
   */
  public Content(steam: any){
    steam.on('data', async (note: any) =>{
      const {Page, Content} = require(`${__dirname}/config/${note.config_name}`);
      const spider = new NewSpider();
      const page = await spider.newPage(newBrowser, Page);
      const res = new Request()
      await res.Write(steam, spider, page, note.url, Content)
      
    });
    steam.on('end', () => {
      steam.end();
    });
  }

  

}

const grpcService = new GrpcServer();
grpcService.Run();
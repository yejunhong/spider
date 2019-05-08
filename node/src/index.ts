import grpc from 'grpc';
import browser from './lib/browser';
import NewSpider from './spider';

let newBrowser: any;

class Request{
  public async Write(steam: any, spider: any, page: any, note: any, config: any, PageCfg: any) {
    const res = await spider.Request(page, note.url, config, PageCfg);
    console.log(`获取数量：${res.data.length}，url：${note.url}`)
    if(res.data.length > 0) {
      steam.write({data: res.data, detail: res.detail, next: res.next?true: false, id: note.id});
      if (res.next === false) {
        await page.close() // 关闭页面
        return
      }

      await this.Write(steam, spider, page, {url: res.next}, config, PageCfg)
      return
    }
    steam.write({data: [], detail: [], next: false, id: note.id});
    await page.close() // 关闭页面
  }

  public async WriteNoDetail(steam: any, spider: any, page: any, note: any, config: any, PageCfg: any) {
    const res = await spider.Request(page, note.url, config, PageCfg);
    console.log(`获取数量：${res.data.length}，url：${note.url}`)
    if(res.data.length > 0) {
      steam.write({data: res.data, next: res.next?true: false, id: note.id});
      if (res.next === false) {
        await page.close() // 关闭页面
        return
      }
      await this.Write(steam, spider, page, {url: res.next}, config, PageCfg)
      return
    }
    steam.write({data: [], next: false, id: note.id});
    await page.close() // 关闭页面
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
      if(note.url == 'end') {
        steam.end();
        return
      }
      // 删除缓冲区 重新加载文件
      delete require.cache[require.resolve(`${__dirname}/config/${note.config_name}`)];
      const {Page, Login, Book} = require(`${__dirname}/config/${note.config_name}`);
      // console.log(note)
      const spider = new NewSpider();
      const spiderPage = await spider.newPage(newBrowser, Page);
      if (Book.islogin != undefined) {
        await spider.LoginPage(spiderPage, Login, Book.islogin);
      }
      const res = new Request()
      await res.Write(steam, spider, spiderPage, note, Book, Page)
    });
    steam.on('end', () => {
      console.log('book steam end')
      steam.end();
    });
  }

  /**
   * 书籍章节
   * @param steam  
   */
  public Chapter(steam: any){
    steam.on('data', async (note: any) => {
      if(note.url == 'end') {
        steam.end();
        return
      }
      // console.log(note)
      delete require.cache[require.resolve(`${__dirname}/config/${note.config_name}`)];
      const {Page, Chapter} = require(`${__dirname}/config/${note.config_name}`);
      const spider = new NewSpider();
      const spiderPage = await spider.newPage(newBrowser, Page);
      const res = new Request()
      await res.Write(steam, spider, spiderPage, note, Chapter, Page)
    });
    steam.on('end', () => {
      console.log('chapter steam end')
      steam.end();
    });
  }

  /**
   * 书籍章节内容
   * @param steam  
   */
  public Content(steam: any){
    steam.on('data', async (note: any) =>{
      if(note.url == 'end') {
        steam.end();
        return
      }
      delete require.cache[require.resolve(`${__dirname}/config/${note.config_name}`)];
      const {Page, Content} = require(`${__dirname}/config/${note.config_name}`);
      const spider = new NewSpider();
      const spiderPage = await spider.newPage(newBrowser, Page);
      const res = new Request()
      await res.WriteNoDetail(steam, spider, spiderPage, note, Content, Page)
    });
    steam.on('end', () => {
      console.log('content steam end')
      steam.end();
    });
  }
}

const grpcService = new GrpcServer();
grpcService.Run();
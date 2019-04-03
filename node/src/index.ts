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
  public async Write(steam: any, spider: any, page: any, url: string, config: any, PageCfg: any) {
    const res = await spider.Request(page, url, config, PageCfg);
    // console.log(res)
    console.log(`获取数量：${res.data.length}，url：${url}`)
    if(res.data.length > 0) {
      steam.write({data: res.data, next: res.next?true: false});
      if (res.next === false) {
        await page.close() // 关闭页面
        // steam.end();
        return
      }
      console.log(`下一页：${res.next}`)
      await this.Write(steam, spider, page, res.next, config, PageCfg)
      return
    }
    steam.write({data: [], next: false});
    await page.close() // 关闭页面
    // steam.end();
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
      await res.Write(steam, spider, spiderPage, note.url, Book, Page)
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
      await res.Write(steam, spider, spiderPage, note.url, Chapter, Page)
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
      await res.Write(steam, spider, spiderPage, note.url, Content, Page)
    });
    steam.on('end', () => {
      console.log('content steam end')
      steam.end();
    });
  }
}

const grpcService = new GrpcServer();
grpcService.Run();
/*
const request = require('superagent');
request.get('http://c976.yinsha5.com/index/book/bookcontent/chapeterid/795250?1554191723')
  .set('User-Agent', 'Mozilla/5.0 (iPhone; CPU iPhone OS 12_1_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16D57 MicroMessenger/7.0.3(0x17000321) NetType/WIFI Language/zh_CN')
  .set('Cookie', `_novelOpenid=oFFzA514uRnqPLc908Y1Zwn8sizc`)
  .redirects(0)
  .then(res => {
    console.log(11111)
    // console.log(res.text)
     // console.log(11111)
    // res.body, res.headers, res.status
  })
  .catch(err => {
    console.log(err.response.text)
  });
  console.log(2222)*/
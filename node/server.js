 
const grpc = require('grpc');
const drive = grpc.load(`${__dirname}/../grpc/drive.proto`).grpc;

class Serve{

  CrawlList(call, callback){
    call.request.url
    call.request.config_name
  }

  CrawlChapter(call, callback){

  }

  CrawlChapterContent(call, callback){

  }

  main(){

    const server = new grpc.Server();
    // testProto.browser.service .proto 设置的service
    // {ping: test} ping提供的函数 test回调的函数
    server.addProtoService(drive.browser.service, {
        CrawlList: this.CrawlList,
        CrawlChapter: this.CrawlChapter,
        CrawlChapterContent: this.CrawlChapterContent
      })
    server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
    server.start()

  }

}

call.request.url
call.request.config_name


CrawlChapter
CrawlChapterContent
function test(call, callback) {
  console.log(testProto)
  callback(null, {message: 'Pong'})
}


 
const grpc = require('grpc');
const DriveClass = require('./drive');
const node_rpc = grpc.load(`${__dirname}/../grpc/drive.proto`).grpc;

class GrpcServer {

  public Run() {
    const server = new grpc.Server();
    server.addProtoService(node_rpc.browser.service, {
      Book: this.Book,
      Chapter: async (call, callback) => { await drive_.CrawlChapter(call, callback) },
      Content: async (call, callback) => { await drive_.CrawlChapterContent(call, callback) },
    })
    server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
    server.start()
  }

  public Book(call) {
    call.on('data', function(note) {
      var key = pointKey(note.location);
      if (route_notes.hasOwnProperty(key)) {
        _.each(route_notes[key], function(note) {
          call.write(note);
        });
      } else {
        route_notes[key] = [];
      }
      // Then add the new note to the list
      route_notes[key].push(JSON.parse(JSON.stringify(note)));
    });
    call.on('end', function() {
      call.end();
    });
  }

}

const grpcService = new GrpcServer();
grpcService.Run();

(async function () {
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
  server.start()
})();

/*
function routeChat(call) {
  call.on('data', function(note) {
    var key = pointKey(note.location);
    if (route_notes.hasOwnProperty(key)) {
      _.each(route_notes[key], function(note) {
        call.write(note);
      });
    } else {
      route_notes[key] = [];
    }
    // Then add the new note to the list
    route_notes[key].push(JSON.parse(JSON.stringify(note)));
  });
  call.on('end', function() {
    call.end();
  });
}*/
 import grpc from 'grpc';

class GrpcServer {

  /**
   * 启动服务器
   */
  public Run() {
    const server = new grpc.Server();
    const node_rpc = grpc.load(`${__dirname}/../grpc/drive.proto`).grpc;
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
    steam.on('data', (note: any) => {
      console.log(note)
      steam.write(note);
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
    steam.on('data', (note: any) => {
      console.log(note)
      steam.write(note);
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
    steam.on('data', (note: any) =>{
      console.log(note)
      steam.write(note);
    });
    steam.on('end', () => {
      steam.end();
    });
  }

}

const grpcService = new GrpcServer();
grpcService.Run();
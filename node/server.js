 
const grpc = require('grpc');

const PROTO_PATH = __dirname + '/../grpc/drive.proto'
const testProto = grpc.load(PROTO_PATH).drive
 
function test(call, callback) {
  console.log(111)
  callback(null, {message: 'Pong'})
}
 
const server = new grpc.Server();
server.addProtoService(testProto.browser.service, {test: '22'})
server.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure())
server.start()
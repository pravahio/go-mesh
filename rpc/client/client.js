var messages = require('./rpc_pb');
var services = require('./rpc_grpc_pb');

var grpc = require('grpc');

function main() {
    var client = new services.MeshClient('127.0.0.1:5555', grpc.credentials.createInsecure());

    var info = new messages.PeerTopicInfo()
    info.setTopic('GGN.BUS')

    var call = client.subscribe(info)

    call.on('data', function(data) {
        console.log(new Buffer.from(data.getRaw()).toString())
    })
}

function sendPublishMsg() {
    let call = client.publish(function(error, stats) {
        console.log(error)
    })

    var pubData = new messages.PublishData()
    var info = new messages.PeerTopicInfo()
    var data = new messages.Data()

    info.setTopic('GGN.BUS')

    data.setRaw(Buffer.from('hello', 'utf-8'))

    pubData.setInfo(info)
    pubData.setData(data)

    call.write(pubData)
    console.log("written")
    //call.end()
    console.log("end")
}

main();

const path = require('path');
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = './proto/crds.proto';
const PORT = 7000;

var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var crds = grpc.loadPackageDefinition(packageDefinition).crds;
var client = new crds.EchoService('localhost:' + PORT,
    grpc.credentials.createInsecure());

export default (req, res) => {
    const {
        query: { message },
    } = req
    client.echo({ input: message }, (err, response) => {
        if (err) {
            res.status(200).json(err)
        } else {
            res.status(200).json(response)
        }
    });
}

// import { NextApiRequest, NextApiResponse } from 'next'
import * as grpc from 'grpc';
import * as protoLoader from '@grpc/proto-loader';

const PROTO_PATH = './proto/crds.proto';
const PORT = 7000;

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
const crds = grpc.loadPackageDefinition(packageDefinition).crds;
const client = new crds.EchoService('localhost:' + PORT,
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

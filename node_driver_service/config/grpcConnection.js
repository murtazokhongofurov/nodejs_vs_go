const protoLoader = require("@grpc/proto-loader");
const grpc = require("@grpc/grpc-js");

const cfg = require('./index');
const log = require('./logger')

const driverService = require("../service/driver");
const carService = require("../service/car");


const PROTO_URL = __dirname + "/../protos/driver_service/driver_service.proto";
const packageDefinition = protoLoader.loadSync(PROTO_URL, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true
});

const driverProto = 
    grpc.loadPackageDefinition(packageDefinition).driver_service;

var server = new grpc.Server();


server.addService(driverProto.DriverService.service, driverService);
server.addService(driverProto.CarService.service, carService);

server.bindAsync(
    "0.0.0.0:" + cfg.RPCPort,
    grpc.ServerCredentials.createInsecure(),
    (err, bindPort) => {
        if (err) {
            throw new Error("Error while binding grpc server to the port");
        }

        log.info("grpc server is running at %s", bindPort);
        server.start();
    }
);
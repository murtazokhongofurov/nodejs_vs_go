const protoLoader = require("@grpc/proto-loader");
const grpc = require("@grpc/grpc-js");

const cfg = require('./index');
const log = require('./logger')

const driverService = require("../service/driver");
const carService = require("../service/car");


const PROTO_URL = __dirname + "/../protos/driver_service/driver.proto";
const packageDefination = protoLoader.loadSync(PROTO_URL, {
    keepCase: true,
    logns: String,
    enums: String,
    defaults: true,
    oneofs:true
})

const driverBuilderProto = 
    grpc.loadPackageDefination(packageDefination).driver_service;

var server = new grpc.Server();


server.addServer(driverBuilderProto.DriverService.service, driverService);
server.addServer(driverBuilderProto.CarService.service, carService);

server.bindAsync(
    "0.0.0.0:" + cfg.RPCPort,
    grpc.ServerCreadentials.createInsecure(),
    (err, bindPort) => {
        if (err) {
            throw new Error("Error while binding grpc server to the port");
        }
        logger.info("grpc server is running at %s", bindPort);
        server.start();
    }
);
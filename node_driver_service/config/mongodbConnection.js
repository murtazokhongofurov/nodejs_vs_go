const cfg = require("./index");
const mongoose = require("mongoose");
const logger = require("./logger");

let mongoDBURL = //`mongodb://node_driver_service:password@localhost:27017/mongo_driver_database_name`
    "mongodb://" + 
    cfg.mongoUser + 
    ":" + cfg.mongoPassword + 
    "@" + cfg.mongoHost +
    ":" + cfg.mongoPort + 
    "/" + 
    cfg.mongoDatabase;

mongoose.set ('strictQuery', true);
mongoose.connect(
    mongoDBURL, 
    {
        useNewUrlParser: true,
        useUnifiedTopology: true,
    },
    (err) => {
        if (err) {
            console.log("Error while connecting to database (" + 
            mongoDBUrl + ") "+ err.message);
            logger.error("Error while connecting to database (" + 
            mongoDBUrl + ") "+ err.message);
        }

        logger.info("Connected to mongoDB")
    }
);
const config = {
    environment: getConf("NODE_ENV", "develop"),
    mongoHost: getConf("MONGO_HOST", "localhost"),
    mongoPort: getConf("MONGO_PORT", "27017"),
    mongoUser: getConf("MONGO_USER", "developer"),
    mongoPassword: getConf("MONGO_PASSWORD", "2002"),
    mongoDatabase: getConf("MONGO_DATABASE", "mongo_driver_db"),
    RPCPort: getConf("RPC_PORT", 2001),

}


function getConf(name, def = "") {
    if (process.env[name]) {
        return process.env[name];
    }
    return def;
}

module.exports = config;
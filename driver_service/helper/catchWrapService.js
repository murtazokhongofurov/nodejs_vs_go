const logger = require("../config/logger");
const grpc = require("@grpc/grpc-js");

module.exports = (namespace, fn) => {
    return async (call, callback) => {
        logger.info(
            `${namespace}: requested - ${JSON.stringify(call.request)}`
        );

        try {
            const resp = await fn(call.request);

            logger.info(`${namespace}: succeeded`);
            callback(null, resp);
        } catch (error) {
            logger.error(`${namespace}: failed with error: ${error.message}`);

            callback({
                code: grpc.status.INTERNAL,
                message: error.message
            });
        }
    };
};

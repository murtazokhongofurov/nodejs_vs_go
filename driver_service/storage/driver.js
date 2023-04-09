const catchWrapDb = require("../helper/catchWrapDb");
const Driver = require("../models/driver");


let NAMESPACE = "storage.driver";

let driverStore =  {
    create: catchWrapDb(`${NAMESPACE}.create`, async(data) => {
        const driver = new Driver(data);
        const response = await driver.save();

        return response;
    }),
    update: catchWrapDb(`${NAMESPACE}.update`, async(data) => {
        const driver = await Driver.updateOne(
            {
                id: data.id
            },
            {
                $set: data
            }
        )
        return driver;
    }),
    getById: catchWrapDb(`${NAMESPACE}.getById`, async(data) => {
        const driver = await Driver.findOne({id: data.id});
        return driver;
    }),
    delete: catchWrapDb(`${NAMESPACE}.delete`, async(data) => {
        const driver = await Driver.deleteOne({id: data.id});
        return driver;
    }),
};

module.exports = driverStore;


const catchWrapService = require("../helper/catchWrapService");
const driverStore = require("../storage/driver");


const driverService = {
    Create: catchWrapService(`service.driver.create`, driverStore.create),
    Update: catchWrapService(`service.driver.update`, driverStore.update),
    GetById: catchWrapService(`service.driver.getById`, driverStore.getById),
    Delete: catchWrapService(`service.driver.delete`, driverStore.delete)
}

module.exports = driverService;
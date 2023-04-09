const catchWrapService = require("../helper/catchWrapService");
const carStore = require("../storage/car");


const carService = {
    Create: catchWrapService(`service.car.create`, carStore.create),
    Update: catchWrapService(`service.car.update`, carStore.update),
    GetById: catchWrapService(`service.car.getById`, carStore.getById),
    Delete: catchWrapService(`service.car.delete`, carStore.delete)
};

module.exports = carService;
const catchWrapDb = require("../helper/catchWrapDb");
const Car = require("../models/car");

let NAMESPACE = "storage.car";
let carStore = {
    create: catchWrapDb(`${NAMESPACE}.create`, async(data) => {
        const car = new Car(data);
        const response = await car.save();

        return response
    }),
    update: catchWrapDb(`${NAMESPACE}.update`, async(data) => {
        const car = Car.updateOne(
            {
                id: data.id
            },
            {
                $set: data
            }
        )
        return car;
    }),
    getById: catchWrapDb(`${NAMESPACE}.getById`, async(data) => {
        const car = Car.findOne(
            {
                id: data.id
            }
        );
        return car;
    }),
    delete: catchWrapDb(`${NAMESPACE}.delete`, async(data) => {
        const car = Car.deleteOne({id: data.id});
        return car;
    }),
}

module.exports = carStore;
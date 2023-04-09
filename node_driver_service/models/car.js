const mongoose = require("mongoose");
const {v4} = require("uuid");

const CarSchema = mongoose.Scema(
    {
        id: {
            type: String,
            default: v4,
            unique: true
        },
        driver_id: {
            type: String,
            unique: true
        },
        car_model: {
            type: String,
            maxlength: 100
        },
        manufacture_year: {
            type: String
        },
        number_plate: {
            type: String,
            maxlength: 9
        },
        technical_licence: {
            type: String,
            maxlength: 9
        },
        car_type: {
            type: String
        },
        car_image: {
            type: String
        },
        color: {
            type: String,
            maxlength: 100
        }
    },
    {
        timestamps: {createdAt: "created_at", updatedAt: "updated_at"},
        toObject: { virtuals: true },
        toJSON: { virtuals: true },
    }
);
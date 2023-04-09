const mongoose = require("mongoose");
const {v4} = require("uuid");

const DriverSchema = mongoose.Schema(
    {
        id: {
            type: String,
            default: v4,
            unique: true
        },
        first_name: {
            type: String,
            required: [true, "Driver must has name"]
        },
        last_name: {
            type: String,
            maxlength: 100
        },
        phone_number: {
            type: String,
            maxlength: 20
        }
    },
    {
        timestamps: {createdAt: "created_at", updatedAt: "updated_at"},
        toObject: { virtuals: true },
        toJSON: { virtuals: true },
    }
);

module.exports = mongoose.model("DriverApp", DriverSchema);
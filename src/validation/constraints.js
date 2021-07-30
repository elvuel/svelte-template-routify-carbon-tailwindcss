export const constraints = {
    username: [{
        required: true,
        type: 'string',
        message: "must be present"
    }, {
        min: 3,
        message: "must be at least 3 characters"
    }],
    password: {
        required: true,
        type: 'string',
        message: "must be present"
    }
}
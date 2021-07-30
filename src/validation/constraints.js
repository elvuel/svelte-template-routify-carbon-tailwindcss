export const defaults = {
    username: {
        presence: true,
        type: string,
        message: "must be present"
    },
    password: {
        presence: true,
        length: {
            minimum: 4,
            message: "must be at least 4 characters"
        },
        message: "must be present"
    }
}
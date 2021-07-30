export function constraints(translate) {
    return {
        username: [{
            required: true,
            type: 'string',
            // message: translate('constraints.username.required')
            // message: x('constraints.username.required')
        }, {
            min: 3,
            // message: () => {
            //     return translate('constraints.username.minlength')
            // }
        }],
        password: [{
            required: true,
            type: 'string',
            // message: "must be present"
        }, {
            min: 6,
            // message: "must be at least 6 characters"
        }]
    }
}
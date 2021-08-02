export default function login(translate) {
    return {
        username: [{
            required: true,
            type: 'string',
            // message: translate('validation.scene.login.username.required')
        }, {
            min: 3,
        }],
        password: [{
            required: true,
            type: 'string',
        }, {
            min: 6,
        }]
    }
}
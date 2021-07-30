export const constraints = {
    username: {
        required: 'Username is required',
        minlength: 'Username must be at least {%d} characters long',
    },
    password: {
        required: 'Password is required',
        minlength: 'Password must be at least {%d} characters long',
    }
}
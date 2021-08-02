import {
    user
} from "../../../model/user"

export const login = {
    username: {
        required: `${user.username}必填`,
        minlength: '%s must be at least %s characters'
    },
    password: {
        required: '密码必填',
        minlength: '密码最少为6个字符'
    }
}
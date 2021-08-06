import {
    auth
} from './auth'
import {
    users,
    createUser,
    getUser,
    updateUser,
    deleteUser,
    deleteUsers
} from './user'


export default {
    auth: auth,
    users: users,
    createUser: createUser,
    getUser: getUser,
    updateUser: updateUser,
    deleteUser: deleteUser,
    deleteUsers: deleteUsers
}
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

import {
    dummies,
    createDummy,
    getDummy,
    updateDummy,
    deleteDummy,
    deleteDummies
} from './dummy'


export default {
    auth: auth,

    users: users,
    createUser: createUser,
    getUser: getUser,
    updateUser: updateUser,
    deleteUser: deleteUser,
    deleteUsers: deleteUsers,

    dummies: dummies,
    createDummy: createDummy,
    getDummy: getDummy,
    updateDummy: updateDummy,
    deleteDummy: deleteDummy,
    deleteDummies: deleteDummies,
}
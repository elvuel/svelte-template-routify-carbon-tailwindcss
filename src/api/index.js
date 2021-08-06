import axios from './common.js'
import qs from 'qs'

export async function auth(data = {}) {
    return axios.post('/auth', data)
}

export async function users(data = {}) {
    let path = "/users"
    if (Object.keys(data).length > 0) {
        path += `?${qs.stringify(data, { arrayFormat: 'repeat' })}`
    }
    return axios.get(path)
}

export async function createUser(data = {}) {
    return axios.post('/users', data)
}

export async function getUser(id) {
    return axios.get(`/users/${id}`)
}

export async function updateUser(id, data = {}) {
    return axios.put(`/users/${id}`, data)
}

export async function deleteUser(id) {
    return axios.delete(`/users/${id}`)
}

export async function deleteUsers(ids) {
    return axios.delete(`/users`, {
        data: ids
    })
}
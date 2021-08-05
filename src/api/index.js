import axios from './common.js'

export async function auth(data = {}) {
    return axios.post('/auth', data)
}

export async function users() {
    return axios.get('/users')
}

export async function createUser(data = {}) {
    return axios.post('/users', data)
}

export async function getUser(id) {
    return axios.get(`/users/${id}`)
}
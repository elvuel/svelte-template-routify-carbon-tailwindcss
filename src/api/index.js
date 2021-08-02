import axios from './common.js'

export async function auth(data = {}) {
    return axios.post('/auth', data)
}

export async function users() {
    return axios.get('/users')
}
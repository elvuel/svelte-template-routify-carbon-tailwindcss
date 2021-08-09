import axios from './common.js'
import qs from 'qs'

export async function dummies(data = {}) {
    let path = "/dummies"
    if (Object.keys(data).length > 0) {
        path += `?${qs.stringify(data, { arrayFormat: 'repeat' })}`
    }
    return axios.get(path)
}

export async function createDummy(data = {}) {
    return axios.post('/dummies', data)
}

export async function getDummy(id) {
    return axios.get(`/dummies/${id}`)
}

export async function updateDummy(id, data = {}) {
    return axios.put(`/dummies/${id}`, data)
}

export async function deleteDummy(id) {
    return axios.delete(`/dummies/${id}`)
}

export async function deleteDummies(ids) {
    return axios.delete(`/dummies`, {
        data: ids
    })
}
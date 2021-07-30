import {
    newAxios
} from "../utils/axios";

// process.env.NODE_ENV
const baseURL = process.env.NODE_ENV === 'production' ? 'https://api.todoist.com' : 'http://localhost:8980';

const axios = newAxios({
    baseURL
})
export default axios
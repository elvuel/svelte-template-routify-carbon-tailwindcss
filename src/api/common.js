import {
    newAxios
} from "../utils/axios";

// process.env.NODE_ENV
const baseURL = process.env.NODE_ENV === 'production' ? 'https://api.todoist.com' : 'http://localhost:8980';

const axios = newAxios({
    baseURL
}, [
    response => {
        if (response.data.error_code && response.data.error_code != 0) {
            // throw new Error(response.data.error_message)
            const {
                error_code,
                error_message
            } = response.data
            return Promise.reject({
                error_code,
                error_message
            });
        }
        return response.data.data
    }
])
export default axios
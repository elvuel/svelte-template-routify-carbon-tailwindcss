import axios from "axios"

export function newAxios(config, responses = []) {
    const instance = axios.create(config)
    instance.defaults.headers.common["Content-Type"] = "application/json; charset=utf-8"
    responses.forEach(func => {
        instance.interceptors.response.use(func)
    })
    // instance.interceptors.response.use(
    //     response => {
    //         if (response.data.error_code && response.data.error_code != 0) {
    //             // throw new Error(response.data.error_message)
    //             const {
    //                 error_code,
    //                 error_message
    //             } = response.data
    //             return Promise.reject({
    //                 error_code,
    //                 error_message
    //             });
    //         }
    //         return response.data.data
    //     }
    // )
    return instance
}

export function setAxiosInstanceRequestInterceptor(instance, conf) {
    instance.interceptors.request.use(conf)
    // 当登陆成功以后可以这样配置
    // instance.interceptors.request.use(async config => {
    //   if (config.url && config.url.charAt(0) === '/') {
    //     config.url = `${baseURL}${config.url}`;
    //   }

    //   config.headers.authorization = `Bearer ${getItem('token')}`;

    //   return config;
    // }
}
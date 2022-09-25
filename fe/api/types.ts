import axios, { AxiosError, AxiosInstance, AxiosRequestConfig } from 'axios'

export interface CustomFailureResponse {
    code: number
    message: string
}

export interface CustomSuccessResponse<T> {
    code: 0
    data: T
}

export type CustomResponse<T = any> = CustomFailureResponse | CustomSuccessResponse<T>

// 判断请求是否成功
export const isSuccess = (r: CustomResponse): r is CustomSuccessResponse<any> => {
  return r.code === 0
}

// 自定义实例类型，把返回值重写成自己的版本
export interface CustomAxiosInstance extends AxiosInstance {
    get<T = any, R = CustomResponse<T>>(url: string, config?: AxiosRequestConfig): Promise<R>;
    post<T = any, R = CustomResponse<T>>(url: string, data?: any, config?: AxiosRequestConfig): Promise<R>;
}

// 返回自己的实例类型
export const createAxios = (): CustomAxiosInstance => {
  const _axios = axios.create({
    baseURL: process.env.baseApiUrl
  })

  // 设置响应拦截器，把返回体重写成自己的样子
  _axios.interceptors.response.use((v: any): any => {
    return v.data
  }, (err: AxiosError) => {
    // 服务器已经响应
    if (err.response) {
      return err.response.data
    }
    // 服务器没有响应
    if (err.message.includes('timeout')) {
      // 响应超时
      return {
        code: -1,
        message: err.message
      }
    }
    if (err.message.includes('Network Error')) {
      // 网络错误
      return {
        code: -2,
        message: err.message
      }
    }
  })

  return _axios
}

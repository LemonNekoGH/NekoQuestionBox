import { createAxios, isSuccess } from './types'
const axios = createAxios()

export interface NewQuestionData {
  question: string // 问题
  value: string // captcha 值
  id: string // captcha id
}

export interface Question {
  question: string
  answeredAt: number
  answer: string
  createAt: number
}

export interface BingWallpaperData {
  images: {
    url: string
  }[]
}

const backend = {
  async isServerAvailable () {
    const res = await axios.get<never>('/ping')
    return isSuccess(res)
  },
  getCaptchaId () {
    return axios.get<string>('/captcha')
  },
  getCaptchaImage (id: string) {
    return axios.get<string>(`/captcha-image?id=${id}`)
  },
  submitQuestion (data: NewQuestionData) {
    return axios.post<never>('/question', data)
  },
  getWallpaper () {
    return axios.get<BingWallpaperData>('/bing')
  },
  getQuestions () {
    return axios.get<Question[]>('/question')
  }
}

export const api = {
  be: backend
}

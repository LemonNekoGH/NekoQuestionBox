import axios, { AxiosError, AxiosResponse } from 'axios'

axios.defaults.headers.get = {
  Accept: 'application/json'
}

axios.defaults.headers.post = {
  Accept: 'application/json'
}

axios.defaults.baseURL = process.env.baseApiUrl

export interface QuestionData {
  question: string
  captchaValue: string
  captchaId: string
}

const backend = {
  async isServerAvailable (): Promise<boolean> {
    try {
      const res = await axios.get('/')
      return res.status === 200
    } catch (e) {
      return false
    }
  },
  async getCaptchaId (): Promise<string> {
    try {
      const res = await axios.get('/captcha')
      return res.data.id
    } catch (e) {
      console.log(e.message)
      return ''
    }
  },
  submitQuestion (data: QuestionData): Promise<number> {
    return new Promise<number>((resolve) => {
      axios.post('/question', data).then((res: AxiosResponse) => {
        resolve(res.status)
      }).catch((e: AxiosError) => {
        if (e.response?.status) {
          resolve(e.response.status)
        }
        resolve(500)
      })
    })
  },
  async getWallpaper (): Promise<string> {
    const res = await axios.get('/bing-wallpaper')
    if (res.data.images[0].url) {
      return res.data.images[0].urlbase
    }
    return ''
  }
}

export const api = {
  be: backend
}

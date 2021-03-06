import axios, { AxiosError, AxiosResponse } from 'axios'

axios.defaults.headers.get = {
  Accept: 'application/json'
}

axios.defaults.headers.post = {
  Accept: 'application/json'
}

axios.defaults.baseURL = process.env.baseApiUrl

export interface SubmitData {
  question: string
  captchaValue: string
  captchaId: string
}

export interface ResponseData {
  question: string
  answerTime: number
  answer: string
  time: number
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
      return ''
    }
  },
  submitQuestion (data: SubmitData): Promise<number> {
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
    try {
      const res = await axios.get('/bing-wallpaper')
      if (res.data.images[0].url) {
        return res.data.images[0].urlbase
      }
      return 'api-failed'
    } catch (e) {
      return 'api-error'
    }
  },
  async getQuestions (): Promise<ResponseData[]> {
    try {
      const res = await axios.get('/question')
      return res.data
    } catch (e) {
      console.log(e)
      return []
    }
  }

}

export const api = {
  be: backend
}

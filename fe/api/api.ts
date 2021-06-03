import axios from 'axios'

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
  async submitQuestion (data: QuestionData) {
    try {
      const res = await axios.post('/question', data)
      console.log(res.data)
      return res.status
    } catch (e) {
      console.log(e.message)
    }
  }
}

export const api = {
  be: backend
}

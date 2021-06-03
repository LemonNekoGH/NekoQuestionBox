import axios from 'axios'

axios.defaults.headers.get = {
  Accept: 'application/json'
}
axios.defaults.baseURL = process.env.baseApiUrl

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
  }
}

export const api = {
  be: backend
}

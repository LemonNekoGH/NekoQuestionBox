import { NuxtOptionsEnv } from '@nuxt/types/config/env'

const development: NuxtOptionsEnv = {
  // 当 IP 换掉之后这里也要换
  baseApiUrl: 'http://192.168.2.62:5000'
}

const production: NuxtOptionsEnv = {
  baseApiUrl: 'https://qboxb.lemonneko.moe'
}

export const env: { development: NuxtOptionsEnv, production: NuxtOptionsEnv, [p: string]: NuxtOptionsEnv } = {
  development,
  production
}

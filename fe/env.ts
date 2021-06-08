import { NuxtOptionsEnv } from '@nuxt/types/config/env'

const development: NuxtOptionsEnv = {
  baseApiUrl: 'http://localhost:5000'
}

const production: NuxtOptionsEnv = {
  baseApiUrl: 'https://qboxb.lemonneko.moe'
}

export const env: { development: NuxtOptionsEnv, production: NuxtOptionsEnv, [p: string]: NuxtOptionsEnv } = {
  development,
  production
}

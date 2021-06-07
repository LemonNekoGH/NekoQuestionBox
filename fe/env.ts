import { NuxtOptionsEnv } from '@nuxt/types/config/env'

const development: NuxtOptionsEnv = {
  baseApiUrl: 'https://localhost'
}

const production: NuxtOptionsEnv = {
  baseApiUrl: 'https://qboxb.lemonneko.moe'
}

export const env: { development: NuxtOptionsEnv, production: NuxtOptionsEnv, [p: string]: NuxtOptionsEnv } = {
  development,
  production
}

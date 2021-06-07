import { NuxtConfig } from '@nuxt/types'
import { env } from './env'

let nodeEnv: string | undefined = process.env.NODE_ENV

if (!nodeEnv) {
  nodeEnv = 'production'
}

const config: NuxtConfig = {
  target: 'static',
  head: {
    titleTemplate: '%s - 柠喵的问题箱',
    title: '柠喵的问题箱',
    htmlAttrs: {
      lang: 'zh'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'LemonNeko\'s Question Box' }
    ],
    link: [
      { rel: 'icon', type: 'image/png', href: '/favicon.png' }
    ]
  },
  css: [
  ],
  plugins: [
  ],
  components: true,
  buildModules: [
    '@nuxt/typescript-build',
    '@nuxtjs/vuetify'
  ],
  modules: [
    '@nuxtjs/axios'
  ],
  axios: {

  },
  vuetify: {
    theme: {
      dark: false,
      themes: {
        light: {
          primary: '#ff8a65'
        }
      }
    },
    breakpoint: {
      mobileBreakpoint: 'sm'
    }
  },
  build: {
  },
  env: env[nodeEnv]
}

export default config

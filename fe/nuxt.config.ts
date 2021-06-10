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
      { hid: 'description', name: 'description', content: 'LemonNeko\'s Question Box' },
      { name: 'twitter:card', content: 'summary' },
      { name: 'twitter:creator', content: '@lemon_neko_cn' },
      { property: 'og:image', content: 'https://qbox.lemonneko.moe/question_48x48.png' },
      { property: 'og:title', content: '柠喵的问题箱' },
      { property: 'og:description', content: '在这里，大家可以匿名向柠喵提问哦' }
    ],
    link: [
      { rel: 'icon', type: 'image/png', href: '/favicon.png' }
    ]
  },
  css: [],
  plugins: [
  ],
  components: true,
  buildModules: [
    '@nuxt/typescript-build',
    '@nuxtjs/vuetify'
  ],
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/robots'
  ],
  robots: {
    UserAgent: '*'
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
      mobileBreakpoint: 'xs'
    }
  },
  build: {
  },
  env: env[nodeEnv],
  server: {
    host: '0.0.0.0'
  }
}

export default config

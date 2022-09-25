<template>
  <v-app class="fill-height">
    <v-img :src="formattedBackground" max-height="100%">
      <v-app-bar dark class="blur">
        <span class="text-h5">柠喵的问题箱</span>
        <v-spacer />
        <v-switch
          v-model="getBackgroundFromBing"
          :loading="loadingBingWallpaper"
          light
          inset
          color="white"
          class="switch-on-app-bar"
          prepend-icon="mdi-microsoft-bing"
        />
        <v-btn v-if="$vuetify.breakpoint.mobile" icon @click="openInfo = true">
          <v-icon>mdi-information</v-icon>
        </v-btn>
      </v-app-bar>
      <v-main>
        <nuxt />
        <neko-footer v-if="!$vuetify.breakpoint.mobile" />
      </v-main>
    </v-img>
    <v-dialog v-model="openInfo" transition="dialog-bottom-transition" fullscreen>
      <v-app-bar flat color="white">
        <v-app-bar-title>关于</v-app-bar-title>
        <v-spacer />
        <v-btn text outlined color="error" @click="openInfo = false">
          关闭
        </v-btn>
      </v-app-bar>
      <v-container class="info-dialog-container">
        <v-row>
          <v-col>
            <v-card outlined>
              <v-card-text>
                <div class="footer-avatar-and-name">
                  <img src="LemonNeko_Avatar.png" height="36" width="36" alt="柠喵的头像" class="footer-avatar">
                  <div class="width-10px" />
                  <span class="footer-name">
                    LemonNeko 柠喵
                  </span>
                </div>
                <div style="height: 10px" />
                <div class="light" style="text-align: center">
                  本项目在 Apache 2.0 许可证下开源
                </div>
                <div style="height: 20px" />
                <v-row class="footer-icons" justify="center">
                  <v-btn icon class="footer-icon" href="https://twitter.com/@lemon_neko_cn" target="_blank">
                    <v-icon color="#1da1f2">
                      mdi-twitter
                    </v-icon>
                  </v-btn>
                  <v-btn icon class="footer-icon" href="https://github.com/LemonNekoGH" target="_blank">
                    <v-icon color="black">
                      mdi-github
                    </v-icon>
                  </v-btn>
                  <v-btn icon class="footer-icon" href="https://t.me/lemonneko" target="_blank">
                    <v-icon color="#179cde">
                      mdi-telegram
                    </v-icon>
                  </v-btn>
                </v-row>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-dialog>
  </v-app>
</template>
<script lang="ts">
import Vue from 'vue'
import { api } from '~/api/api'
import { isSuccess } from '~/api/types'
import NekoFooter from '~/components/neko-footer/index.vue'

export default Vue.extend({
  components: {
    NekoFooter
  },
  data () {
    return {
      background: '',
      getBackgroundFromBing: true,
      loadingBingWallpaper: false,
      openInfo: false
    }
  },
  computed: {
    formattedBackground (): string {
      if (!this.background) {
        if ((this as any).$vuetify.breakpoint.mobile) {
          return '/placeholder_mobile.jpeg'
        }
        return '/placeholder_desktop.jpeg'
      }
      let react = '1920x1080'
      if ((this as any).$vuetify.breakpoint.mobile) {
        react = '480x800'
      }
      return `https://www.bing.com${this.background}_${react}.jpg`
    }
  },
  watch: {
    getBackgroundFromBing () {
      this.getBackground()
    }
  },
  mounted () {
    this.getBackground()
    const media = window.matchMedia('(prefers-color-scheme: dark)')
    media.addEventListener('change', this.changeColorScheme)
  },
  methods: {
    async getBackground () {
      this.loadingBingWallpaper = true
      if (!this.getBackgroundFromBing) {
        this.background = ''
        return
      }
      const res = await api.be.getWallpaper()
      if (isSuccess(res)) {
        this.background = res.data.images[0].url
      }
    },
    changeColorScheme () {
      let link: HTMLLinkElement | null = document.querySelector("link[rel~='icon']")
      if (!link) {
        link = document.createElement('link')
      }
      link.rel = 'icon'
      link.type = 'image/png'
      if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
        link.href = '/favicon-dark.png'
      } else {
        link.href = '/favicon.png'
      }
    },
    copyEmail () {
      navigator.clipboard.writeText('chheese048@gmail.com')
      this.$msg.success('复制邮箱成功')
    }
  }
})
</script>
<style lang="less" src="../static/app.less"></style>

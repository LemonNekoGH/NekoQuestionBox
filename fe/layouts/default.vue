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
    <v-snackbar v-model="snackbar.show" elevation="0" :color="snackbar.color" top app>
      {{ snackbar.text }}
    </v-snackbar>
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
import NekoFooter from '~/components/neko-footer/index.vue'

export default Vue.extend({
  components: {
    NekoFooter
  },
  data () {
    return {
      background: '',
      getBackgroundFromBing: true,
      snackbar: {
        show: false,
        color: '',
        text: ''
      },
      loadingBingWallpaper: false,
      openInfo: false
    }
  },
  computed: {
    formattedBackground (): string {
      if (!this.background) {
        if (this.$vuetify.breakpoint.mobile) {
          return '/placeholder_mobile.jpeg'
        }
        return '/placeholder_desktop.jpeg'
      }
      let react = '1920x1080'
      if (this.$vuetify.breakpoint.mobile) {
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
    getBackground () {
      this.loadingBingWallpaper = true
      if (!this.getBackgroundFromBing) {
        this.background = ''
        return
      }
      api.be.getWallpaper().then((res) => {
        if (res === 'api-failed') {
          this.snackbar.color = 'warning'
          this.snackbar.text = '获取 Bing 壁纸失败，Bing 壁纸的 API 可能发生了变动'
          this.snackbar.show = true
          this.getBackgroundFromBing = false
        } else if (res === 'api-error') {
          this.snackbar.color = 'error'
          this.snackbar.text = '获取 Bing 壁纸失败，服务器可能没有准备好'
          this.snackbar.show = true
          this.getBackgroundFromBing = false
        } else {
          this.background = res
        }
      }).finally(() => {
        this.loadingBingWallpaper = false
      })
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
      this.snackbar.color = 'success'
      this.snackbar.text = '已复制柠喵的邮箱'
      this.snackbar.show = true
    }
  }
})
</script>
<style lang="less" src="../static/app.less"></style>

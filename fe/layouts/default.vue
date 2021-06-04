<template>
  <v-app>
    <v-img :src="formattedBackground">
      <v-app-bar dark class="blur">
        <v-app-bar-title>柠喵的问题箱</v-app-bar-title>
      </v-app-bar>
      <v-main>
        <v-container>
          <nuxt />
        </v-container>
      </v-main>
      <v-footer
        dark
        app
        class="blur"
      >
        <span>LemonNeko &copy; {{ new Date().getFullYear() }}</span>
      </v-footer>
    </v-img>
  </v-app>
</template>
<script lang="ts">
import Vue from 'vue'
import { api } from '~/api/api'

export default Vue.extend({
  data () {
    return {
      background: ''
    }
  },
  computed: {
    formattedBackground (): string {
      if (!this.background) {
        return ''
      }
      let react = '1920x1080'
      if (this.$vuetify.breakpoint.mobile) {
        react = '480x800'
      }
      return `https://www.bing.com${this.background}_${react}.jpg`
    }
  },
  mounted () {
    api.be.getWallpaper().then((res) => {
      this.background = res
    })
  }
})
</script>
<style lang="less" src="../static/app.less"></style>

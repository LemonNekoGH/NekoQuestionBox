<template>
  <v-app class="fill-height">
    <v-img :src="formattedBackground" max-height="100%">
      <v-app-bar dark class="blur">
        <span class="text-h5">柠喵的问题箱</span>
        <v-spacer />
        <v-switch
          v-model="getBackgroundFromBing"
          light
          inset
          color="white"
          class="switch-on-app-bar"
          prepend-icon="mdi-microsoft-bing"
        />
      </v-app-bar>
      <v-main>
        <nuxt />
      </v-main>
    </v-img>
  </v-app>
</template>
<script lang="ts">
import Vue from 'vue'
import { api } from '~/api/api'

export default Vue.extend({
  data () {
    return {
      background: '',
      getBackgroundFromBing: true
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
  },
  methods: {
    getBackground () {
      if (!this.getBackgroundFromBing) {
        this.background = ''
        return
      }
      api.be.getWallpaper().then((res) => {
        this.background = res
      })
    }
  }
})
</script>
<style lang="less" src="../static/app.less"></style>

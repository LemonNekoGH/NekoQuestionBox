<template>
  <v-row v-if="!isServerAvailable">
    <v-col>
      <v-card dark class="blur" outlined rounded="lg">
        <v-card-text>
          <div class="flex-box align-items">
            <span class="flex-1">
              服务器似乎没有准备好
            </span>
            <v-btn icon :loading="refreshingAvailable" @click="checkServer">
              <v-icon>mdi-refresh</v-icon>
            </v-btn>
          </div>
        </v-card-text>
      </v-card>
    </v-col>
  </v-row>
  <v-row v-else>
    <v-col>
      <v-card dark class="blur" rounded="lg">
        <v-card-title>开始提问</v-card-title>
        <v-divider />
        <v-card-actions class="send-question-card-content">
          <div>
            <v-text-field
              v-model="question"
              aria-autocomplete="none"
              :dense="$vuetify.breakpoint.mobile"
              outlined
              label="在这里输入想要问柠喵的问题"
              prepend-icon="mdi-help-circle-outline"
            />
          </div>
          <div class="flex-box">
            <v-text-field
              v-model="captchaValue"
              aria-autocomplete="none"
              prepend-icon="mdi-numeric"
              :dense="$vuetify.breakpoint.mobile"
              class="flex-1"
              outlined
              label="在这里输入右侧图案中的数字"
            />
            <div class="width-10px" />
            <v-tooltip bottom>
              <template
                #activator="{ attr, on }"
              >
                <v-btn
                  v-show="captchaLoadFailed"
                  color="
                primary"
                  outlined
                  v-bind="attr"
                  class="captcha-reload-btn"
                  v-on="on"
                  @click="getCaptchaId"
                >
                  <v-icon left>
                    mdi-refresh
                  </v-icon>
                  加载失败
                </v-btn>
              </template>
              重新获取
            </v-tooltip>
            <v-tooltip
              bottom
            >
              <template #activator="{ attr, on }">
                <v-img
                  v-show="!captchaLoadFailed"
                  v-bind="attr"
                  :src="captchaUrl"
                  :max-height="$vuetify.breakpoint.mobile ? 40 : 56"
                  max-width="112"
                  class="captcha-img"
                  @click="getCaptchaId"
                  v-on="on"
                />
              </template>
              点击重新获取
            </v-tooltip>
          </div>
        </v-card-actions>
        <v-divider />
        <v-card-actions>
          <v-spacer />
          <v-btn text outlined :loading="submitting" @click="submitQuestion">
            提问
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
    <v-snackbar v-model="snackbar.show" elevation="0" :color="snackbar.color" top app>
      {{ snackbar.text }}
    </v-snackbar>
  </v-row>
</template>

<script lang="ts">
import Vue from 'vue'
import { api } from '~/api/api'

type SnackbarColor = 'success' | 'warning' | 'error' | ''

interface ComponentData {
  isServerAvailable: boolean
  refreshingAvailable: boolean
  question: string
  captchaId: string
  captchaValue: string
  captchaLoadFailed: boolean
  snackbar: {
    show: boolean
    color: SnackbarColor
    text: string
  }
  submitting: boolean
}

export default Vue.extend({
  data (): ComponentData {
    return {
      isServerAvailable: false,
      refreshingAvailable: false,
      question: '',
      captchaId: '',
      captchaValue: '',
      snackbar: {
        show: false,
        color: '',
        text: ''
      },
      submitting: false,
      captchaLoadFailed: false
    }
  },
  computed: {
    captchaUrl (): string {
      if (!this.captchaId) {
        return ''
      }
      return process.env.baseApiUrl + '/captcha-image?id=' + this.captchaId
    }
  },
  mounted () {
    this.checkServer()
    this.getCaptchaId()
  },
  methods: {
    async checkServer () {
      this.refreshingAvailable = true
      this.isServerAvailable = await api.be.isServerAvailable()
      this.refreshingAvailable = false
    },
    async getCaptchaId () {
      const id = await api.be.getCaptchaId()
      if (!id) {
        console.log('captcha load failed')
        this.captchaLoadFailed = true
        return
      }
      this.captchaLoadFailed = false
      this.captchaId = id
    },
    async submitQuestion () {
      if (this.question === '') {
        this.showSnackbar('没有输入问题哦', 'error')
        return
      }
      if (this.captchaValue === '') {
        this.showSnackbar('没有输入验证码哦', 'error')
      }

      this.submitting = true
      const statusCode = await api.be.submitQuestion({
        question: this.question,
        captchaId: this.captchaId,
        captchaValue: this.captchaValue
      })
      switch (statusCode) {
        case 200:
          this.reset()
          this.showSnackbar('提问成功啦，等待柠喵回复吧', 'success')
          break
        case 406:
          this.showSnackbar('验证失败，请点击验证码图案再来一次', 'warning')
          break
        case 500:
          this.showSnackbar('服务器出了点问题，可能过一会就好了', 'error')
          break
      }
      this.submitting = false
    },
    showSnackbar (text: string, color: SnackbarColor) {
      if (color === '') { return }
      this.snackbar.text = text
      this.snackbar.color = color
      this.snackbar.show = true
    },
    reset () {
      this.captchaValue = ''
      this.question = ''
      this.getCaptchaId()
    }
  }
})
</script>
<style lang="less" src="../static/app.less"></style>
<style lang="less" scoped>
.captcha-img {
  transition: all 250ms;
  border: 1px rgba(255,255,255,0.27) solid;
  border-radius: 4px;
  &:hover {
    border: 1px rgba(255,255,255,0.87) solid;
  }
}
.send-question-card-content {
  display: block;
}
@media screen and (max-width: 600px) {
  .captcha-reload-btn {
    height: 40px;
  }
}
</style>

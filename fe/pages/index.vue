<template>
  <v-row v-if="!isServerAvailable">
    <v-col>
      <v-card outlined>
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
      <v-card outlined>
        <v-card-title>开始提问</v-card-title>
        <v-divider />
        <v-card-actions class="send-question-card-content">
          <div>
            <v-text-field v-model="question" :dense="$vuetify.breakpoint.mobile" outlined label="在这里输入想要问柠喵的问题" />
          </div>
          <div class="flex-box">
            <v-text-field v-model="question" :dense="$vuetify.breakpoint.mobile" class="flex-1" outlined label="在这里输入看到的验证码" />
            <div class="width-10px" />
            <v-tooltip bottom>
              <template #activator="{ attr, on }">
                <v-img
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
          <v-btn text outlined>
            提问
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-col>
  </v-row>
</template>

<script lang="ts">
import Vue from 'vue'
import { api } from '~/api/api'

interface ComponentData {
  isServerAvailable: boolean
  refreshingAvailable: boolean
  question: string
  captchaId: string
  captchaValue: string
}

interface QuestionData {
  question: string
  captchaValue: string
}

export default Vue.extend({
  data (): ComponentData {
    return {
      isServerAvailable: false,
      refreshingAvailable: false,
      question: '',
      captchaId: '',
      captchaValue: ''
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
      this.captchaId = await api.be.getCaptchaId()
    }
  }
})
</script>
<style lang="less" src="../static/app.less"></style>
<style lang="less" scoped>
.captcha-img {
  transition: all 250ms;
  border: 1px rgba(0,0,0,0.5) solid;
  border-radius: 3px;
  &:hover {
    border: 1px rgba(0,0,0,0.87) solid;
  }
}
.send-question-card-content {
  display: block;
}
</style>

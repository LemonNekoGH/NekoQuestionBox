<template>
  <v-container class="overflow">
    <v-row v-if="!isServerAvailable">
      <v-col>
        <v-card dark class="blur" rounded="lg">
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
                label="想要问柠喵的问题"
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
                label="右侧图案的数字"
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
    </v-row>
    <v-row v-for="(item, index) in questionsWithDefault" :key="index" :dense="$vuetify.breakpoint.mobile">
      <v-col>
        <v-card class="blur" dark rounded="lg">
          <v-card-text>
            <div :class="$vuetify.breakpoint.mobile || item.question.length > 10 ? '' : 'flex-box align-items'">
              <div class="text-h6" :class="$vuetify.breakpoint.mobile ? '' : 'flex-1'">
                提问：{{ item.question }}
              </div>
              <div class="text-body-1">
                {{ formatTime(item.time) }}
              </div>
            </div>
          </v-card-text>
          <v-divider />
          <v-card-text>
            <div :class="$vuetify.breakpoint.mobile || item.question.length > 10 ? '' : 'flex-box align-items'">
              <div class="text-h6" :class="$vuetify.breakpoint.mobile ? '' : 'flex-1'">
                {{ item.answer }}
              </div>
              <div v-if="item.answerTime" class="text-body-1">
                {{ formatTime(item.answerTime) }}
              </div>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue'
import moment from 'moment'
import { api, Question } from '~/api/api'
import { isSuccess } from '~/api/types'

interface ComponentData {
  isServerAvailable: boolean
  refreshingAvailable: boolean
  question: string
  captchaId: string
  captchaValue: string
  captchaLoadFailed: boolean
  submitting: boolean
  questions: Question[]
}

export default Vue.extend({
  data (): ComponentData {
    return {
      isServerAvailable: false,
      refreshingAvailable: false,
      question: '',
      captchaId: '',
      captchaValue: '',
      submitting: false,
      captchaLoadFailed: false,
      questions: []
    }
  },
  computed: {
    captchaUrl (): string {
      if (!this.captchaId) {
        return ''
      }
      return process.env.baseApiUrl + '/captcha-image?id=' + this.captchaId
    },
    questionsWithDefault (): Question[] {
      const ret: Question[] = []
      this.questions.forEach((it) => {
        ret.push({
          ...it,
          answer: it.answer ? it.answer : '还没有回答'
        })
      })
      return ret
    }
  },
  mounted () {
    this.checkServer()
    this.getCaptchaId()
    this.getQuestions()
  },
  methods: {
    async checkServer () {
      this.refreshingAvailable = true
      this.isServerAvailable = await api.be.isServerAvailable()
      this.refreshingAvailable = false
    },
    async getCaptchaId () {
      const resp = await api.be.getCaptchaId()
      if (isSuccess(resp)) {
        this.captchaLoadFailed = false
        this.captchaId = resp.data
        return
      }
      this.captchaLoadFailed = true
    },
    async submitQuestion () {
      if (this.question === '') {
        this.$msg.error('没有输入问题哦')
        return
      }
      if (this.captchaValue === '') {
        this.$msg.error('没有输入验证码哦')
      }

      this.submitting = true
      const resp = await api.be.submitQuestion({
        question: this.question,
        id: this.captchaId,
        value: this.captchaValue
      })
      if (isSuccess(resp)) {
        this.reset()
        this.getQuestions()
        this.$msg.success('问题提交成功')
        this.submitting = false
        return
      }

      switch (resp.code) {
        case 40601:
          this.$msg.warning('校验失败，请重试')
          break
        case 40901:
          this.$msg.warning('相同的问题已存在')
          break
        case 50001:
          this.$msg.error('服务器内部错误')
          break
      }
      this.submitting = false
    },
    reset () {
      this.captchaValue = ''
      this.question = ''
      this.getCaptchaId()
    },
    async getQuestions () {
      const res = await api.be.getQuestions()
      if (isSuccess(res)) {
        this.questions = res.data
        return
      }
      this.$msg.error('获取问题集时出错')
    },
    formatTime (time: number): string {
      return moment(time).format('YYYY 年 M 月 D 日 HH:mm:ss')
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
    background: white;
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

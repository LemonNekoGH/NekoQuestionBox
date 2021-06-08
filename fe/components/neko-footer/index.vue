<template>
  <v-footer class="blur footer" dark app>
    <v-row class="footer-row" justify="center">
      <v-col align-self="center" :cols="$vuetify.breakpoint.mobile ? 'auto' : 8">
        <div class="footer-avatar-and-name">
          <img src="LemonNeko_Avatar.png" height="36" width="36" alt="柠喵的头像" class="footer-avatar">
          <div class="width-10px" />
          <span class="footer-name">
            LemonNeko 柠喵
          </span>
        </div>
        <div style="height: 10px" />
        <div class="light">
          本项目在 Apache 2.0 许可证下开源
        </div>
      </v-col>
      <v-col v-if="!$vuetify.breakpoint.mobile" align-self="center">
        <v-btn icon class="footer-icon" href="https://twitter.com/@lemon_neko_cn" target="_blank">
          <v-icon>mdi-twitter</v-icon>
        </v-btn>
        <v-btn icon class="footer-icon" href="https://t.me/lemonneko" target="_blank">
          <v-icon>mdi-telegram</v-icon>
        </v-btn>
        <v-btn icon class="footer-icon" href="https://github.com/LemonNekoGH" target="_blank">
          <v-icon>mdi-github</v-icon>
        </v-btn>
        <v-tooltip top>
          <template #activator="{attr, on}">
            <v-btn icon class="footer-icon" v-bind="attr" v-on="on" @click="copyEmail">
              <v-icon>mdi-email</v-icon>
            </v-btn>
          </template>
          chheese048@gmail.com
        </v-tooltip>
      </v-col>
    </v-row>
    <v-row v-if="$vuetify.breakpoint.mobile" justify="center" align="center">
      <v-btn icon class="footer-icon" href="https://twitter.com/@lemon_neko_cn" target="_blank">
        <v-icon>mdi-twitter</v-icon>
      </v-btn>
      <v-btn icon class="footer-icon" href="https://t.me/lemonneko" target="_blank">
        <v-icon>mdi-telegram</v-icon>
      </v-btn>
      <v-btn icon class="footer-icon" href="https://github.com/LemonNekoGH" target="_blank">
        <v-icon>mdi-github</v-icon>
      </v-btn>
      <v-tooltip top>
        <template #activator="{attr, on}">
          <v-btn icon class="footer-icon" v-bind="attr" v-on="on" @click="copyEmail">
            <v-icon>mdi-email</v-icon>
          </v-btn>
        </template>
        chheese048@gmail.com
      </v-tooltip>
    </v-row>
    <v-snackbar v-model="copySuccess" color="success" timeout="1500">
      邮箱已复制到剪贴板
    </v-snackbar>
  </v-footer>
</template>
<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  data () {
    return {
      copySuccess: false
    }
  },
  methods: {
    copyEmail () {
      navigator.clipboard.writeText('chheese048@gmail.com')
      this.copySuccess = true
    }
  }
})
</script>
<style lang="less" src="../../static/app.less"></style>
<style lang="less" scoped>
.footer {
  padding: 16px;
}
.footer-row {
  color: rgba(255,255,255,.8);
  font-family: 'JetBrains Mono',sans-serif;
}
.footer-avatar {
  border-radius: 5px;
}
.footer-name {
  font-size: 1.5em;
  font-weight: 400;
}
.footer-avatar-and-name {
  display: flex;
  align-items: center;
}
.footer-icon {
  color: rgba(255,255,255,.6) !important;
  &:hover {
    color: white !important;
  }
}
@media screen and (max-width: 600px) {
  .footer-avatar-and-name {
    justify-content: center;
  }
}
</style>

<template>
  <v-app>
    <v-app-bar app dark flat color="primary">
      <v-app-bar-title>柠喵的问题箱</v-app-bar-title>
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>
    <v-footer
      app
    >
      <span>&copy; {{ new Date().getFullYear() }}</span>
    </v-footer>
  </v-app>
</template>
<script>
export default {
  data () {
    return {
      drawer: false
    }
  }
}
</script>

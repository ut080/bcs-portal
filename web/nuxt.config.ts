// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    css: [
      '~/assets/css/fonts.css',
    ],
  devtools: { enabled: true },
  modules: [
      '@nuxt/content',
      '@nuxtjs/tailwindcss',
  ]
})

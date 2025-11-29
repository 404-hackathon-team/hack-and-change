// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2025-07-15',
  devtools: {
    enabled: true,

    timeline: {
      enabled: true
    }
  },

  modules: [
    '@nuxt/eslint',
    '@nuxt/fonts',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxt/test-utils',
    '@formkit/auto-animate',
    'dayjs-nuxt',
    '@unocss/nuxt',
    '@vite-pwa/nuxt',
    '@nuxtjs/mdc'
  ],

    css: ['@/assets/fonts.css', '@/assets/styles.css'],

    runtimeConfig: {
        public: {
            apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',
        },
    },
})
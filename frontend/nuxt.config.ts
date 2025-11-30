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

    routeRules: {
        '/remote/api/**': {
            proxy: {
                to: 'http://localhost:8080/api/v1/**',
            },
        },
    },
})
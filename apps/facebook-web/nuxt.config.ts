// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  devServer: {
    port: 5173,
  },
  modules: ['@unocss/nuxt', '@hebilicious/vue-query-nuxt', '@vueuse/nuxt', '@vee-validate/nuxt'],
  veeValidate: {
    autoImports: true,
  },
  vueQuery: {
    queryClientOptions: {
      defaultOptions: {
        queries: {
          refetchOnWindowFocus: false,
          retry: false,
        },
      },
    },
  },
});

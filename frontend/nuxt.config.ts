// https://nuxt.com/docs/api/configuration/nuxt-config
//
import tailwindcss from "@tailwindcss/vite";
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  devtools: { enabled: true },
  modules: ["@nuxt/fonts", "shadcn-nuxt"],
  vite: {
    plugins: [tailwindcss()],
  },
  css: ["./app/assets/css/main.css"],
  shadcn: {
    prefix: "Ui",
    componentDir: "@/components/ui",
  },
  typescript: {
    strict: false,
  },
  runtimeConfig: {
    public: {
      BACKEND_URL: import.meta.env.NUXT_BACKEND_URL,
    },
  },
});

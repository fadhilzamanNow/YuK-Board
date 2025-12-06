import { useRuntimeConfig } from "#app";
import axios from "axios";

export default defineNuxtPlugin((nuxtApp) => {
  const baseApi = useRuntimeConfig().public.BACKEND_URL;
  console.log(baseApi);

  const api = axios.create({
    baseURL: baseApi,
    timeout: 10000,
  });
  api.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`;
    return config;
  });

  return {
    provide: {
      baseApi: api,
    },
  };
});

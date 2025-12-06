import axios from "axios";

export default defineNuxtPlugin(() => {
  const baseApi = useRuntimeConfig().public.BACKEND_URL;

  const api = axios.create({
    baseURL: baseApi,
    timeout: 10000,
  });

  api.interceptors.request.use((config) => {
    if (import.meta.client) {
      const token = localStorage.getItem("authToken");
      if (token) config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  });

  return {
    provide: {
      baseApi: api,
    },
  };
});

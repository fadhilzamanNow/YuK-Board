import { AxiosError } from "axios";
import { baseApi } from "./api";

export async function postLogin(params: LoginParams): Promise<LoginResponse> {
  try {
    const response = await baseApi().post("/login", params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function postRegister(params: RegisterParams): Promise<LoginResponse> {
  try {
    const response = await baseApi().post("/register", params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function getMe(): Promise<MeResponse> {
  try {
    const response = await baseApi().get("/me");
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

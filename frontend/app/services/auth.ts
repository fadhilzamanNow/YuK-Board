import { AxiosError } from "axios";
import { baseApi } from "./api";

export async function postLogin(params: LoginParams) {
  try {
    const response = await baseApi().post("/login", {
      email: params.email,
      password: params.password,
    });
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError) {
      throw err.response.data as CustomErrorResponse<{
        message: string;
        code: number;
      }>;
    }
  }
}

export async function postRegister(params: RegisterParams) {
  try {
    const response = await baseApi().post("/register", {
      email: params.email,
      password: params.password,
      name: params.name,
    });
    return response.data as LoginResponse;
  } catch (err) {
    if (err instanceof AxiosError) {
      throw err.response.data as CustomErrorResponse<{
        message: string;
        code: number;
      }>;
    }
  }
}

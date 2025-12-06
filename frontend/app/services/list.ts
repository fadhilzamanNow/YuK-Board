import { AxiosError } from "axios";
import { baseApi } from "./api";

export async function getLists(): Promise<ListsResponse> {
  try {
    const response = await baseApi().get("/lists");
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function getList(id: string): Promise<ListResponse> {
  try {
    const response = await baseApi().get(`/lists/${id}`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function createList(params: CreateListParams): Promise<ListResponse> {
  try {
    const response = await baseApi().post("/lists", params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function updateList(id: string, params: UpdateListParams): Promise<ListResponse> {
  try {
    const response = await baseApi().put(`/lists/${id}`, params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function deleteList(id: string): Promise<{ message: string }> {
  try {
    const response = await baseApi().delete(`/lists/${id}`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

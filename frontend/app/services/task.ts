import { AxiosError } from "axios";
import { baseApi } from "./api";

export async function getTasks(listId: string): Promise<TasksResponse> {
  try {
    const response = await baseApi().get(`/lists/${listId}/tasks`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function createTask(listId: string, params: CreateTaskParams): Promise<TaskResponse> {
  try {
    const response = await baseApi().post(`/lists/${listId}/tasks`, params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function updateTask(listId: string, taskId: string, params: UpdateTaskParams): Promise<TaskResponse> {
  try {
    const response = await baseApi().put(`/lists/${listId}/tasks/${taskId}`, params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function deleteTask(listId: string, taskId: string): Promise<{ message: string }> {
  try {
    const response = await baseApi().delete(`/lists/${listId}/tasks/${taskId}`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function updateTaskStatus(listId: string, taskId: string): Promise<TaskResponse> {
  try {
    const response = await baseApi().patch(`/lists/${listId}/tasks/${taskId}/status`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

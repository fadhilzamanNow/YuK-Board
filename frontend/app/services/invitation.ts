import { AxiosError } from "axios";
import { baseApi } from "./api";

export async function getMyInvitations(): Promise<InvitationsResponse> {
  try {
    const response = await baseApi().get("/invitations");
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function inviteUser(listId: string, params: InviteParams): Promise<InvitationResponse> {
  try {
    const response = await baseApi().post(`/lists/${listId}/invite`, params);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function acceptInvitation(id: string): Promise<{ message: string }> {
  try {
    const response = await baseApi().post(`/invitations/${id}/accept`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function declineInvitation(id: string): Promise<{ message: string }> {
  try {
    const response = await baseApi().post(`/invitations/${id}/decline`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

export async function cancelInvitation(id: string): Promise<{ message: string }> {
  try {
    const response = await baseApi().delete(`/invitations/${id}`);
    return response.data;
  } catch (err) {
    if (err instanceof AxiosError && err.response) {
      throw err.response.data as CustomErrorResponse;
    }
    throw err;
  }
}

import { useQuery, useMutation } from "@tanstack/vue-query";
import { postLogin, postRegister, getMe } from "~/services/auth";

export function useQueryMe() {
  return useQuery({
    queryKey: ["me"],
    queryFn: getMe,
  });
}

export function useMutationLogin() {
  return useMutation({
    mutationFn: (params: LoginParams) => postLogin(params),
  });
}

export function useMutationRegister() {
  return useMutation({
    mutationFn: (params: RegisterParams) => postRegister(params),
  });
}

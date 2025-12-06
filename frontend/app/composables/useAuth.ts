import { useMutation } from "@tanstack/vue-query";
import { postLogin, postRegister } from "~/services/auth";

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

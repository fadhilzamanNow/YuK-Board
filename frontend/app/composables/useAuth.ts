import { useMutation } from "@tanstack/vue-query";
import { postLogin } from "~/services/auth";

export function useMutationLogin() {
  return useMutation({
    mutationFn: async ({ email, password }: LoginParams) =>
      await postLogin({ email, password }),
  });
}

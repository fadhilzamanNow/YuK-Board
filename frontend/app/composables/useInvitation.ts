import { useQuery, useMutation, useQueryClient } from "@tanstack/vue-query";
import { getMyInvitations, inviteUser, acceptInvitation, declineInvitation, cancelInvitation } from "~/services/invitation";

export function useQueryInvitations() {
  return useQuery({
    queryKey: ["invitations"],
    queryFn: getMyInvitations,
  });
}

export function useMutationInviteUser() {
  return useMutation({
    mutationFn: ({ listId, email }: { listId: string; email: string }) => inviteUser(listId, { email }),
  });
}

export function useMutationAcceptInvitation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => acceptInvitation(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["invitations"] });
      queryClient.invalidateQueries({ queryKey: ["lists"] });
    },
  });
}

export function useMutationDeclineInvitation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => declineInvitation(id),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["invitations"] }),
  });
}

export function useMutationCancelInvitation() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => cancelInvitation(id),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["invitations"] }),
  });
}

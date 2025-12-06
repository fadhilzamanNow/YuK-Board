import { useQuery, useMutation, useQueryClient } from "@tanstack/vue-query";
import { getLists, getList, createList, updateList, deleteList } from "~/services/list";

export function useQueryLists() {
  return useQuery({
    queryKey: ["lists"],
    queryFn: getLists,
  });
}

export function useQueryList(id: Ref<string | undefined> | ComputedRef<string | undefined>) {
  return useQuery({
    queryKey: ["list", id],
    queryFn: () => getList(id.value!),
    enabled: () => !!id.value,
  });
}

export function useMutationCreateList() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (params: CreateListParams) => createList(params),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["lists"] }),
  });
}

export function useMutationUpdateList() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ id, params }: { id: string; params: UpdateListParams }) => updateList(id, params),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["lists"] }),
  });
}

export function useMutationDeleteList() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: (id: string) => deleteList(id),
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["lists"] }),
  });
}

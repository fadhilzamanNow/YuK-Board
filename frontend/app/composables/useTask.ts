import { useQuery, useMutation, useQueryClient } from "@tanstack/vue-query";
import { getTasks, createTask, updateTask, deleteTask, updateTaskStatus } from "~/services/task";

export function useQueryTasks(listId: Ref<string | undefined> | ComputedRef<string | undefined>) {
  return useQuery({
    queryKey: ["tasks", listId],
    queryFn: () => getTasks(listId.value!),
    enabled: () => !!listId.value,
  });
}

export function useMutationCreateTask() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ listId, params }: { listId: string; params: CreateTaskParams }) => createTask(listId, params),
    onSuccess: (_, { listId }) => queryClient.invalidateQueries({ queryKey: ["tasks", listId] }),
  });
}

export function useMutationUpdateTask() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ listId, taskId, params }: { listId: string; taskId: string; params: UpdateTaskParams }) =>
      updateTask(listId, taskId, params),
    onSuccess: (_, { listId }) => queryClient.invalidateQueries({ queryKey: ["tasks", listId] }),
  });
}

export function useMutationDeleteTask() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ listId, taskId }: { listId: string; taskId: string }) => deleteTask(listId, taskId),
    onSuccess: (_, { listId }) => queryClient.invalidateQueries({ queryKey: ["tasks", listId] }),
  });
}

export function useMutationToggleTaskStatus() {
  const queryClient = useQueryClient();
  return useMutation({
    mutationFn: ({ listId, taskId }: { listId: string; taskId: string }) => updateTaskStatus(listId, taskId),
    onSuccess: (_, { listId }) => queryClient.invalidateQueries({ queryKey: ["tasks", listId] }),
  });
}

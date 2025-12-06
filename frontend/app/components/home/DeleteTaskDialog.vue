<script setup lang="ts">
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { toast } from "vue-sonner";

const open = defineModel<boolean>("open", { default: false });
const props = defineProps<{ task?: Task; listId?: string }>();

const { mutateAsync: deleteTask, isPending } = useMutationDeleteTask();

const handleDelete = () => {
  if (!props.task || !props.listId) return;
  deleteTask(
    { listId: props.listId, taskId: props.task.id },
    {
      onSuccess: () => {
        toast.success("Tugas berhasil dihapus");
        open.value = false;
      },
      onError: (error: CustomErrorResponse) => {
        toast.error(error.message || "Gagal menghapus tugas");
      },
    }
  );
};
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Hapus Tugas</DialogTitle>
        <DialogDescription>
          Apakah kamu yakin ingin menghapus "{{ task?.title }}"?
        </DialogDescription>
      </DialogHeader>
      <DialogFooter>
        <Button variant="outline" @click="open = false">Batal</Button>
        <Button variant="destructive" :disabled="isPending" @click="handleDelete">
          {{ isPending ? "Menghapus..." : "Hapus" }}
        </Button>
      </DialogFooter>
    </DialogContent>
  </Dialog>
</template>

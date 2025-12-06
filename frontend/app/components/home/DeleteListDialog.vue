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
const props = defineProps<{ list?: TodoList }>();
const emit = defineEmits<{ deleted: [] }>();

const { mutateAsync: deleteList, isPending } = useMutationDeleteList();

const handleDelete = () => {
  if (!props.list) return;
  deleteList(props.list.id, {
    onSuccess: () => {
      toast.success("Daftar berhasil dihapus");
      open.value = false;
      emit("deleted");
    },
    onError: (error: CustomErrorResponse) => {
      toast.error(error.message || "Gagal menghapus daftar");
    },
  });
};
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Hapus Daftar</DialogTitle>
        <DialogDescription>
          Apakah kamu yakin ingin menghapus "{{ list?.title }}"? Semua tugas di dalamnya akan ikut terhapus.
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

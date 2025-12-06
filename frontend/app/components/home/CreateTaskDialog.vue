<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Field, FieldGroup, FieldLabel } from "@/components/ui/field";
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
const props = defineProps<{ listId?: string }>();

const title = ref("");
const description = ref("");
const errors = ref<Record<string, string>>({});

const { mutateAsync: createTask, isPending } = useMutationCreateTask();

const handleSubmit = () => {
  if (!props.listId) return;
  errors.value = {};
  createTask(
    { listId: props.listId, params: { title: title.value, description: description.value } },
    {
      onSuccess: (response) => {
        toast.success(response.message || "Tugas berhasil dibuat");
        open.value = false;
        title.value = "";
        description.value = "";
      },
      onError: (error: CustomErrorResponse) => {
        toast.error(error.message || "Gagal membuat tugas");
        if (error.errors) errors.value = error.errors;
      },
    }
  );
};
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent>
      <DialogHeader>
        <DialogTitle>Buat Tugas Baru</DialogTitle>
        <DialogDescription>Tambah tugas ke daftar</DialogDescription>
      </DialogHeader>
      <form @submit.prevent="handleSubmit">
        <FieldGroup>
          <Field>
            <FieldLabel for="task-title">Judul</FieldLabel>
            <Input id="task-title" v-model="title" placeholder="Nama tugas" required />
            <span v-if="errors.title" class="text-red-500 text-[10px]">*{{ errors.title }}</span>
          </Field>
          <Field>
            <FieldLabel for="task-description">Deskripsi</FieldLabel>
            <Input id="task-description" v-model="description" placeholder="Deskripsi (opsional)" />
            <span v-if="errors.description" class="text-red-500 text-[10px]">*{{ errors.description }}</span>
          </Field>
        </FieldGroup>
        <DialogFooter class="mt-4">
          <Button type="submit" :disabled="isPending">
            {{ isPending ? "Menyimpan..." : "Simpan" }}
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>

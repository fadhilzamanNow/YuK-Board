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
const props = defineProps<{ task?: Task; listId?: string }>();

const title = ref("");
const description = ref("");
const errors = ref<Record<string, string>>({});

const { mutateAsync: updateTask, isPending } = useMutationUpdateTask();

watch(() => props.task, (task) => {
  if (task) {
    title.value = task.title;
    description.value = task.description ?? "";
  }
}, { immediate: true });

const handleSubmit = () => {
  if (!props.task || !props.listId) return;
  errors.value = {};
  updateTask(
    { listId: props.listId, taskId: props.task.id, params: { title: title.value, description: description.value } },
    {
      onSuccess: (response) => {
        toast.success(response.message || "Tugas berhasil diperbarui");
        open.value = false;
      },
      onError: (error: CustomErrorResponse) => {
        toast.error(error.message || "Gagal memperbarui tugas");
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
        <DialogTitle>Edit Tugas</DialogTitle>
        <DialogDescription>Perbarui informasi tugas</DialogDescription>
      </DialogHeader>
      <form @submit.prevent="handleSubmit">
        <FieldGroup>
          <Field>
            <FieldLabel for="edit-task-title">Judul</FieldLabel>
            <Input id="edit-task-title" v-model="title" placeholder="Nama tugas" required />
            <span v-if="errors.title" class="text-red-500 text-[10px]">*{{ errors.title }}</span>
          </Field>
          <Field>
            <FieldLabel for="edit-task-description">Deskripsi</FieldLabel>
            <Input id="edit-task-description" v-model="description" placeholder="Deskripsi (opsional)" />
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

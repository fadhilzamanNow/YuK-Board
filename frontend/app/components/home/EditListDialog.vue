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
const props = defineProps<{ list?: TodoList }>();

const title = ref("");
const description = ref("");
const errors = ref<Record<string, string>>({});

const { mutateAsync: updateList, isPending } = useMutationUpdateList();

watch(() => props.list, (list) => {
  if (list) {
    title.value = list.title;
    description.value = list.description ?? "";
  }
}, { immediate: true });

const handleSubmit = () => {
  if (!props.list) return;
  errors.value = {};
  updateList(
    { id: props.list.id, params: { title: title.value, description: description.value } },
    {
      onSuccess: (response) => {
        toast.success(response.message || "Daftar berhasil diperbarui");
        open.value = false;
      },
      onError: (error: CustomErrorResponse) => {
        toast.error(error.message || "Gagal memperbarui daftar");
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
        <DialogTitle>Edit Daftar</DialogTitle>
        <DialogDescription>Perbarui informasi daftar tugas</DialogDescription>
      </DialogHeader>
      <form @submit.prevent="handleSubmit">
        <FieldGroup>
          <Field>
            <FieldLabel for="edit-title">Judul</FieldLabel>
            <Input id="edit-title" v-model="title" placeholder="Nama daftar" required />
            <span v-if="errors.title" class="text-red-500 text-[10px]">*{{ errors.title }}</span>
          </Field>
          <Field>
            <FieldLabel for="edit-description">Deskripsi</FieldLabel>
            <Input id="edit-description" v-model="description" placeholder="Deskripsi (opsional)" />
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

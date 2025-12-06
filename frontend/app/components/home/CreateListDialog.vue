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

const title = ref("");
const description = ref("");
const errors = ref<Record<string, string>>({});

const { mutateAsync: createList, isPending } = useMutationCreateList();

const handleSubmit = () => {
  errors.value = {};
  createList(
    { title: title.value, description: description.value },
    {
      onSuccess: (response) => {
        toast.success(response.message || "Daftar berhasil dibuat");
        open.value = false;
        title.value = "";
        description.value = "";
      },
      onError: (error: CustomErrorResponse) => {
        toast.error(error.message || "Gagal membuat daftar");
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
        <DialogTitle>Buat Daftar Baru</DialogTitle>
        <DialogDescription>Buat daftar tugas untuk tim kamu</DialogDescription>
      </DialogHeader>
      <form @submit.prevent="handleSubmit">
        <FieldGroup>
          <Field>
            <FieldLabel for="title">Judul</FieldLabel>
            <Input id="title" v-model="title" placeholder="Nama daftar" required />
            <span v-if="errors.title" class="text-red-500 text-[10px]">*{{ errors.title }}</span>
          </Field>
          <Field>
            <FieldLabel for="description">Deskripsi</FieldLabel>
            <Input id="description" v-model="description" placeholder="Deskripsi (opsional)" />
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

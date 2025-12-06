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

const email = ref("");
const errors = ref<Record<string, string>>({});

const { mutateAsync: invite, isPending } = useMutationInviteUser();

const handleSubmit = () => {
  if (!props.list) return;
  errors.value = {};
  invite(
    { listId: props.list.id, email: email.value },
    {
      onSuccess: (response) => {
        toast.success(response.message || "Undangan berhasil dikirim");
        open.value = false;
        email.value = "";
      },
      onError: (error: CustomErrorResponse) => {
        toast.error(error.message || "Gagal mengirim undangan");
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
        <DialogTitle>Undang Anggota</DialogTitle>
        <DialogDescription>Undang pengguna ke "{{ list?.title }}"</DialogDescription>
      </DialogHeader>
      <form @submit.prevent="handleSubmit">
        <FieldGroup>
          <Field>
            <FieldLabel for="invite-email">Email</FieldLabel>
            <Input id="invite-email" type="email" v-model="email" placeholder="email@example.com" required />
            <span v-if="errors.email" class="text-red-500 text-[10px]">*{{ errors.email }}</span>
          </Field>
        </FieldGroup>
        <DialogFooter class="mt-4">
          <Button type="submit" :disabled="isPending">
            {{ isPending ? "Mengirim..." : "Kirim Undangan" }}
          </Button>
        </DialogFooter>
      </form>
    </DialogContent>
  </Dialog>
</template>

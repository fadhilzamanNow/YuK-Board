<script setup lang="ts">
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Check, X } from "lucide-vue-next";
import { toast } from "vue-sonner";

const open = defineModel<boolean>("open", { default: false });

const { data, isLoading } = useQueryInvitations();
const { mutateAsync: accept, isPending: isAccepting } = useMutationAcceptInvitation();
const { mutateAsync: decline, isPending: isDeclining } = useMutationDeclineInvitation();

const invitations = computed(() => data.value?.invitations ?? []);

const handleAccept = (id: string) => {
  accept(id, {
    onSuccess: () => toast.success("Undangan diterima"),
    onError: (err: CustomErrorResponse) => toast.error(err.message || "Gagal menerima undangan"),
  });
};

const handleDecline = (id: string) => {
  decline(id, {
    onSuccess: () => toast.success("Undangan ditolak"),
    onError: (err: CustomErrorResponse) => toast.error(err.message || "Gagal menolak undangan"),
  });
};
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent class="max-w-md">
      <DialogHeader>
        <DialogTitle>Undangan</DialogTitle>
        <DialogDescription>Undangan untuk bergabung ke daftar tugas</DialogDescription>
      </DialogHeader>
      <div class="space-y-3 max-h-80 overflow-y-auto">
        <div v-if="isLoading" class="text-muted-foreground text-sm text-center py-4">Memuat...</div>
        <div v-else-if="!invitations.length" class="text-muted-foreground text-sm text-center py-4">
          Tidak ada undangan
        </div>
        <Card v-for="inv in invitations" :key="inv.id" class="flex flex-row flex-wrap items-center justify-between p-4">
          <div>
            <p class="font-medium text-sm">{{ inv.todo_list?.title }}</p>
            <p class="text-muted-foreground text-xs">Dari: {{ inv.inviter?.name }}</p>
          </div>
          <div class="flex gap-2">
            <Button size="sm" :disabled="isAccepting || isDeclining" @click="handleAccept(inv.id)">
              <Check class="size-4 mr-1" /> Terima
            </Button>
            <Button size="sm" variant="outline" :disabled="isAccepting || isDeclining" @click="handleDecline(inv.id)">
              <X class="size-4 mr-1" /> Tolak
            </Button>
          </div>
        </Card>
      </div>
    </DialogContent>
  </Dialog>
</template>

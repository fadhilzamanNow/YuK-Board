<script setup lang="ts">
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { Badge } from "@/components/ui/badge";
import { Crown } from "lucide-vue-next";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";

const open = defineModel<boolean>("open", { default: false });
const props = defineProps<{ members: ListMember[] }>();

const getInitial = (name?: string) => name?.charAt(0).toUpperCase() ?? "?";
</script>

<template>
  <Dialog v-model:open="open">
    <DialogContent class="max-w-sm">
      <DialogHeader>
        <DialogTitle>Anggota</DialogTitle>
        <DialogDescription>Daftar anggota dalam list ini</DialogDescription>
      </DialogHeader>
      <div class="space-y-3 max-h-60 overflow-y-auto">
        <div v-for="member in members" :key="member.id" class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <Avatar class="size-8">
              <AvatarFallback>{{ getInitial(member.user?.name) }}</AvatarFallback>
            </Avatar>
            <div>
              <p class="text-sm font-medium">{{ member.user?.name }}</p>
              <p class="text-xs text-muted-foreground">{{ member.user?.email }}</p>
            </div>
          </div>
          <Badge variant="secondary">
            <Crown v-if="member.role === 'owner'" class="size-3 mr-1" />
            {{ member.role === 'owner' ? 'Owner' : 'Member' }}
          </Badge>
        </div>
      </div>
    </DialogContent>
  </Dialog>
</template>

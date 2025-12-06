<script setup lang="ts">
import { Avatar, AvatarFallback } from "@/components/ui/avatar";

const props = defineProps<{ members: ListMember[]; max?: number }>();
const emit = defineEmits<{ click: [] }>();

const maxShow = computed(() => props.max ?? 3);
const visibleMembers = computed(() => props.members.slice(0, maxShow.value));
const remaining = computed(() => Math.max(0, props.members.length - maxShow.value));

const getInitial = (name?: string) => name?.charAt(0).toUpperCase() ?? "?";
</script>

<template>
  <button class="flex -space-x-2 cursor-pointer hover:opacity-80" @click="emit('click')">
    <Avatar v-for="member in visibleMembers" :key="member.id" class="size-7 border-2 border-background">
      <AvatarFallback class="text-xs">{{ getInitial(member.user?.name) }}</AvatarFallback>
    </Avatar>
    <Avatar v-if="remaining > 0" class="size-7 border-2 border-background">
      <AvatarFallback class="text-xs bg-muted">+{{ remaining }}</AvatarFallback>
    </Avatar>
  </button>
</template>

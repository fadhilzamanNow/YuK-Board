<script setup lang="ts">
import { LogOut } from "lucide-vue-next";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

const props = defineProps<{
  userName: string;
  userEmail?: string;
}>();

const initial = computed(() => props.userName?.charAt(0).toUpperCase() ?? "?");

const handleLogout = () => {
  localStorage.removeItem("authToken");
  navigateTo("/login");
};
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <button class="cursor-pointer">
        <Avatar class="size-8">
          <AvatarFallback>{{ initial }}</AvatarFallback>
        </Avatar>
      </button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end" class="w-56">
      <DropdownMenuLabel>
        <p class="font-medium">{{ userName }}</p>
        <p class="text-xs text-muted-foreground font-normal">{{ userEmail }}</p>
      </DropdownMenuLabel>
      <DropdownMenuSeparator />
      <DropdownMenuItem @click="handleLogout">
        <LogOut class="size-4 mr-2" />
        Keluar
      </DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>

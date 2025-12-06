<script setup lang="ts">
import { GalleryVerticalEnd, Plus, ListTodo, Mail, EllipsisVertical, Pencil, Trash2 } from "lucide-vue-next";
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupAction,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSkeleton,
  SidebarMenuAction,
} from "@/components/ui/sidebar";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";

const props = defineProps<{
  lists: TodoList[];
  selectedListId?: string;
  isLoading?: boolean;
  invitationCount?: number;
}>();

const emit = defineEmits<{
  selectList: [id: string];
  createList: [];
  editList: [list: TodoList];
  deleteList: [list: TodoList];
  openInvitations: [];
}>();
</script>

<template>
  <Sidebar>
    <SidebarHeader>
      <SidebarMenu>
        <SidebarMenuItem>
          <SidebarMenuButton size="lg" as-child>
            <NuxtLink to="/">
              <div class="bg-primary text-primary-foreground flex size-8 items-center justify-center rounded-md">
                <GalleryVerticalEnd class="size-5" />
              </div>
              <span class="font-bold">YukBoard</span>
            </NuxtLink>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarHeader>

    <SidebarContent>
      <SidebarGroup>
        <SidebarGroupLabel>Daftar Tugas</SidebarGroupLabel>
        <SidebarGroupAction @click="emit('createList')">
          <Plus class="size-4" />
        </SidebarGroupAction>
        <SidebarGroupContent>
          <SidebarMenu>
            <template v-if="isLoading">
              <SidebarMenuSkeleton v-for="i in 3" :key="i" />
            </template>
            <template v-else-if="lists.length">
              <SidebarMenuItem v-for="list in lists" :key="list.id">
                <SidebarMenuButton
                  :is-active="selectedListId === list.id"
                  @click="emit('selectList', list.id)"
                >
                  <ListTodo class="size-4" />
                  <span>{{ list.title }}</span>
                </SidebarMenuButton>
                <DropdownMenu>
                  <DropdownMenuTrigger as-child>
                    <SidebarMenuAction>
                      <EllipsisVertical class="size-4" />
                    </SidebarMenuAction>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent side="right" align="start">
                    <DropdownMenuItem @click="emit('editList', list)">
                      <Pencil class="size-4 mr-2" />
                      Edit
                    </DropdownMenuItem>
                    <DropdownMenuItem class="text-destructive" @click="emit('deleteList', list)">
                      <Trash2 class="size-4 mr-2" />
                      Hapus
                    </DropdownMenuItem>
                  </DropdownMenuContent>
                </DropdownMenu>
              </SidebarMenuItem>
            </template>
            <template v-else>
              <p class="text-muted-foreground text-sm px-2">Belum ada daftar</p>
            </template>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>

    <SidebarFooter>
      <SidebarMenu>
        <SidebarMenuItem>
          <SidebarMenuButton @click="emit('openInvitations')">
            <Mail class="size-4" />
            <span>Undangan</span>
            <span v-if="invitationCount" class="ml-auto bg-primary text-primary-foreground text-xs px-2 py-0.5 rounded-full">
              {{ invitationCount }}
            </span>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarFooter>
  </Sidebar>
</template>

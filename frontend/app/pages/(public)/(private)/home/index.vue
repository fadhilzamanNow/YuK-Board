<script setup lang="ts">
import { SidebarProvider, SidebarInset, SidebarTrigger } from "@/components/ui/sidebar";
import AppSidebar from "@/components/home/AppSidebar.vue";
import CreateListDialog from "@/components/home/CreateListDialog.vue";
import EditListDialog from "@/components/home/EditListDialog.vue";
import DeleteListDialog from "@/components/home/DeleteListDialog.vue";
import UserDropdown from "@/components/home/UserDropdown.vue";

const selectedListId = ref<string>();
const showCreateDialog = ref(false);
const showEditDialog = ref(false);
const showDeleteDialog = ref(false);
const showInvitations = ref(false);
const editingList = ref<TodoList>();
const deletingList = ref<TodoList>();

const { data: meData } = useQueryMe();
const { data: listsData, isLoading } = useQueryLists();

const user = computed(() => meData.value?.user);
const lists = computed(() => listsData.value?.lists ?? []);
const selectedList = computed(() => lists.value.find((l) => l.id === selectedListId.value));

watch(lists, (newLists) => {
  if (newLists.length && !selectedListId.value) {
    selectedListId.value = newLists[0].id;
  }
}, { immediate: true });

const handleEditList = (list: TodoList) => {
  editingList.value = list;
  showEditDialog.value = true;
};

const handleDeleteList = (list: TodoList) => {
  deletingList.value = list;
  showDeleteDialog.value = true;
};
</script>

<template>
  <SidebarProvider>
    <AppSidebar
      :lists="lists"
      :selected-list-id="selectedListId"
      :is-loading="isLoading"
      :invitation-count="0"
      @select-list="selectedListId = $event"
      @create-list="showCreateDialog = true"
      @edit-list="handleEditList"
      @delete-list="handleDeleteList"
      @open-invitations="showInvitations = true"
    />
    <SidebarInset>
      <header class="flex h-14 items-center justify-between border-b px-4">
        <div class="flex items-center gap-2">
          <SidebarTrigger />
          <h1 class="font-semibold">{{ selectedList?.title ?? 'Pilih daftar' }}</h1>
        </div>
        <div class="flex items-center gap-3">
          <ClientOnly>
            <ModeToggle />
          </ClientOnly>
          <UserDropdown :user-name="user?.name ?? ''" />
        </div>
      </header>
      <main class="flex-1 p-4">
        <div v-if="!selectedListId" class="text-muted-foreground text-center py-20">
          Pilih atau buat daftar tugas untuk memulai
        </div>
        <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div class="bg-muted/50 rounded-lg p-4">
            <h2 class="font-semibold mb-4">Todo</h2>
            <p class="text-muted-foreground text-sm">Belum ada tugas</p>
          </div>
          <div class="bg-muted/50 rounded-lg p-4">
            <h2 class="font-semibold mb-4">In Progress</h2>
            <p class="text-muted-foreground text-sm">Belum ada tugas</p>
          </div>
          <div class="bg-muted/50 rounded-lg p-4">
            <h2 class="font-semibold mb-4">Done</h2>
            <p class="text-muted-foreground text-sm">Belum ada tugas</p>
          </div>
        </div>
      </main>
    </SidebarInset>

    <CreateListDialog v-model:open="showCreateDialog" />
    <EditListDialog v-model:open="showEditDialog" :list="editingList" />
    <DeleteListDialog v-model:open="showDeleteDialog" :list="deletingList" @deleted="selectedListId = undefined" />
  </SidebarProvider>
</template>

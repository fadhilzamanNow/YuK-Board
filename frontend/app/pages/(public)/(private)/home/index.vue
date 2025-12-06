<script setup lang="ts">
import { SidebarProvider, SidebarInset, SidebarTrigger } from "@/components/ui/sidebar";
import { Button } from "@/components/ui/button";
import { UserPlus, ListTodo } from "lucide-vue-next";
import AppSidebar from "@/components/home/AppSidebar.vue";
import CreateListDialog from "@/components/home/CreateListDialog.vue";
import EditListDialog from "@/components/home/EditListDialog.vue";
import DeleteListDialog from "@/components/home/DeleteListDialog.vue";
import InviteUserDialog from "@/components/home/InviteUserDialog.vue";
import InvitationsDialog from "@/components/home/InvitationsDialog.vue";
import CreateTaskDialog from "@/components/home/CreateTaskDialog.vue";
import EditTaskDialog from "@/components/home/EditTaskDialog.vue";
import DeleteTaskDialog from "@/components/home/DeleteTaskDialog.vue";
import KanbanBoard from "@/components/home/KanbanBoard.vue";
import MembersAvatarStack from "@/components/home/MembersAvatarStack.vue";
import MembersDialog from "@/components/home/MembersDialog.vue";
import UserDropdown from "@/components/home/UserDropdown.vue";

const selectedListId = ref<string>();
const showCreateDialog = ref(false);
const showEditDialog = ref(false);
const showDeleteDialog = ref(false);
const showInviteDialog = ref(false);
const showInvitationsDialog = ref(false);
const showCreateTaskDialog = ref(false);
const showEditTaskDialog = ref(false);
const showDeleteTaskDialog = ref(false);
const showMembersDialog = ref(false);
const editingList = ref<TodoList>();
const deletingList = ref<TodoList>();
const invitingList = ref<TodoList>();
const editingTask = ref<Task>();
const deletingTask = ref<Task>();

const { data: meData } = useQueryMe();
const { data: listsData, isLoading } = useQueryLists();
const { data: invitationsData } = useQueryInvitations();
const { data: listDetailData } = useQueryList(selectedListId);
const { data: tasksData } = useQueryTasks(selectedListId);

const user = computed(() => meData.value?.user);
const lists = computed(() => listsData.value?.lists ?? []);
const selectedList = computed(() => lists.value.find((l) => l.id === selectedListId.value));
const invitationCount = computed(() => invitationsData.value?.invitations?.length ?? 0);
const tasks = computed(() => tasksData.value?.tasks ?? []);
const members = computed(() => listDetailData.value?.list?.members ?? []);

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

const handleInviteUser = (list: TodoList) => {
  invitingList.value = list;
  showInviteDialog.value = true;
};

const handleEditTask = (task: Task) => {
  editingTask.value = task;
  showEditTaskDialog.value = true;
};

const handleDeleteTask = (task: Task) => {
  deletingTask.value = task;
  showDeleteTaskDialog.value = true;
};
</script>

<template>
  <SidebarProvider>
    <AppSidebar
      :lists="lists"
      :selected-list-id="selectedListId"
      :is-loading="isLoading"
      :invitation-count="invitationCount"
      @select-list="selectedListId = $event"
      @create-list="showCreateDialog = true"
      @edit-list="handleEditList"
      @delete-list="handleDeleteList"
      @open-invitations="showInvitationsDialog = true"
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
          <UserDropdown :user-name="user?.name ?? ''" :user-email="user?.email" />
        </div>
      </header>
      <main class="flex-1 p-4">
        <div v-if="!selectedListId" class="flex flex-col items-center justify-center py-20">
          <ListTodo class="size-16 text-muted-foreground/50 mb-4" />
          <p class="text-muted-foreground">Pilih atau buat daftar tugas untuk memulai</p>
        </div>
        <div v-else>
          <div class="flex items-center justify-between mb-4">
            <Button size="sm" variant="outline" @click="handleInviteUser(selectedList!)">
              <UserPlus class="size-4 mr-2" /> Undang
            </Button>
            <MembersAvatarStack :members="members" @click="showMembersDialog = true" />
          </div>
          <KanbanBoard
            :list-id="selectedListId"
            :tasks="tasks"
            @create-task="showCreateTaskDialog = true"
            @edit-task="handleEditTask"
            @delete-task="handleDeleteTask"
          />
        </div>
      </main>
    </SidebarInset>

    <CreateListDialog v-model:open="showCreateDialog" />
    <EditListDialog v-model:open="showEditDialog" :list="editingList" />
    <DeleteListDialog v-model:open="showDeleteDialog" :list="deletingList" @deleted="selectedListId = undefined" />
    <InviteUserDialog v-model:open="showInviteDialog" :list="invitingList" />
    <InvitationsDialog v-model:open="showInvitationsDialog" />
    <CreateTaskDialog v-model:open="showCreateTaskDialog" :list-id="selectedListId" />
    <EditTaskDialog v-model:open="showEditTaskDialog" :task="editingTask" :list-id="selectedListId" />
    <DeleteTaskDialog v-model:open="showDeleteTaskDialog" :task="deletingTask" :list-id="selectedListId" />
    <MembersDialog v-model:open="showMembersDialog" :members="members" />
  </SidebarProvider>
</template>

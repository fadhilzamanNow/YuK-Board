<script setup lang="ts">
import { Plus } from "lucide-vue-next";
import { Button } from "@/components/ui/button";
import TaskCard from "./TaskCard.vue";

const props = defineProps<{
  listId: string;
  tasks: Task[];
}>();

const emit = defineEmits<{
  createTask: [];
  editTask: [task: Task];
  deleteTask: [task: Task];
}>();

const todoTasks = computed(() => props.tasks.filter((t) => t.status === "todo"));
const inProgressTasks = computed(() => props.tasks.filter((t) => t.status === "in_progress"));
const doneTasks = computed(() => props.tasks.filter((t) => t.status === "done"));
</script>

<template>
  <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
    <!-- Todo -->
    <div class="bg-muted/50 rounded-lg p-4">
      <div class="flex items-center justify-between mb-4">
        <h2 class="font-semibold">Todo</h2>
        <Button size="icon" variant="ghost" class="size-7" @click="emit('createTask')">
          <Plus class="size-4" />
        </Button>
      </div>
      <div class="space-y-2">
        <TaskCard
          v-for="task in todoTasks"
          :key="task.id"
          :task="task"
          :list-id="listId"
          @edit="emit('editTask', $event)"
          @delete="emit('deleteTask', $event)"
        />
        <p v-if="!todoTasks.length" class="text-muted-foreground text-sm">Belum ada tugas</p>
      </div>
    </div>

    <!-- In Progress -->
    <div class="bg-muted/50 rounded-lg p-4">
      <div class="flex items-center justify-between mb-4">
        <h2 class="font-semibold">In Progress</h2>
      </div>
      <div class="space-y-2">
        <TaskCard
          v-for="task in inProgressTasks"
          :key="task.id"
          :task="task"
          :list-id="listId"
          @edit="emit('editTask', $event)"
          @delete="emit('deleteTask', $event)"
        />
        <p v-if="!inProgressTasks.length" class="text-muted-foreground text-sm">Belum ada tugas</p>
      </div>
    </div>

    <!-- Done -->
    <div class="bg-muted/50 rounded-lg p-4">
      <div class="flex items-center justify-between mb-4">
        <h2 class="font-semibold">Done</h2>
      </div>
      <div class="space-y-2">
        <TaskCard
          v-for="task in doneTasks"
          :key="task.id"
          :task="task"
          :list-id="listId"
          @edit="emit('editTask', $event)"
          @delete="emit('deleteTask', $event)"
        />
        <p v-if="!doneTasks.length" class="text-muted-foreground text-sm">Belum ada tugas</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";
import { EllipsisVertical, Pencil, Trash2, Circle, Clock, CheckCircle } from "lucide-vue-next";
import { toast } from "vue-sonner";

const props = defineProps<{
  task: Task;
  listId: string;
}>();

const emit = defineEmits<{
  edit: [task: Task];
  delete: [task: Task];
}>();

const { mutateAsync: updateTask, isPending } = useMutationUpdateTask();

const creatorInitial = computed(() => props.task.creator?.name?.charAt(0).toUpperCase() ?? "?");

const handleStatusChange = (status: TaskStatus) => {
  if (status === props.task.status) return;
  updateTask(
    { listId: props.listId, taskId: props.task.id, params: { title: props.task.title, status } },
    {
      onError: (err: CustomErrorResponse) => toast.error(err.message || "Gagal mengubah status"),
    }
  );
};
</script>

<template>
  <Card class="p-3 hover:shadow-md transition-shadow">
    <div class="flex justify-between items-start gap-2">
      <div class="flex-1 min-w-0">
        <p class="font-medium text-sm truncate">{{ task.title }}</p>
        <p v-if="task.description" class="text-muted-foreground text-xs truncate mt-1">{{ task.description }}</p>
        <div class="flex items-center gap-1.5 mt-4">
          <Avatar class="size-5">
            <AvatarFallback class="text-[10px]">{{ creatorInitial }}</AvatarFallback>
          </Avatar>
          <span class="text-muted-foreground text-xs">{{ task.creator?.name }}</span>
        </div>
      </div>
      <div class="flex items-center gap-1">
        <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button size="icon" variant="ghost" class="size-7" :disabled="isPending">
              <Circle v-if="task.status === 'todo'" class="size-4" />
              <Clock v-else-if="task.status === 'in_progress'" class="size-4 text-yellow-500" />
              <CheckCircle v-else class="size-4 text-green-500" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem @click="handleStatusChange('todo')">
              <Circle class="size-4 mr-2" /> Todo
            </DropdownMenuItem>
            <DropdownMenuItem @click="handleStatusChange('in_progress')">
              <Clock class="size-4 mr-2 text-yellow-500" /> In Progress
            </DropdownMenuItem>
            <DropdownMenuItem @click="handleStatusChange('done')">
              <CheckCircle class="size-4 mr-2 text-green-500" /> Done
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
        <DropdownMenu>
          <DropdownMenuTrigger as-child>
            <Button size="icon" variant="ghost" class="size-7">
              <EllipsisVertical class="size-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuItem @click="emit('edit', task)">
              <Pencil class="size-4 mr-2" /> Edit
            </DropdownMenuItem>
            <DropdownMenuItem class="text-destructive" @click="emit('delete', task)">
              <Trash2 class="size-4 mr-2" /> Hapus
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
    </div>
  </Card>
</template>

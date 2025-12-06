type TaskStatus = "todo" | "in_progress" | "done";

interface Task {
  id: string;
  todo_list_id: string;
  title: string;
  description: string;
  status: TaskStatus;
  created_by: string;
  assigned_to?: string;
  created_at: string;
  updated_at: string;
  creator?: UserInfo;
  assigned_user?: UserInfo;
}

interface CreateTaskParams {
  title: string;
  description?: string;
  assigned_to?: string;
}

interface UpdateTaskParams {
  title: string;
  description?: string;
  status?: TaskStatus;
  assigned_to?: string;
}

interface TaskResponse {
  message: string;
  task: Task;
}

interface TasksResponse {
  message: string;
  tasks: Task[];
}

interface TodoList {
  id: string;
  title: string;
  description: string;
  owner_id: string;
  invite_code: string;
  created_at: string;
  updated_at: string;
  owner?: UserInfo;
  members?: ListMember[];
}

interface ListMember {
  id: string;
  todo_list_id: string;
  user_id: string;
  role: "owner" | "member";
  joined_at: string;
  user?: UserInfo;
}

interface CreateListParams {
  title: string;
  description?: string;
}

interface UpdateListParams {
  title: string;
  description?: string;
}

interface ListResponse {
  message: string;
  list: TodoList;
}

interface ListsResponse {
  message: string;
  lists: TodoList[];
}

interface Invitation {
  id: string;
  todo_list_id: string;
  inviter_id: string;
  invitee_id: string;
  status: "pending" | "accepted" | "declined" | "cancelled";
  token: string;
  created_at: string;
  responded_at?: string;
  todo_list?: TodoList;
  inviter?: UserInfo;
  invitee?: UserInfo;
}

interface InviteParams {
  email: string;
}

interface InvitationResponse {
  message: string;
  invitation: Invitation;
}

interface InvitationsResponse {
  message: string;
  invitations: Invitation[];
}

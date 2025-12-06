interface LoginParams {
  email: string;
  password: string;
}

interface RegisterParams extends LoginParams {
  name: string;
}

interface UserInfo {
  id: string;
  name: string;
  email: string;
  created_at: string;
  updated_at: string;
}

interface LoginResponse {
  message: string;
  token: string;
  user: UserInfo;
}

interface MeResponse {
  message: string;
  user: UserInfo;
}

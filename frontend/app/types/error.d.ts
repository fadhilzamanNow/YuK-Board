interface CustomErrorResponse<T> {
  message: string;
  errors: {
    [key: T]: string;
  };
}

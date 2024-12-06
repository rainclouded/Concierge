export interface ApiResponse<T> {
  message?: string;
  data: T;
	timestamp: string;
}

export interface ApiGenPasswordResponse {
    message?: string;
    password?: number;
    status?: string
}
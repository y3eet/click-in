import axios, {
  AxiosError,
  AxiosResponse,
  InternalAxiosRequestConfig,
} from "axios";
import { config } from "./config";

export const api = axios.create({
  baseURL: config.apiUrl,
  withCredentials: true,
});

interface QueueItem {
  resolve: (value?: unknown) => void;
  reject: (reason?: unknown) => void;
}

// Flag to prevent multiple refresh attempts
let isRefreshing = false;
let failedQueue: QueueItem[] = [];

const processQueue = (
  error: AxiosError | null,
  token: string | null = null
): void => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });

  failedQueue = [];
};

// Response interceptor to handle token refresh
api.interceptors.response.use(
  (response: AxiosResponse) => response,
  async (error: AxiosError) => {
    const originalRequest = error.config as InternalAxiosRequestConfig & {
      _retry?: boolean;
    };
    // Check token expiration
    const exp = localStorage.getItem("exp");
    if (!exp) return Promise.reject(error);
    const expDate = new Date(Number(exp) * 1000);
    const now = new Date();

    if (
      (expDate <= now || error.response?.status === 401) &&
      !originalRequest._retry
    ) {
      if (isRefreshing) {
        // If already refreshing, queue this request
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then(() => api(originalRequest))
          .catch((err) => Promise.reject(err));
      }

      originalRequest._retry = true;
      isRefreshing = true;

      try {
        // Call refresh token endpoint
        await api.post("/auth/refresh");
        console.log("Token refreshed successfully");
        isRefreshing = false;
        processQueue(null, null);

        // Retry the original request
        return api(originalRequest);
      } catch (refreshError) {
        console.log("Token refresh failed" + refreshError);
        processQueue(refreshError as AxiosError, null);
        isRefreshing = false;

        // const currentPath = window.location.pathname;
        // Redirect to login or handle refresh failure
        // if (currentPath !== "/auth/login") {
        //   window.location.href = "/auth/login";
        // }
        return Promise.reject(refreshError);
      }
    }

    return Promise.reject(error);
  }
);

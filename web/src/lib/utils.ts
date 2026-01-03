import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { ResponseError, Result } from "./types";
import axios from "axios";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export async function tryCatch<T, E = ResponseError>(
  promise: Promise<T>
): Promise<Result<T, E>> {
  try {
    const data = await promise;
    return { data, error: null };
  } catch (error) {
    return { data: null, error: error as E };
  }
}

export function parseErrorMessage(error: ResponseError): string {
  return error.response.data.error || "An unknown error occurred.";
}

export function buildFileUrl(key: string): string {
  return `${process.env.NEXT_PUBLIC_API_URL}/api/file/${key}`;
}

export function parseError(error: any) {
  if (axios.isAxiosError(error)) {
    return (error as ResponseError)?.response?.data.error;
  }
  return JSON.stringify(error).replace(/"/g, "");
}

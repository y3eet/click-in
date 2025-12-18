import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";
import { ResponseError, Result } from "./types";

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

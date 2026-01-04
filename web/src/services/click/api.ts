import { api } from "@/lib/axios";

export function createClick(clickableId: number) {
  return api.post<{ message: string }>("/api/click", {
    clickable_id: clickableId,
  });
}

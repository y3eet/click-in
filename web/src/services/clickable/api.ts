import { api } from "@/lib/axios";
import { CreateClickable, Clickable } from "./types";

export function getAllClickable() {
  return api.get<Clickable[]>("/api/clickable");
}

export function createClickable(data: CreateClickable) {
  return api.post<Clickable>("/api/clickable", data);
}

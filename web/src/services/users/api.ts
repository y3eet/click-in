import { api } from "@/lib/axios";
import { User } from "./types";

export function getAllUsers() {
  return api.get<User[]>("/api/users");
}

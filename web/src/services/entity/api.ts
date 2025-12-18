import { api } from "@/lib/axios";
import { CreateEntity, Entity } from "./types";

export function getAllEntities() {
  return api.get<Entity[]>("/api/entity");
}

export function createEntity(data: CreateEntity) {
  return api.post<Entity>("/api/entity", data);
}

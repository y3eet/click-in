import { User } from "../users/types";

export type Entity = {
  id: string;
  name: string;
  image_key: string;
  user_id: string;
  mp3_key: string;
  createdAt: Date;
  updatedAt: Date;
  user: User;
};

export type CreateEntity = Pick<Entity, "name" | "image_key" | "mp3_key">;

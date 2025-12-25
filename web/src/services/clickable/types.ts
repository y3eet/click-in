import { User } from "../users/types";

export type Clickable = {
  id: string;
  name: string;
  image_key: string;
  user_id: string;
  mp3_key: string;
  createdAt: Date;
  updatedAt: Date;
  user: User;
};

export type CreateClickable = Pick<Clickable, "name" | "image_key" | "mp3_key">;

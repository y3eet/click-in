import { Prettify } from "@/lib/types";

export type User = {
  id: number;
  provider_id: string;
  email: string;
  username: string;
  avatar_url: string;
  provider: string;
  created_at: string;
  updated_at: string;
};

export type UserPayload = User & {
  iss: string;
  exp: number;
  nbf: number;
  iat: number;
};

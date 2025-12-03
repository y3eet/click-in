import { User } from "../users/types";

export type ExchangeTokenResponse = {
  message: string;
  access_token: string;
  user: User;
};

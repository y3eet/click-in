import { User, UserPayload } from "../users/types";

export type ExchangeTokenResponse = {
  message: string;
  access_token: string;
  user: User;
};

export type AuthContextType = {
  currentUser?: UserPayload | null;
  isLoading: boolean;
  getCurrentUser: () => void;
  handleLogout: () => void;
};

import { api } from "@/lib/axios";
import { ExchangeTokenResponse } from "./types";
import { Prettify } from "@/lib/types";
import { UserPayload } from "../users/types";

export function exhangeToken(exchangeToken: string) {
  return api.post<ExchangeTokenResponse>("/auth/exchange", {
    exchange_token: exchangeToken,
  });
}

export function refreshToken() {
  return api.post("/auth/refresh");
}

export function currentUser() {
  return api.get<Prettify<UserPayload>>("/auth/current-user");
}

import { api } from "@/lib/axios";
import { ExchangeTokenResponse } from "./types";

export function exhangeToken(exchangeToken: string) {
  return api.post<ExchangeTokenResponse>("/auth/exchange", {
    exchange_token: exchangeToken,
  });
}

import { useMutation } from "@tanstack/react-query";
import { exhangeToken } from "./api";

export function useExchangeToken() {
  return useMutation({ mutationFn: exhangeToken });
}

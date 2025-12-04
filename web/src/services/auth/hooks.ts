import { useMutation, useQuery } from "@tanstack/react-query";
import { currentUser, exhangeToken, logout, refreshToken } from "./api";

export function useExchangeToken() {
  return useMutation({ mutationFn: exhangeToken });
}

export function useRefreshToken() {
  return useMutation({ mutationFn: refreshToken });
}

export function useCurrentUser() {
  return useQuery({
    queryKey: ["currentUser"],
    queryFn: currentUser,
  });
}

export function useLogout() {
  return useMutation({
    mutationFn: logout,
  });
}

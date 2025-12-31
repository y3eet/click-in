import { useQuery, useMutation } from "@tanstack/react-query";
import { createClickable, getAllClickable, getClickableById } from "./api";

export function useFetchClickable() {
  return useQuery({
    queryKey: ["clickable"],
    queryFn: getAllClickable,
  });
}

export function useCreateClickable() {
  return useMutation({
    mutationFn: createClickable,
  });
}

export function useFetchClickableById(id: string) {
  return useQuery({
    queryKey: ["clickable", id],
    queryFn: () => getClickableById(id),
  });
}

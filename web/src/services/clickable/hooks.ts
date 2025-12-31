import { useQuery, useMutation } from "@tanstack/react-query";
import { createClickable, getAllClickable } from "./api";

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

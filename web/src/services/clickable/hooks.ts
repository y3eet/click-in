import { useQuery, useMutation } from "@tanstack/react-query";
import { createClickable, getAllEntities } from "./api";

export function useFetchEntities() {
  return useQuery({
    queryKey: ["entities"],
    queryFn: getAllEntities,
  });
}

export function useCreateClickable() {
  return useMutation({
    mutationFn: createClickable,
  });
}

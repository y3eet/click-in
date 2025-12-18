import { useQuery, useMutation } from "@tanstack/react-query";
import { createEntity, getAllEntities } from "./api";

export function useFetchEntities() {
  return useQuery({
    queryKey: ["entities"],
    queryFn: getAllEntities,
  });
}

export function useCreateEntity() {
  return useMutation({
    mutationFn: createEntity,
  });
}

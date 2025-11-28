import { useQuery } from "@tanstack/react-query";
import { getAllUsers } from "./api";

export function useGetUsers() {
  return useQuery({
    queryKey: ["users"],
    queryFn: getAllUsers,
  });
}

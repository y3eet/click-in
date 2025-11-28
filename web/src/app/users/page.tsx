"use client";

import { useGetUsers } from "@/services/users/hooks";

export default function Page() {
  const { data, error } = useGetUsers();
  return (
    <div>
      <div>{JSON.stringify({ data, error }, null, 2)}</div>
    </div>
  );
}

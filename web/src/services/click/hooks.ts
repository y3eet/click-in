import { useMutation } from "@tanstack/react-query";
import { createClick } from "./api";
import { useState } from "react";
import { useSSE } from "@/hooks/use-sse";

export function useCreateClick() {
  return useMutation({
    mutationFn: createClick,
  });
}

export function useStreamClickCount(clickable_id: number) {
  const [clickCount, setClickCount] = useState<number | null>(null);
  const [error, setError] = useState<string | null>(null);

  useSSE<{ count: number }>(
    `/api/click/event/count/${clickable_id}`,
    (data) => {
      setClickCount(data.count);
    },
    () => {
      setError("Error receiving click count updates.");
    }
  );

  return { clickCount, error };
}

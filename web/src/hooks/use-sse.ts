import { useEffect } from "react";

type SSEResponse<T> = T & {
  data: T;
  timestamp: string;
};

export function useSSE<T>(
  url: string,
  onMessage: (data: SSEResponse<T>) => void,
  onError: (error: Event) => void
) {
  useEffect(() => {
    const resolvedUrl = process.env.NEXT_PUBLIC_API_URL + url;
    const eventSource = new EventSource(resolvedUrl, { withCredentials: true });

    eventSource.onmessage = (event) => {
      try {
        const data: SSEResponse<T> = JSON.parse(event.data);
        onMessage(data);
      } catch (error: unknown) {
        console.log("Error parsing SSE message:", error);
        onMessage(event.data);
      }
    };

    eventSource.onerror = (error: Event) => {
      onError(error);
      eventSource.close();
    };

    return () => {
      eventSource.close();
    };
  }, [url, onMessage, onError]);
}

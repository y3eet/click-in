import { useEffect, useRef } from "react";

export function useWebSocket<T extends object>(
  url: string,
  onMessage: (data: T) => void,
  onOpen?: () => void,
  onClose?: () => void
) {
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket(url);
    socketRef.current = socket;

    socket.onopen = () => {
      console.log("WebSocket connection established");
      if (onOpen) {
        onOpen();
      }
    };

    socket.onmessage = (event) => {
      console.log("Message from server ", event.data);
      const data: T = JSON.parse(event.data);
      onMessage(data);
    };

    socket.onclose = () => {
      console.log("WebSocket connection closed");
    };

    return () => {
      socket.close();
      socketRef.current = null;
      if (onClose) {
        onClose();
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [url]);

  function sendMessage(message: T) {
    if (socketRef.current && socketRef.current.readyState === WebSocket.OPEN) {
      socketRef.current.send(JSON.stringify(message));
    }
  }

  return { sendMessage };
}

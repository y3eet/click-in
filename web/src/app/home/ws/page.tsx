"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useWebSocket } from "@/hooks/use-websocket";
import { useRef } from "react";

export default function WebSocketPage() {
  const { sendMessage } = useWebSocket<{ data: string }>(
    "ws://localhost:8080/ws",
    (data) => {
      console.log("Received data:", data);
    },
    () => {
      console.log("WebSocket opened");
    },
    () => {
      console.log("WebSocket closed");
    }
  );

  const inputRef = useRef<HTMLInputElement>(null);

  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold mb-4">WebSocket Test Page</h1>
      <div className="flex space-x-2">
        <Input ref={inputRef} placeholder="Type a message..." />
        <Button
          onClick={() => {
            if (inputRef.current) {
              sendMessage({ data: inputRef.current.value });
              inputRef.current.value = "";
            }
          }}
        >
          Send
        </Button>
      </div>
    </div>
  );
}

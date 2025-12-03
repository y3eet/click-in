"use client";

import { tryCatch } from "@/lib/utils";
import { useExchangeToken } from "@/services/auth/hooks";
import { useSearchParams } from "next/navigation";
import { useEffect } from "react";

export default function Page() {
  const searchParams = useSearchParams();
  const exchangeToken = searchParams.get("exchange_token");

  const { mutateAsync, isPending } = useExchangeToken();
  async function handleExchangeToken() {
    if (!exchangeToken) return;
    const { data, error } = await tryCatch(mutateAsync(exchangeToken));
    console.log({ data, error });
  }

  useEffect(() => {
    handleExchangeToken();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  return <div>{isPending && <>Lading.....</>}</div>;
}

"use client";

import {
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty";
import { Spinner } from "@/components/ui/spinner";
import { tryCatch } from "@/lib/utils";
import { useExchangeToken } from "@/services/auth/hooks";
import { useAuthContext } from "@/services/auth/provider";
import { useRouter, useSearchParams } from "next/navigation";
import { Suspense, useEffect } from "react";

export default function Page() {
  return (
    <Suspense>
      <AuthLoader />
    </Suspense>
  );
}

function AuthLoader() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const exchangeToken = searchParams.get("exchange_token");

  const { mutateAsync, isPending } = useExchangeToken();
  const { getCurrentUser } = useAuthContext();

  async function handleExchangeToken() {
    if (!exchangeToken) {
      router.push("/home");
      return;
    }
    const { error } = await tryCatch(mutateAsync(exchangeToken));
    if (!error) {
      getCurrentUser();
      router.push("/home");
    }
  }

  useEffect(() => {
    handleExchangeToken();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <Empty className="w-full">
      <EmptyHeader>
        <EmptyMedia variant="icon">{isPending && <Spinner />}</EmptyMedia>
        <EmptyTitle>Authenticating</EmptyTitle>
        <EmptyDescription>
          Please wait while we securely authenticate your account...
        </EmptyDescription>
      </EmptyHeader>
    </Empty>
  );
}

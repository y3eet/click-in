"use client";

import { Clickable } from "@/components/clickable";
import { useParams } from "next/navigation";
import { useFetchClickableById } from "@/services/clickable/hooks";
import { buildImageUrl } from "@/lib/utils";
import Loader from "@/components/loader";

export default function ClickablePage() {
  const params = useParams<{ id: string }>();

  const { data, isLoading, error } = useFetchClickableById(params.id);
  return (
    <div>
      <Loader loading={isLoading} error={error}>
        {data && (
          <Clickable
            imageUrl={buildImageUrl(data.data.image_key!)}
            alt={`Clickable ${params.id}`}
          />
        )}
      </Loader>
    </div>
  );
}

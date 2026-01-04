"use client";

import { Clickable } from "@/components/clickable";
import { useParams } from "next/navigation";
import { useFetchClickableById } from "@/services/clickable/hooks";
import { buildFileUrl } from "@/lib/utils";
import Loader from "@/components/loader";
import { useCreateClick, useStreamClickCount } from "@/services/click/hooks";

export default function ClickablePage() {
  const params = useParams<{ id: string }>();
  const { data, isLoading, error } = useFetchClickableById(params.id);
  const { clickCount, error: streamError } = useStreamClickCount(
    Number(params.id)
  );
  const { mutate } = useCreateClick();

  function handleClickableClick() {
    mutate(Number(params.id));
  }

  return (
    <div className="w-full">
      <Loader loading={isLoading} error={error}>
        <div className="flex items-center justify-center px-4 py-8">
          <div className="w-full max-w-2xl space-y-8">
            <div className="text-center space-y-2">
              <h1 className="text-3xl font-bold tracking-tight">
                {data?.data.name}
              </h1>
              <p className="text-muted-foreground">
                {clickCount ?? "Loading..."} Clicks
              </p>
            </div>

            {data && (
              <div className="flex justify-center items-center">
                <Clickable
                  imageUrl={buildFileUrl(data.data.image_key!)}
                  alt={`Clickable ${params.id}`}
                  onClick={() => {
                    handleClickableClick();
                    const audio = document.getElementById(
                      `clickable-audio-${params.id}`
                    ) as HTMLAudioElement | null;

                    if (!audio) return;
                    audio.pause();
                    audio.currentTime = 0;
                    audio.play();
                  }}
                />
              </div>
            )}

            {data?.data?.mp3_key ? (
              <div className="bg-card rounded-lg border shadow-sm p-6 space-y-4">
                <h2 className="text-lg font-semibold tracking-tight">
                  Audio Controls
                </h2>
                <audio
                  id={`clickable-audio-${params.id}`}
                  src={buildFileUrl(data.data.mp3_key)}
                  preload="auto"
                  controls
                  className="w-full"
                />
              </div>
            ) : (
              <div>
                <p className="text-center text-sm text-muted-foreground">
                  No audio available for this clickable.
                </p>
              </div>
            )}
          </div>
        </div>
      </Loader>
    </div>
  );
}

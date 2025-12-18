"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Progress } from "@/components/ui/progress";
import { Spinner } from "@/components/ui/spinner";
import { useFileUpload } from "@/hooks/use-file-upload";
import { tryCatch } from "@/lib/utils";
import { useCreateEntity } from "@/services/entity/hooks";
import Image from "next/image";
import { useState } from "react";

export default function Page() {
  const [name, setName] = useState("");
  const [imageKey, setImageKey] = useState("");
  const [audioKey, setAudioKey] = useState("");

  const { isPending, mutateAsync } = useCreateEntity();

  async function handleCreate() {
    const { error } = await tryCatch(
      mutateAsync(
        {
          name,
          image_key: imageKey,
          mp3_key: audioKey,
        },
        {
          onSuccess: () => {
            setName("");
            setImageKey("");
            setAudioKey("");
          },
        }
      )
    );
    if (error) {
      console.error("Error creating entity:", error);
    }
  }

  return (
    <div className="pt-10">
      <h1 className="text-3xl font-bold">Create Entity</h1>
      <div className="mt-4 grid w-full max-w-sm items-center gap-2">
        <Label htmlFor="name">Name</Label>
        <Input
          value={name}
          onChange={(e) => setName(e.target.value)}
          type="text"
          id="name"
          placeholder="Enter entity name"
          className="mt-1 w-full max-w-sm"
        />
      </div>
      <UploadImage onImageUpload={(e) => setImageKey(e)} />
      <Mp3Upload onMp3Upload={(e) => setAudioKey(e)} />
      <Button className="mt-6" onClick={handleCreate} disabled={isPending}>
        {isPending && <Spinner className="mr-2" />}
        Create
      </Button>
    </div>
  );
}
type UploadImageProps = {
  onImageUpload: (key: string) => void;
};

function UploadImage(props: UploadImageProps) {
  const { uploadFile, progress, fileUrl, error, uploading } = useFileUpload();

  async function handleFileUpload(file: File) {
    const key = await uploadFile(file);
    props.onImageUpload(key);
  }

  return (
    <div className="mt-4 grid w-full max-w-sm items-center gap-2">
      <Label htmlFor="picture">Image</Label>
      <div className="flex items-center gap-2">
        <Input
          placeholder="Upload an image"
          type="file"
          id="picture"
          accept="image/*"
          onChange={(e) => {
            const file = e.target.files?.[0];
            if (file) {
              handleFileUpload(file);
            }
          }}
        />
      </div>
      {uploading && <Progress value={progress} className="w-full" />}
      {fileUrl && (
        <Image
          unoptimized
          src={fileUrl}
          alt="Uploaded"
          className="mt-2 max-h-48 w-auto rounded-md border"
          width={192}
          height={192}
        />
      )}
      {error && <p className="text-sm text-red-600">{error}</p>}
    </div>
  );
}

type Mp3UploadProps = {
  onMp3Upload: (key: string) => void;
};

function Mp3Upload(props: Mp3UploadProps) {
  const { uploadFile, progress, fileUrl, error, uploading } = useFileUpload();

  async function handleFileUpload(file: File) {
    const key = await uploadFile(file);
    props.onMp3Upload(key);
  }
  return (
    <div className="mt-4 grid w-full max-w-sm items-center gap-2">
      <Label htmlFor="audio">Audio</Label>
      <div className="flex items-center gap-2">
        <Input
          placeholder="Upload an audio file"
          type="file"
          id="audio"
          accept="audio/mpeg"
          onChange={(e) => {
            const file = e.target.files?.[0];
            if (file) {
              handleFileUpload(file);
            }
          }}
        />
      </div>
      {uploading && <Progress value={progress} className="w-full" />}
      {fileUrl && (
        <audio className="mt-2 w-full" controls>
          <source src={fileUrl} type="audio/mpeg" />
          Your browser does not support the audio element.
        </audio>
      )}
      {error && <p className="text-sm text-red-600">{error}</p>}
    </div>
  );
}

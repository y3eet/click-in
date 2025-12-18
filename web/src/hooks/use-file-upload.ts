import { api } from "@/lib/axios";
import { parseErrorMessage, tryCatch } from "@/lib/utils";
import { useState } from "react";

export function useFileUpload() {
  const [progress, setProgress] = useState(0);
  const [uploading, setUploading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [fileUrl, setFileUrl] = useState<string | null>(null);

  const uploadFile = async (file: File): Promise<string> => {
    const formData = new FormData();
    formData.append("file", file);

    const { error, data } = await tryCatch(
      api.post("/api/file/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
        onUploadProgress: (p) => {
          setUploading(true);
          setError(null);
          const total = p.total || 1;
          const percentCompleted = Math.round((p.loaded * 100) / total);
          setProgress(percentCompleted);
        },
      })
    );

    if (error) {
      setUploading(false);
      setProgress(0);
      setError(parseErrorMessage(error));
      return "";
    }

    setUploading(false);
    const fileUrl = `${process.env.NEXT_PUBLIC_API_URL}/api/file/${
      data?.data.file || ""
    }`;

    setFileUrl(fileUrl);
    return fileUrl;
  };

  return {
    uploading,
    error,
    fileUrl,
    uploadFile,
    progress,
  };
}

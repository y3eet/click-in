"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { buildImageUrl } from "@/lib/utils";
import { useFetchEntities } from "@/services/entity/hooks";
import Link from "next/link";
import Image from "next/image";

export default function BrowsePage() {
  const { data } = useFetchEntities();
  return (
    <div className="py-10">
      <h1 className="text-3xl font-bold">Browse Page</h1>
      <div className="flex gap-4 items-center mt-4 w-full">
        <Input placeholder="Search..." className="w-full" />
        <Link href="/home/browse/create">
          <Button>Create Entity</Button>
        </Link>
      </div>
      <div className="mt-6 grid gap-4">
        {data?.data.map((entity) => (
          <div
            key={entity.id}
            className="p-4 border rounded-md hover:shadow-md transition"
          >
            <h2 className="text-xl font-semibold">{entity.name}</h2>
            {entity.image_key && (
              <Image
                unoptimized
                src={buildImageUrl(entity.image_key)}
                alt={entity.name}
                className="mt-2 w-32 h-32 object-cover"
                width={128}
                height={128}
              />
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

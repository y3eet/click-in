"use client";

import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { buildFileUrl } from "@/lib/utils";
import { useFetchClickable } from "@/services/clickable/hooks";
import Link from "next/link";
import Image from "next/image";
import { useRouter } from "next/navigation";

export default function BrowsePage() {
  const { data } = useFetchClickable();
  const router = useRouter();

  return (
    <div className="py-10">
      <h1 className="text-3xl font-bold">Browse Page</h1>
      <div className="flex gap-4 items-center mt-4 w-full">
        <Input placeholder="Search..." className="w-full" />
        <Link href="/home/browse/create">
          <Button>Create Clickable</Button>
        </Link>
      </div>
      <div className="mt-6 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5 gap-6">
        {data?.data.map((clickable) => (
          <div
            onClick={() => router.push(`clickable/${clickable.id}`)}
            key={clickable.id}
            className="p-4 border rounded-md hover:shadow-md transition"
          >
            <h2 className="text-xl font-semibold">{clickable.name}</h2>
            {clickable.image_key && (
              <Image
                unoptimized
                src={buildFileUrl(clickable.image_key)}
                alt={clickable.name}
                width={350}
                height={350}
              />
            )}
          </div>
        ))}
      </div>
    </div>
  );
}

import { Input } from "@/components/ui/input";

export default function BrowsePage() {
  return (
    <div className="py-10">
      <h1 className="text-3xl font-bold">Browse Page</h1>
      <div>
        <Input placeholder="Search..." className="mt-4 w-full max-w-md" />
      </div>
    </div>
  );
}

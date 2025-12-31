import { parseError } from "@/lib/utils";
import { Spinner } from "./ui/spinner";

export default function Loader({
  loading,
  error,
  children,
}: {
  loading: boolean;
  error?: Error | null;
  children?: React.ReactNode;
}) {
  if (loading) {
    return (
      <div className="flex justify-center items-center h-full">
        <Spinner />
      </div>
    );
  }
  if (error) {
    return <div>Error: {parseError(error)}</div>;
  }
  return <>{children}</>;
}

import { createContext, useContext, useEffect } from "react";
import { AuthContextType } from "./types";
import { useCurrentUser, useLogout } from "./hooks";
import { useRouter } from "next/navigation";

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const { data, isLoading, refetch } = useCurrentUser();
  const { mutate: logout } = useLogout();

  useEffect(() => {
    const user = data?.data;
    if (user) {
      const exp = user.exp;
      localStorage.setItem("exp", String(exp));
    }
    console.log(user);
  }, [data]);

  function getCurrentUser() {
    refetch();
  }

  function handleLogout() {
    logout(undefined, {
      onSuccess: () => {
        localStorage.removeItem("exp");
        router.push("/auth/login");
        refetch();
      },
    });
  }

  const value = {
    currentUser: data?.data,
    isLoading,
    getCurrentUser,
    handleLogout,
  };
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const useAuthContext = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuthContext must be used within an AuthProvider");
  }
  return context;
};

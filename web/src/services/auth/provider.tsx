import { createContext, useContext, useEffect } from "react";
import { AuthContextType } from "./types";
import { useCurrentUser } from "./hooks";

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const { data, isLoading, refetch } = useCurrentUser();

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

  const value = {
    currentUser: data?.data,
    isLoading,
    getCurrentUser,
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

import { createContext, useContext, useState } from "react";
import { ApiClient } from "@/api/apiClient";
import { TokenStorage } from "@/utils/tokenStorage";

interface AuthContextType {
  isAuthenticated: boolean;
  signIn: (username: string, password: string) => Promise<void>;
  signOut: () => void;
  signUp: (username: string, email: string, password: string) => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [isAuthenticated, setIsAuthenticated] = useState<boolean>(
    !!TokenStorage.getToken(),
  );

  const signIn = async (username: string, password: string) => {
    const response = await ApiClient.getInstance().request<{ token: string }>(
      "/sign/in",
      {
        method: "POST",
        body: JSON.stringify({ username, password }),
        headers: { "Content-Type": "application/json" },
      },
    );

    if (response.data?.token) {
      TokenStorage.setToken(response.data.token);
      setIsAuthenticated(true);
    }
  };

  const signUp = async (username: string, email: string, password: string) => {
    await ApiClient.getInstance().request("/sign/up", {
      method: "POST",
      body: JSON.stringify({ username, email, password }),
      headers: { "Content-Type": "application/json" },
    });
  };

  const signOut = () => {
    TokenStorage.clearToken();
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, signIn, signOut, signUp }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within an AuthProvider");
  return context;
};

/*
    useEffect(() => {
     
      const validateToken = async () => {
        const token = TokenStorage.getToken();
        if (token) {
          try {
            await ApiClient.getInstance().request('/auth/validate');
            setIsAuthenticated(true);
          } catch (error) {
            TokenStorage.clearToken();
            setIsAuthenticated(false);
          }
        }
      };
      validateToken();
    }, []); */

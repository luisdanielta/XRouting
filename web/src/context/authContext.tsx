import React, { createContext, useContext, useState } from "react";
import { AuthService } from "@/ports/authServices";

interface AuthContextType {
  isAuthenticated: boolean;
  role: "user" | "admin" | null;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [isAuthenticated, setIsAuthenticated] = useState(
    AuthService.prototype.isAuthenticated(),
  );
  const [role, setRole] = useState<"user" | "admin" | null>(
    AuthService.prototype.getUserRole(),
  );

  const login = async (email: string, password: string) => {
    await AuthService.prototype.signIn(email, password);
    setIsAuthenticated(true);
    setRole(AuthService.prototype.getUserRole());
  };

  const logout = () => {
    AuthService.prototype.signOut();
    setIsAuthenticated(false);
    setRole(null);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, role, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};

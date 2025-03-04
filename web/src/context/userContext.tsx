import { User } from "@/domain/entities/user";
import { createContext, useState } from "react";

// Define the structure of the UserContext
interface UserContextType {
  user: User | null; // Stores the authenticated user or null if not logged in
  isAuthenticated: boolean; // Indicates if the user is logged in
  signIn: (userData: User) => void; // Logs in the user
  signOut: () => void; // Logs out the user
}

// Create the UserContext
const UserContext = createContext<UserContextType | undefined>(undefined);

// Provider component for UserContext
export const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [user, setUser] = useState<User | null>(null);

  // Determine if the user is authenticated based on the user state
  const isAuthenticated = !!user;

  // Function to log in the user
  const signIn = (userData: User) => {
    setUser(userData);
  };

  // Function to log out the user
  const signOut = () => {
    setUser(null);
  };

  return (
    <UserContext.Provider value={{ user, isAuthenticated, signIn, signOut }}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext;

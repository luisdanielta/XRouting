import { useContext } from "react";
import UserContext from "@/context/userContext";

export const UseUserContext = () => {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserProvider");
  }
  return context;
};

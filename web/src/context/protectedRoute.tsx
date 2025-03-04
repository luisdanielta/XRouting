/*
import React from "react";
import { Navigate, Outlet } from "react-router";
import { UseUserContext } from "@/hooks/useUserContext";

interface ProtectedRouteProps {
  allowedRoles: Array<"user" | "admin">;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({
  allowedRoles,
}) => {

  const { user, isAuthenticated } = useUserContext();
  if (!isAuthenticated) {
    return <Navigate to="/sign/in" replace />;
  }

  if (user?.role && !allowedRoles.includes(user.role)) {
    return <Navigate to="/sign/in" replace />;
  }
  

  return <Outlet />;
};
*/

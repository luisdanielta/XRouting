import React from "react";
import { Navigate, Outlet } from "react-router";

interface ProtectedRouteProps {
  allowedRoles: Array<"user" | "moderator">;
}

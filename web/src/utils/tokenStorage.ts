export interface ITokenStorage {
  setToken(token: string): void;
  getToken(): string | null;
  clearToken(): void;
  setRole(role: "user" | "admin"): void;
  getRole(): "user" | "admin" | null;
  clearRole(): void;
}

export const TokenStorage: ITokenStorage = {
  setToken: (token: string): void => {
    localStorage.setItem("jwt_token", token);
  },

  getToken: (): string | null => {
    return localStorage.getItem("jwt_token");
  },

  clearToken: (): void => {
    localStorage.removeItem("jwt_token");
  },

  setRole: (role: "user" | "admin"): void => {
    localStorage.setItem("user_role", role);
  },

  getRole: (): "user" | "admin" | null => {
    return localStorage.getItem("user_role") as "user" | "admin" | null;
  },

  clearRole: (): void => {
    localStorage.removeItem("user_role");
  },
};

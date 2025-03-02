import { jwtDecode } from "jwt-decode";
import { ApiClient } from "../apiClient";
import { ITokenStorage, TokenStorage } from "@/utils/tokenStorage";
import { User } from "@/domain/entities/User";
import { Endpoints as E } from "../endpoints";

/**
 * Represents the decoded JWT payload.
 */
interface DecodedToken {
  id: string;
  name: string;
  username: string;
  email: string;
  role: "user" | "moderator";
  exp: number; // Expiration timestamp in seconds
}

export interface IAuthService {
  login(
    username: string,
    password: string,
  ): Promise<{ user: User; token: string }>;
  signup(
    username: string,
    email: string,
    password: string,
  ): Promise<{ message: string }>;
  logout(): Promise<void>;
  isAuthenticated(): boolean;
  getRole(): "user" | "moderator" | null;
  getUserFromToken(): User | null;
}

/**
 * Authentication service that handles login, signup, and token-based authentication.
 */
export class AuthService implements IAuthService {
  private apiClient: ApiClient;
  private tokenStorage: ITokenStorage;

  /**
   * Constructor with dependency injection for API client and token storage.
   * @param apiClient - Instance of ApiClient for API calls.
   * @param tokenStorage - Instance implementing ITokenStorage for token handling.
   */
  constructor(apiClient: ApiClient, tokenStorage: ITokenStorage) {
    this.apiClient = apiClient;
    this.tokenStorage = tokenStorage;
  }

  /**
   * Logs in the user by making an API request and decoding the token.
   * @returns A promise resolving to the user object and token.
   */
  public async login(
    username: string,
    password: string,
  ): Promise<{ user: User; token: string }> {
    try {
      const response = await this.apiClient.request<{ token: string }>(
        E.auth.login,
        {
          method: "POST",
          body: { username, password },
        },
      );

      const token = response.data.token;
      const decoded = jwtDecode<DecodedToken>(token);
      const user = new User(
        decoded.name,
        decoded.username,
        decoded.email,
        true,
        decoded.role,
      );

      this.tokenStorage.setToken(token);

      return { user, token };
    } catch (error) {
      console.error("Error during login:", error);
      throw new Error("Failed to log in. Please try again.");
    }
  }

  /**
   * Registers a new user via API.
   * @returns A promise with a success message.
   */
  public async signup(
    username: string,
    email: string,
    password: string,
  ): Promise<{ message: string }> {
    try {
      await this.apiClient.request(E.auth.signup, {
        method: "POST",
        body: { username, email, password },
      });
      return { message: "User created successfully" };
    } catch (error) {
      console.error("Error during signup:", error);
      throw new Error("Failed to sign up. Please try again.");
    }
  }

  /**
   * Logs out the user, clearing token storage and notifying the API.
   */
  public async logout(): Promise<void> {
    try {
      this.tokenStorage.clearToken();
      await this.apiClient.request(E.auth.logout, { method: "POST" });
    } catch (error) {
      console.error("Error during logout:", error);
      throw new Error("Failed to log out. Please try again.");
    }
  }

  /**
   * Checks if the user is authenticated by verifying the stored token.
   * @returns True if authenticated, false otherwise.
   */
  public isAuthenticated(): boolean {
    const token = this.tokenStorage.getToken();
    if (!token) return false;

    try {
      const decoded = jwtDecode<DecodedToken>(token);
      const isExpired = decoded.exp * 1000 < Date.now();
      if (isExpired) {
        this.tokenStorage.clearToken();
        return false;
      }
      return true;
    } catch (error) {
      console.error("Error decoding token:", error);
      return false;
    }
  }

  /**
   * Retrieves the user's role from the stored token.
   * @returns The user role or null if unavailable.
   */
  public getRole(): "user" | "moderator" | null {
    const token = this.tokenStorage.getToken();
    if (!token) return null;

    try {
      const decoded = jwtDecode<DecodedToken>(token);
      return decoded.role;
    } catch (error) {
      console.error("Error decoding token:", error);
      return null;
    }
  }

  /**
   * Retrieves the User instance from the stored token.
   * @returns The user instance or null if the token is invalid.
   */
  public getUserFromToken(): User | null {
    const token = this.tokenStorage.getToken();
    if (!token) return null;

    try {
      const decoded = jwtDecode<DecodedToken>(token);
      return new User(
        decoded.name,
        decoded.username,
        decoded.email,
        true,
        decoded.role,
      );
    } catch (error) {
      console.error("Error decoding token:", error);
      return null;
    }
  }
}

export const authServiceInstance = new AuthService(
  ApiClient.getInstance(),
  TokenStorage,
);

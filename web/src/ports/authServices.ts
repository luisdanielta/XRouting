import { TokenStorage } from "@/utils/tokenStorage";
import { ApiClient } from "@/api/apiClient";

interface AuthResponse {
    token: string;
    role: "user" | "admin";
}

export class AuthService {
    private apiClient: ApiClient;

    constructor() {
        this.apiClient = ApiClient.getInstance();
    }

    /**
     * Logs in a user by sending credentials to the backend.
     * @param email User email.
     * @param password User password.
     * @returns A promise that resolves when the login is successful.
     */
    public async signIn(email: string, password: string): Promise<void> {
        const response = await this.apiClient.request<AuthResponse>("/sign/in", {
            method: "POST",
            body: JSON.stringify({ email, password }),
            headers: { "Content-Type": "application/json" },
        });

        if (response.data) {
            TokenStorage.setToken(response.data.token);
            TokenStorage.setRole(response.data.role);
        }
    }

    /**
     * Logs out the user by clearing stored tokens and roles.
     */
    public signOut(): void {
        TokenStorage.clearToken();
        TokenStorage.clearRole();
    }

    /**
     * Checks if the user is authenticated.
     * @returns True if a token is stored, false otherwise.
     */
    public isAuthenticated(): boolean {
        return !!TokenStorage.getToken();
    }

    /**
     * Gets the current user's role.
     * @returns The role of the authenticated user or null if not set.
     */
    public getUserRole(): "user" | "admin" | null {
        return TokenStorage.getRole();
    }
}
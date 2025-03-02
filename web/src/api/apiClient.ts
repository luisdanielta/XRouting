import { TokenStorage } from "@/utils/tokenStorage"
import { HttpClient, HttpRequestConfig, HttpResponse } from "./httpClient"

/**
 * ApiClient is a singleton class that extends HttpClient to provide
 * a centralized API client for making HTTP requests to the backend.
 * It includes token-based authentication by adding an Authorization
 * header to each request if a token is available.
 */
export class ApiClient extends HttpClient {
    /**
     * The singleton instance of the ApiClient.
     */
    private static instance: ApiClient

    /**
     * Private constructor to prevent direct instantiation.
     * Initializes the HttpClient with the base URL for the API.
     */
    private constructor() {
        super("http://localhost:9050/v1/api")
    }

    /**
     * Returns the singleton instance of the ApiClient.
     * If the instance does not exist, it creates one.
     *
     * @returns {ApiClient} The singleton instance of the ApiClient.
     */
    public static getInstance(): ApiClient {
        if (!ApiClient.instance) {
            ApiClient.instance = new ApiClient()
        }
        return ApiClient.instance
    }

    /**
     * Makes an HTTP request to the specified endpoint with the given configuration.
     * Adds an Authorization header with a token if available.
     *
     * @template T The type of the response data.
     * @param {string} endpoint The API endpoint to request.
     * @param {HttpRequestConfig} [config={}] The configuration for the HTTP request.
     * @returns {Promise<HttpResponse<T>>} A promise that resolves to the HTTP response.
     */
    public async request<T>(
        endpoint: string,
        config: HttpRequestConfig = {},
    ): Promise<HttpResponse<T>> {
        const token = TokenStorage.getToken()
        const headers = {
            ...config.headers,
            ...(token ? { Authorization: `Bearer ${token}` } : {}),
        }

        return super.request<T>(endpoint, { ...config, headers })
    }
}
export interface HttpRequestConfig {
    method?: "GET" | "POST" | "PUT" | "DELETE"
    headers?: Record<string, string>
    body?: unknown
}

export interface HttpResponse<T = unknown> {
    data: T
    status: number
}

export abstract class HttpClient {
    private readonly baseUrl: string

    constructor(baseUrl: string) {
        this.baseUrl = baseUrl
    }

    protected async request<T>(
        endpoint: string,
        { method = "GET", headers = {}, body }: HttpRequestConfig = {},
    ): Promise<HttpResponse<T>> {
        const url = `${this.baseUrl}${endpoint}`
        const config: RequestInit = {
            method,
            headers: {
                "Content-Type": "application/json",
                ...headers,
            },
            body: body ? JSON.stringify(body) : undefined,
        }

        const response = await fetch(url, config)

        if (!response.ok) {
            const errorData = await response.json()
            throw new Error(errorData.message || "Error en la solicitud")
        }

        const data = await response.json()
        return { data, status: response.status }
    }
}
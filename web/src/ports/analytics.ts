import { ApiClient } from "@/api/apiClient";
import { Analytic, AnalyticSchema } from "@/domain/entities/analytic";

export class AnalyticService {
  private apiClient = ApiClient.getInstance();
  private baseEndpoint = "/analytic";

  async create(analytic: Analytic): Promise<Analytic> {
    const validatedAnalytic = AnalyticSchema.parse(analytic);
    return this.apiClient
      .request<Analytic>(this.baseEndpoint, {
        method: "POST",
        body: JSON.stringify(validatedAnalytic),
        headers: { "Content-Type": "application/json" },
      })
      .then((res) => res.data);
  }

  async findById(id: string): Promise<Analytic | null> {
    return this.apiClient
      .request<Analytic>(`${this.baseEndpoint}/${id}`, {
        method: "GET",
      })
      .then((res) => res.data)
      .catch(() => null);
  }

  async update(
    id: string,
    updateData: Partial<Analytic>,
  ): Promise<Analytic | null> {
    return this.apiClient
      .request<Analytic>(`${this.baseEndpoint}/${id}`, {
        method: "PUT",
        body: JSON.stringify(updateData),
        headers: { "Content-Type": "application/json" },
      })
      .then((res) => res.data)
      .catch(() => null);
  }

  async delete(id: string): Promise<boolean> {
    return this.apiClient
      .request<void>(`${this.baseEndpoint}/${id}`, {
        method: "DELETE",
      })
      .then(() => true)
      .catch(() => false);
  }

  async findAll(): Promise<Analytic[]> {
    return this.apiClient
      .request<Analytic[]>(this.baseEndpoint + "s", {
        method: "GET",
      })
      .then((res) => res.data);
  }
}

import { ApiClient } from "@/api/apiClient";
import { Maintenance, MaintenanceSchema } from "@/domain/entities/maintenance";

export class MaintenanceService {
  private apiClient = ApiClient.getInstance();
  private baseEndpoint = "/maintenance";

  async create(maintenance: Maintenance): Promise<Maintenance> {
    const validatedMaintenance = MaintenanceSchema.parse(maintenance);
    return this.apiClient
      .request<Maintenance>(this.baseEndpoint, {
        method: "POST",
        body: JSON.stringify(validatedMaintenance),
        headers: { "Content-Type": "application/json" },
      })
      .then((res) => res.data);
  }

  async findById(id: string): Promise<Maintenance | null> {
    return this.apiClient
      .request<Maintenance>(`${this.baseEndpoint}/${id}`, {
        method: "GET",
      })
      .then((res) => res.data)
      .catch(() => null);
  }

  async update(
    id: string,
    updateData: Partial<Maintenance>,
  ): Promise<Maintenance | null> {
    return this.apiClient
      .request<Maintenance>(`${this.baseEndpoint}/${id}`, {
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

  async findAll(): Promise<Maintenance[]> {
    return this.apiClient
      .request<Maintenance[]>(this.baseEndpoint + "s", {
        method: "GET",
      })
      .then((res) => res.data);
  }
}
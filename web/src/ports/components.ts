import { ApiClient } from "@/api/apiClient";
import { Component, ComponentSchema } from "@/domain/entities/component";

export class ComponentService {
  private apiClient = ApiClient.getInstance();
  private baseEndpoint = "/component";

  async create(component: Component): Promise<Component> {
    const validatedComponent = ComponentSchema.parse(component);
    return this.apiClient
      .request<Component>(this.baseEndpoint, {
        method: "POST",
        body: JSON.stringify(validatedComponent),
        headers: { "Content-Type": "application/json" },
      })
      .then((res) => res.data);
  }

  async findById(id: string): Promise<Component | null> {
    return this.apiClient
      .request<Component>(`${this.baseEndpoint}/${id}`, {
        method: "GET",
      })
      .then((res) => res.data)
      .catch(() => null);
  }

  async update(
    id: string,
    updateData: Partial<Component>,
  ): Promise<Component | null> {
    return this.apiClient
      .request<Component>(`${this.baseEndpoint}/${id}`, {
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

  async findAll(): Promise<Component[]> {
    return this.apiClient
      .request<Component[]>(this.baseEndpoint + "s", {
        method: "GET",
      })
      .then((res) => res.data);
  }
}

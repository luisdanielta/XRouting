import { Maintenance, MaintenanceSchema } from "../entities/maintenance";

export class MaintenanceRepository {
  private maintenances: Map<string, Maintenance> = new Map();

  async create(maintenance: Maintenance): Promise<Maintenance> {
    const validatedMaintenance = MaintenanceSchema.parse(maintenance);
    this.maintenances.set(validatedMaintenance.id, validatedMaintenance);
    return validatedMaintenance;
  }

  async findById(id: string): Promise<Maintenance | null> {
    return this.maintenances.get(id) || null;
  }

  async update(
    id: string,
    updateData: Partial<Maintenance>,
  ): Promise<Maintenance | null> {
    if (!this.maintenances.has(id)) return null;

    const existingMaintenance = this.maintenances.get(id)!;
    const updatedMaintenance = { ...existingMaintenance, ...updateData };
    const validatedMaintenance = MaintenanceSchema.parse(updatedMaintenance);

    this.maintenances.set(id, validatedMaintenance);
    return validatedMaintenance;
  }

  async delete(id: string): Promise<boolean> {
    return this.maintenances.delete(id);
  }

  async findAll(): Promise<Maintenance[]> {
    return Array.from(this.maintenances.values());
  }
}

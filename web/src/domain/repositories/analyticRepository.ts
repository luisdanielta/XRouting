import { Analytic, AnalyticSchema } from "../entities/analytic";

export class AnalyticRepository {
  private analytics: Map<string, Analytic> = new Map();

  async create(analytic: Analytic): Promise<Analytic> {
    const validatedAnalytic = AnalyticSchema.parse(analytic);
    this.analytics.set(validatedAnalytic.id, validatedAnalytic);
    return validatedAnalytic;
  }

  async findById(id: string): Promise<Analytic | null> {
    return this.analytics.get(id) || null;
  }

  async update(
    id: string,
    updateData: Partial<Analytic>,
  ): Promise<Analytic | null> {
    if (!this.analytics.has(id)) return null;

    const existingAnalytic = this.analytics.get(id)!;
    const updatedAnalytic = { ...existingAnalytic, ...updateData };
    const validatedAnalytic = AnalyticSchema.parse(updatedAnalytic);

    this.analytics.set(id, validatedAnalytic);
    return validatedAnalytic;
  }

  async delete(id: string): Promise<boolean> {
    return this.analytics.delete(id);
  }

  async findAll(): Promise<Analytic[]> {
    return Array.from(this.analytics.values());
  }
}

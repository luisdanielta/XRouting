import { Component, ComponentSchema } from "../entities/component";

export class ComponentRepository {
  private components: Map<string, Component> = new Map();

  async create(component: Component): Promise<Component> {
    const validatedComponent = ComponentSchema.parse(component);
    this.components.set(validatedComponent.id, validatedComponent);
    return validatedComponent;
  }

  async findById(id: string): Promise<Component | null> {
    return this.components.get(id) || null;
  }

  async update(id: string, updateData: Partial<Component>): Promise<Component | null> {
    if (!this.components.has(id)) return null;
    
    const existingComponent = this.components.get(id)!;
    const updatedComponent = { ...existingComponent, ...updateData };
    const validatedComponent = ComponentSchema.parse(updatedComponent);
    
    this.components.set(id, validatedComponent);
    return validatedComponent;
  }

  async delete(id: string): Promise<boolean> {
    return this.components.delete(id);
  }

  async findAll(): Promise<Component[]> {
    return Array.from(this.components.values());
  }
}

import { z } from "zod";

export const ComponentType = z.enum([
  "core",
  "capsule",
  "engine",
  "ship",
  "launchpad",
  "rocket",
]);

export type ComponentType = z.infer<typeof ComponentType>;

export const ComponentStatus = z.enum([
  "active",
  "inactive",
  "damaged",
  "destroyed",
]);
export type ComponentStatus = z.infer<typeof ComponentStatus>;

export const ComponentCategory = z.enum(["fixed", "launched"]);
export type ComponentCategory = z.infer<typeof ComponentCategory>;

export const ComponentSchema = z.object({
  id: z.string(),
  componentId: z.string().describe("The ID of the component"),
  type: ComponentType.describe("The type of the component"),
  status: ComponentStatus.describe("The status of the component"),
  category: ComponentCategory.describe("The category of the component"),
  missions: z
    .array(z.string())
    .default([])
    .describe("The missions this component has been a part of"),
  hasSubcomponents: z
    .boolean()
    .default(false)
    .describe("Whether this component has subcomponents"),
});

export type Component = z.infer<typeof ComponentSchema>;

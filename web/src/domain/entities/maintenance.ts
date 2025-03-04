import { z } from "zod";

export const FailureRiskLevel = z.enum(["low", "medium", "high"]);
export type FailureRiskLevel = z.infer<typeof FailureRiskLevel>;

export const FailurePredictionSchema = z.object({
  riskLevel: FailureRiskLevel.describe(
    "Risk level of failure (low, medium, high)",
  ),
  estimatedFailureDate: z.coerce
    .date()
    .nullable()
    .describe("Predicted failure date if available"),
  reason: z.string().describe("Reason for failure prediction"),
});
export type FailurePrediction = z.infer<typeof FailurePredictionSchema>;

export const MaintenanceSchema = z.object({
  id: z.string(),
  maintenanceId: z.string().describe("The unique ID of the maintenance"),
  componentId: z.string().describe("The ID of the component affected"),
  subcomponentsAffected: z
    .array(z.string())
    .default([])
    .describe("List of affected subcomponent IDs"),
  date: z.coerce.date().describe("Date of maintenance"),
  performedBy: z.string().describe("User ID who performed the maintenance"),
  cost: z.number().describe("Cost of the maintenance"),
  description: z
    .string()
    .nullable()
    .describe("Description of the maintenance performed"),
  nextMaintenance: z.coerce
    .date()
    .nullable()
    .describe("Scheduled next maintenance date"),
  failurePrediction: FailurePredictionSchema.nullable().describe(
    "Prediction of potential future failures",
  ),
});

export type Maintenance = z.infer<typeof MaintenanceSchema>;

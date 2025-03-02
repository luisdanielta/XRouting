import { z } from "zod"

export const logInSchema = z.object({
  username: z
    .string()
    .min(2, {
      message: "El nombre de usuario debe tener al menos 2 caracteres.",
    })
    .max(50, {
      message: "El nombre de usuario no puede exceder los 50 caracteres.",
    }),
  password: z
    .string()
    .min(8, { message: "La contraseña debe tener al menos 8 caracteres." }),
})

export const signUpSchema = z.object({
  email: z
    .string()
    .email({ message: "Por favor, ingresa un correo electrónico válido." }),
  username: z
    .string()
    .min(2, {
      message: "El nombre de usuario debe tener al menos 2 caracteres.",
    })
    .max(50, {
      message: "El nombre de usuario no puede exceder los 50 caracteres.",
    }),
  password: z
    .string()
    .min(8, { message: "La contraseña debe tener al menos 8 caracteres." }),
})
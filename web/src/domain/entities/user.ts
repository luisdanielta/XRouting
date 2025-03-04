export class User {
  constructor(
    public readonly id: string | null,
    public readonly username: string,
    public readonly role: "user" | "admin" = "user", // Restricting roles
  ) {}
}

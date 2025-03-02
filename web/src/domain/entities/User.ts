export class User {
  constructor(
    // public readonly id: string | null,
    public readonly name: string,
    public readonly username: string,
    public readonly email: string,
    public readonly status: boolean = true,
    public readonly role: "user" | "moderator" = "user", // Restricting roles
  ) { }

  public getDisplayName(): string {
    return `${this.name} (${this.username})`
  }
}
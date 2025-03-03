from pydantic import BaseModel, Field
from enum import Enum

class UserRole(str, Enum):
    admin = "admin"
    engineer = "engineer"
    viewer = "viewer"

class User(BaseModel):
    id: int
    name: str = Field(..., description="Name of the user")
    email: str = Field(..., description="Email of the user")
    password: str = Field(..., description="Password of the user")
    role: UserRole = Field(..., description="Role of the user")
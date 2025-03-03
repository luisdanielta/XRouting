from pydantic import BaseModel, Field

class Comment(BaseModel):
    id: str
    componentId: str = Field(..., description="Component ID of the comment")
    userId: str = Field(..., description="User ID of the comment")
    message: str = Field(..., description="Message of the comment")
    updatedDate: str = Field(..., description="Updated date of the comment")


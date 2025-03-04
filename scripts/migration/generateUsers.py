import uuid
from core.entities.user import User, UserRole
from core.repositories.userRepository import UserRepository

def generateUsers():
    return [
        User(id=str(uuid.uuid4()), name="Admin SpaceX", email="admin@example.com", password="securepass", role=UserRole.admin),
        User(id=str(uuid.uuid4()), name="Jhon Doe", email="engineer1@example.com", password="engineerpass", role=UserRole.engineer),
        User(id=str(uuid.uuid4()), name="Lena Lopez", email="engineer2@example.com", password="engineerpass", role=UserRole.engineer),
        User(id=str(uuid.uuid4()), name="Alex Morgan", email="engineer3@example.com", password="engineerpass", role=UserRole.engineer),
        User(id=str(uuid.uuid4()), name="Viewer SpaceX", email="viewer@example.com", password="viewerpass", role=UserRole.viewer),
    ]

def insertUsers():
    userRepo = UserRepository()
    users = generateUsers()
    try:
        userRepo.saveBatch(users) 
        print("Users inserted successfully.")
    except Exception as e:
        print(f"Failed to insert users: {e}")

if __name__ == "__main__":
    insertUsers()

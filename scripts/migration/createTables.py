from adapters.repositories.componentRepository import ComponentRepository
from adapters.repositories.subComponentRepository import SubComponentRepository
from adapters.repositories.commentsRepository import CommentRepository
from adapters.repositories.maintenanceRepository import MaintenanceRepository
from adapters.repositories.analyticRepository import AnalyticRepository
from adapters.repositories.userRepository import UserRepository


def createTables():
    repositories = [
        ComponentRepository(),
        SubComponentRepository(),
        CommentRepository(),
        MaintenanceRepository(),
        AnalyticRepository(),
        UserRepository()
    ]

    for repo in repositories:
        try:
            print(f"Creating table: {repo.tableName}...")
            repo.createTableIfNotExists()
            print(f" Table '{repo.tableName}' created successfully.")
        except Exception as e:
            print(f"Error creating table '{repo.tableName}': {e}")

    print("All tables processed.")

if __name__ == "__main__":
    createTables()

# CRUD application

## Goals

This project a command-line interface (CLI) TODO application that allows users to manage their tasks. The application support CRUD operations (Create, Read, Update, Delete) for both users and tasks entities.

## Application assumptions

1. User Management
    - The application allow users to create, read, update, and delete user profiles.
    - Users are able to provide details such as their name, email, phone
2. Task Management
    - The application enables users to create, read, update, and delete tasks.
    - Tasks have properties such as a title, description, due date, status.
3. User-Task Relationship
    - Each task is associated with a specific user.
    - Users are able to view their own tasks as well as tasks assigned to other users.
4. Command-Line Interface
    - The CLI provides a clear and intuitive syntax for entering commands and interacting with the application.
    - It displays instructions, usage guidelines, and error messages when necessary.
5. Data Storage
    - The application persists user and task data to a file.
    - Users and tasks are stored separately and linked appropriately.
6. Error Handling
    - The application handles errors gracefully and provide informative error messages to guide the user.
    - It validates user inputs, ensuring they meet the required format and constraints.
7. User-Friendly Outputs
    - The application displays task lists and user profiles in a user-friendly and readable format.
    - Task lists provides relevant details and clearly indicate task status and due dates.

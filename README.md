# Task Management System

A simple task management system using the actor model in Go.

## Overview

This project implements a basic task management system where tasks and users are represented as actors. The actor model is employed to simulate interactions between tasks and users through message passing.

## Project Structure

The project follows the following structure:

```plaintext
task-manager/
|-- main.go
|-- model/
|   |-- actor.go
|   |-- message.go
|   |-- task.go
|   |-- user.go
|-- manager/
|   |-- task_manager.go
|   |-- user_manager.go
```

## Features

- **User Registration:** Register users with a unique username and role.
- **Task Management:** Create tasks with a status, process ID (PID), and an associated message.
- **Task Assignment:** Assign tasks from one user to another.


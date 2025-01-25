# Contribute Guidelines

1. Clone the repository to your local machine.
2. Create a branch (`/feature`, `/bugfix`, or `/hotfix`) based on your needs.
3. Make the necessary changes.
4. Push your changes to the Git repository.
5. Create a pull request.
6. Add appropriate labels to the pull request.

You can also use the provided `docker-compose` file for local testing. More information about local environment setup can be found [here](/doc/local-env.md/).

**Note!** Don't forget to validate your code via make validate (for GO only) or make validate_codebase (for whole CodeBase)

# Go - Practices and Coding Conventions

## Introduction

This document outlines the coding conventions and practices for the **Go** project, specifically tailored for GoLang development using Gin Gonic. Adhering to these standards ensures code consistency, readability, and maintainability across the project.

## Table of Contents

1. [General Coding Standards](#general-coding-standards)
   - [Naming Conventions](#naming-conventions)
   - - [Variables and Functions](#variables-and-functions)
   - - [Structs](#structs)
   - - [Interfaces](#interfaces)
   - - [Constants](#constants)
   - [Code Structure](#code-structure)
   - - [Packages](#packages)
   - - [File Naming](#file-naming)
   - - [Folder Structure](#folder-structure)
   - - [Test Files](#test-files)
   - [Formatting](#formatting)
   - - [Newlines](#newlines)
   - - [Commenting](#commenting)
   - [Error Handling](#error-handling)
   - - [Error return](#error-return)
   - - [Error Messages](#error-messages)
   - - [Error shorthand](#error-shorthand)
   - [Miscellaneous](#miscellaneous)
   - - [One-letter variable names](#one-letter-variable-names)
   - - [Enumerations](#enumerations)
2. [Layer-Specific Guidelines](#layer-specific-guidelines)
   - [Router (using Gin Gonic)](#router-using-gin-gonic)
   - [Controller](#controller)
   - [Service](#service)
   - [Repository](#repository)
3. [Best Practices for PRs](#best-practices-for-prs)
4. [Commit Messages](#commit-messages)
5. [Conclusion](#conclusion)

---

## General Coding Standards

### Naming Conventions

#### **Variables and Functions**:

Use `camelCase` for variable.

- Additional words can be added to disambiguate similar names, for example userCount and projectCount.
- Do not simply drop letters to save typing. For example `sandbox` is preferred over `Sbx`, particularly for exported names.
- Omit types and type-like words from most variable names.
- - For a number, `userCount` is a better name than `numUsers` or `usersInt`.
- - For a slice, `users` is a better name than `userSlice`.
- Omit words that are clear from the surrounding context. For example, in the implementation of a `UserCount` method, a local variable called `userCount` is probably redundant; `count`, `users` are just as readable.
- Use `camelCase` for function names if they are not exported, and `PascalCase` if they are exported. For example, `getUserByID` is a private function, and `GetUserByID` is a public function. All that is not exported should be added after the // private comment.

#### **Structs**:

Use `PascalCase` for struct names.

**Good Example**:

```go
type User struct {
    ID   int
    Name string
}

users := []User{
    {ID: 1, Name: "Alice"},
    {ID: 2, Name: "Bob"},
}

var userCount int

func GetUserByID(userID int) {
    // Implementation
}

// private
func getUserByID(userID int) {
    // Implementation
}
```

**Bad Example**:

```go
type UserStruct struct {
    UserID   int
    UserName string
}

userSlice := []UserStruct{
    {UserID: 1, UserName: "Alice"},
    {UserID: 2, UserName: "Bob"},
}

var user_count int

func getUserByID(userID int) {
    // Implementation
}

// private
func GetUserByID(userID int) {
    // Implementation
}
```

---

#### **Interfaces**:

- Interfaces should be prefixed with `I` (e.g., `IService`) to indicate it's an interface.

Good Example:

```go
type IUserService interface {
    GetUserByID(id int) (*User, error)
}
```

Bad Example:

```go
type UserService interface {
    GetUserByID(id int) (*User, error)
}
```

---

#### **Constants**:

Use MixedCaps for constants.

- **MixedCaps**: For example, a constant is MaxLength (not MAX_LENGTH) if exported and maxLength (not max_length) if unexported.

**Good Example:**

```go
const MaxLength = 10
const minPasswordLength = 8
```

**Bad Example:**

```go
const MAX_LENGTH = 10
const MIN_PASSWORD_LENGTH = 8
```

---

### Code Structure

#### **Packages**:

Go package names should be short and contain only lowercase letters. A package name composed of multiple words should be left unbroken in all lowercase. For example, the package tabwriter is not named `tabWriter`, `TabWriter`, or `tab_writer`

#### **File Naming**:

Use descriptive names, and separate words with underscores (e.g., `user.go`, `user_validations.go`).

#### **Folder Structure**:

Organize code into meaningful packages and folders (e.g., `controller/`, `service/`, `repository/`, `model/`)

#### **Test Files**:

Name test files with `_test.go` suffix (e.g., `user_test.go`).

**Good Example:**

```go
package user
package uservalidations
```

```
app/
    ├── controller/
    │   ├── user.go
    |   ├── user_test.go
    │   └── ...
    ├── service/
    │   ├── user.go
    |   ├── user_test.go
    │   └── ...
    ├── repository/
    │   ├── user.go
    |   ├── user_test.go
    │   └── ...
    |── model/
    |   ├── user.go
    └── main.go
```

**Bad Example:**

```go
package User
package user_validations
```

```
app/
    ├── controllers/
    │   ├── controller_user.go
    |   ├── controller_userTest.go
    │   └── ...
    ├── services/
    │   ├── userService.go
    |   ├── userService_test.go
    │   └── ...
    ├── repositories/
    │   ├── userRepository.go
    |   ├── userRepositorytest.go
    │   └── ...
    |── models/
    |   ├── user_model.go
    └── main.go
```

---

### Formatting

#### **Newlines**:

Ensure there is a newline after `}` when there is a `return` or new `var` and before `var` or any other line following a closing brace. Also, include a newline after each `case` in a `switch` statement.

**Good Example:**

```go
func example() {
    if condition {
        return
    }

    var x int
    x = 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }

    switch x {
    case 1:
        fmt.Println("x is 1")

    case 2:
        fmt.Println("x is 2")

    default:
        fmt.Println("x is neither 1 nor 2")
    }

    return x
}
```

**Bad Example:**

```go
func example() {
    if condition {
        return
    }
    var x int
    x = 10
    if x > 5 {
        fmt.Println("x is greater than 5")
    }
    switch x {
    case 1:
        fmt.Println("x is 1")
    case 2:
        fmt.Println("x is 2")
    default:
        fmt.Println("x is neither 1 nor 2")
    }
    return x
}
```

#### **Commenting**:

Add comments to explain complex logic or non-obvious code.

---

### Error Handling

#### **Error return**:

Prefer returning errors explicitly instead of using panic.

**Good Example:**

```go
func getUserByID(userID int) (*User, error) {
    user, err := userRepository.FindByID(userID)
    if err != nil {
        return nil, fmt.Errorf("finding user by ID failed: %v", err)
    }

    return user, nil
}
```

**Bad Example:**

```go
func getUserByID(userID int) *User {
    user, err := userRepository.FindByID(userID)
    if err != nil {
        panic(err)
    }

    return user
}
```

---

#### **Error Messages**:

Provide meaningful error messages when returning errors.

**Good Example:**

```go
func getUserByID(userID int) (*User, error) {
    user, err := userRepository.FindByID(userID)
    if err != nil {
        return nil, fmt.Errorf("finding user by ID failed: %v", err)
    }

    return user, nil
}
```

**Bad Example:**

```go
func getUserByID(userID int) (*User, error) {
    user, err := userRepository.FindByID(userID)
    if err != nil {
        return nil, err
    }

    return user, nil
}
```

---

#### **Error shorthand**:

Use the `err` shorthand for error variables.

**Good Example:**

```go
if err := someFunction(); err != nil {
    return err
}
```

**Bad Example:**

```go
err := someFunction()
if err != nil {
    return err
}
```

---

### Miscellaneous

#### One-letter variable names:

Avoid using one-letter variable names except in cases like loop indices (`i`, `j`, `k`).

**Good Example:**

```go
func calculateArea(length, width int) int {
    return length * width
}
```

**Bad Example:**

```go
func calculateArea(l, w int) int {
    return l * w
}
```

#### Enumerations:

Enums should be defined within the `model` package with their respective constants.

**Good Example:**

```go
package enum

type Coverage string

const (
	CoverageMostSegments      Coverage = "MOST_SEGMENTS"
	CoverageAtLeastOneSegment Coverage = "AT_LEAST_ONE_SEGMENT"
	CoverageAllSegments       Coverage = "ALL_SEGMENTS"
)

var Coverages = []string{
	string(CoverageMostSegments),
	string(CoverageAtLeastOneSegment),
	string(CoverageAllSegments),
}
```

**Bad Example:**

```go
package enum

const (
    MostSegments      = "MOST_SEGMENTS"
    AtLeastOneSegment = "AT_LEAST_ONE_SEGMENT"
    AllSegments       = "ALL_SEGMENTS"
)

var Coverages = []string{
    MostSegments,
    AtLeastOneSegment,
    AllSegments,
}
```

---

## Layer-Specific Guidelines

The software architecture that we are using is "Layered Architecture" or "Layered Design Pattern." In this architecture, the different components of the system are organized into separate layers, each with its own specific responsibility.

### Router (using Gin Gonic)

- Define routes using Gin Gonic's Router.
- Use meaningful route paths and HTTP methods.

#### Example: Gin Router Setup

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    // Example route with handler
    r.GET("/users/:id", getUserHandler)

    r.Run(":8080")
}

func getUserHandler(c *gin.Context) {
    // Implementation to fetch user by ID
    userID := c.Param("id")

    // Call service layer to fetch user details
    user, err := userService.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}
```

### Controller

- Controllers should handle HTTP request/response lifecycle.
- Validate request, query, path parameters, and call service methods.

Example: **UserController**

```go
package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

// UserController handles user-related requests
type UserController struct {
    userService IUserService
}

// NewUserController creates a new UserController
func NewUserController(us IUserService) *UserController {
    return &UserController{userService: us}
}

// GetUserByID retrieves a user by ID
func (uc *UserController) GetUserByID(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
        return
    }

    user, err := uc.userService.GetUser(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}
```

### Service

- Business logic should reside here.
- Follow Single Responsibility Principle (SRP) for service methods.

Example: **UserService**

```go
package services

import (
    "errors"
)

// UserService provides user-related operations
type UserService struct {
    // Service dependencies if any
}

// NewUserService creates a new instance of UserService
func NewUserService() *UserService {
    return &UserService{}
}

// GetUser retrieves a user by ID
func (us *UserService) GetUser(id int) (*User, error) {
    // Example logic to fetch user from repository layer
    user, err := userRepository.FindByID(id)
    if err != nil {
        return nil, errors.New("failed to fetch user")
    }

    return user, nil
}
```

### Repository

- Data access layer for interacting with databases or external services.
- Implement data persistence logic.

Example: **UserRepository**

```go
package repositories

// UserRepository provides data access operations for users
type UserRepository struct {
    // Repository dependencies if any
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

// FindByID retrieves a user by ID from the database
func (ur *UserRepository) FindByID(id int) (*User, error) {
    // Example implementation to fetch user from database or external service
    // Placeholder code for demonstration purposes
    user := &User{
        ID:   id,
        Name: "John Doe",
        Age:  30,
    }

    return user, nil
}
```

## Best Practices for PRs

- Consistency: Ensure adherence to coding standards and guidelines.
- - Review your code before creating a PR to catch any errors or issues.
- Code Reviews: All PRs must be reviewed by at least one team member.
- Formatting: Check for proper indentation, brace placement, and comment usage.

## Commit Messages

[Conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) are used for commit messages.

Commit messages should be clear and descriptive. They should follow the format `type: subject` where type is one of the following:

- `feat`: a new feature
- `fix`: a bug fix
- `docs`: changes to documentation
- `style`: formatting, missing semi colons, etc; no code change
- `refactor`: refactoring production code
- `test`: adding tests, refactoring test; no production code change

## Conclusion

By following these coding conventions and best practices, we aim to maintain a high standard of code quality, readability, and maintainability across the **Go** project. Consistency in coding style and structure will help streamline development and collaboration among team members. For any questions or suggestions regarding these practices, please reach out to the team.

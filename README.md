# EnbekTap1.0

This project implements a job search platform similar to HeadHunter, built using Golang for the backend and MongoDB for the database.  It provides functionality for job seekers to search and apply for jobs, and for companies to post and manage job listings.

## Features

* **Job Searching:**  Users can search for jobs based on keywords, location, salary, and other criteria.
* **Job Posting:** Companies can create and manage job postings, including details like job description, requirements, and salary.
* **User Profiles:** Job seekers can create profiles, upload resumes, and track their job applications.
* **Company Profiles:** Companies can create profiles to showcase their information and branding.
* **Application Management:**  Users can track their job applications, and companies can manage received applications.
* **Authentication and Authorization:** Secure user authentication and authorization to protect sensitive data.

## Technologies Used

* **Golang:**  Backend API development.
* **MongoDB:**  NoSQL database for storing job listings, user profiles, and other data.
* **Gin (Optional):**  Web framework for Golang (recommended for routing and middleware).
* **mgo/mongo-driver:** MongoDB driver for Golang.  (mgo is older but sometimes simpler; the official mongo-driver is recommended for new projects).
* **Bcrypt (or similar):**  For password hashing.
* **JWT (Optional):**  For stateless authentication.


## Getting Started

1. **Prerequisites:** Make sure you have Go and MongoDB installed on your system.

2. **Clone the repository:**

```bash
git clone https://github.com/your-username/job-search-platform.git
```

Run the application:
```bash
cd backend
go run main.go
```
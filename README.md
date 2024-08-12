# PASSWORD SERVICE

### User Story for Password Manager

#### Title: Password Manager for Secure Storage of User Credentials

As a user,  
I want to securely store and manage my passwords for various websites,  
so that I can easily access them without needing to remember each one individually.

---

Acceptance Criteria:

1. User Registration and Login:
   - Given a new user,  
   When they visit the password manager site,  
   Then they should be able to register by providing a username, email, and a secure password.
   - Given an existing user,  
   When they enter their credentials,  
   Then they should be able to log in to their account.

2. Password Storage:
   - Given a logged-in user,  
   When they enter a website URL, username/email, and password,  
   Then the password should be securely encrypted before being saved to the database.
   - Given the sensitivity of passwords,  
   When a user stores a password,  
   Then it should be stored in a way that ensures it cannot be accessed by an unauthorized individuals. (Another user should not be able to see a different user's stored password)

3. Password Retrieval:
   - Given a logged-in user,  
   When they request to view their saved passwords,  
   Then they should be able to decrypt and view them in a secure manner.

4. Security and Permissions:
   - Given the importance of security,  
   When a user attempts to access saved passwords,  
   Then they must be logged in and authorized to view their own data.
   - Given that passwords are sensitive information,  
   When a user logs out,  
   Then all access to stored passwords should be revoked immediately.


5. Additional Security Measures:
   - Given the risk of unauthorized access,  
   When a user registers or logs in,  
   Then their password should follow best practices for security (e.g., minimum length, complexity requirements).
   - Given the potential threats,  
   When storing passwords,  
   Then they should be hashed using a strong encryption algorithm and never stored in plain text.

---

This project will help in understanding and implementing crucial backend topics such as user authentication and authorization, database management, encryption, and overall security best practices specifically in golang.
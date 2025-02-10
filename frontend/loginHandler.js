const ngrokURL = "https://d249-212-96-87-84.ngrok-free.app"
const localhost = "http://localhost:8080"

document.addEventListener('DOMContentLoaded', () => {
    const loginForm = document.getElementById("login-container");
    const signupForm = document.getElementById("signup-container");
    const authForm = document.getElementById("auth-container");
    const loginTag = document.getElementById("log-in-tag");
    const signupTag = document.getElementById("sign-up-tag");
    const loginBttn = document.getElementById("log-in-bttn");
    const signupBttn = document.getElementById("sign-up-bttn");
    const authFormBttn = document.getElementById("auth-code-bttn");
    const logoutbttn = document.getElementById("logout")

    let signupEmail, passwordFirst, passwordSecond, name, role;

    logoutbttn.addEventListener("click", async (event) => {
        event.preventDefault();  // Prevent default form submission if applicable
    
        try {
            const token = localStorage.getItem("access_token"); // Make sure token is available
    
            const response = await fetch(`${localhost}/auth/logout`, {  // ✅ Correct endpoint
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${token}`  // ✅ Send token if required by backend
                },
                credentials: 'include' // ✅ Required if using cookies
            });
    
            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || "Logout failed");
            }
    
            // ✅ Only parse JSON if there's content
            let data = {};
            if (response.headers.get("content-type")?.includes("application/json")) {
                data = await response.json();
            }
    
            // ✅ Clear token from localStorage
            localStorage.removeItem("access_token");
            localStorage.removeItem("email");
            localStorage.removeItem("role")
    
            // ✅ Redirect to login page
            window.location.href = "./login.html";
    
        } catch (error) {
            console.error("Logout error:", error);
            alert(error.message || "Logout failed. Please try again.");
        }
    });

    signupTag.addEventListener("click", () => {
        loginForm.style.display = "none";
        signupForm.style.display = "block";
        document.getElementById("loginForm").reset();
    });

    loginTag.addEventListener("click", () => {
        signupForm.style.display = "none";
        loginForm.style.display = "block";
        document.getElementById("signupForm").reset();
    });

    loginBttn.addEventListener("click", async (event) => {
        event.preventDefault();
        const loginEmail = document.getElementById("loginEmail").value.trim();
        const loginPassword = document.getElementById("loginPassword").value.trim();

        try {
            const response = await fetch(`${localhost}/api/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                credentials: 'include',
                body: JSON.stringify({
                    Email: loginEmail,
                    Password: loginPassword
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.error || "Login failed");
            }

            const data = await response.json();
            
            localStorage.setItem("access_token", data.jwt);

            window.location.href = "./main.html";

        } catch (error) {
            console.error("Login error:", error);
            alert(error.message || "Login failed. Please try again.");
        }
    });

    signupBttn.addEventListener("click", (event) => {
        signupEmail = document.getElementById("signupEmail").value.trim();
        passwordFirst = document.getElementById("passwordFirst").value.trim();
        passwordSecond = document.getElementById("passwordSecond").value.trim();
        name = document.getElementById("fullName").value.trim();
        role = document.getElementById("role").value.trim();

        if (!signupEmail || !passwordFirst || !passwordSecond) {
            alert("Please fill in all fields.");
            event.preventDefault();
            return;
        }

        if (passwordFirst !== passwordSecond) {
            alert("Passwords do not match.");
            event.preventDefault();
            return;
        }

        const emailPattern = /^[a-zA-Z0-9._%+-]+@gmail\.com$/;
        if (!emailPattern.test(signupEmail)) {
            alert("Please enter a valid email address ending with @gmail.com.");
            event.preventDefault();
            return;
        }

        fetch(`${localhost}/api/send-code`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "ngrok-skip-browser-warning": "true"
            },
            body: JSON.stringify({ email: signupEmail })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error("Code sending failed");
            }
            return response.json();
        })
        .then(data => {
            console.log("Code sent successfully:", data);
        })
        .catch(error => {
            console.error("Error during code sending:", error);
            alert("An error occurred. Please try again.");
        });

        signupForm.style.display = "none";
        authForm.style.display = "block";
    });

    authFormBttn.addEventListener("click", async (event) => {
        event.preventDefault();
        const authCode = document.getElementById("authCode").value.trim();

        if (!authCode) {
            alert("Please enter verification code");
            return;
        }

        try {
            const response = await fetch(`${localhost}/api/signup`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "ngrok-skip-browser-warning": "true",
                },
                body: JSON.stringify({
                    FullName: name,
                    Email: signupEmail,
                    Password: passwordFirst,
                    Role: role,
                    VerificationCode: authCode
                })
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message || "Registration failed");
            }

            const data = await response.json();
            console.log("Registration successful:", data);
            alert("Registration successful! Please login.");
            window.location.href = "./login.html";

        } catch (error) {
            console.error("Registration error:", error);
            alert(error.message || "Registration failed. Please try again.");
        }
    });

    function updateNavbarForLoggedInUser() {
        const loginLink = document.getElementById('login-link');
        loginLink.textContent = 'Logout';
        loginLink.href = '#';
        loginLink.addEventListener('click', () => {
            // Implement logout functionality here
            alert('Logged out successfully');
            window.location.href = 'login.html';
        });

        const profileLink = document.createElement('a');
        profileLink.href = 'profile.html';
        profileLink.className = 'profile-link';
        profileLink.textContent = 'Profile';
        document.querySelector('.nav-right').appendChild(profileLink);
    }
});
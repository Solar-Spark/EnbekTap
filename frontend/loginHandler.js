const ngrokURL = "https://d249-212-96-87-84.ngrok-free.app"

document.addEventListener('DOMContentLoaded', () => {

    const loginForm = document.getElementById("login-container");
    const signupForm = document.getElementById("signup-container");
    const authForm = document.getElementById("auth-container")
    const loginTag = document.getElementById("log-in-tag");//login anchor tag
    const signupTag = document.getElementById("sign-up-tag");//signup anchor tag

    const loginBttn = document.getElementById("log-in-bttn");
    const signupBttn = document.getElementById("sign-up-bttn");

    let signupEmail, passwordFirst, passwordSecond, name;

    const authFormBttn = document.getElementById("auth-code-bttn")


    //when signup link clicked
    signupTag.addEventListener("click", () => {
        loginForm.style.display = "none";//hide login form
        signupForm.style.display = "block";//display signup form
        document.getElementById("loginForm").reset();//reset input values of login form
    });

    //when login link clicked
    loginTag.addEventListener("click", () => {
        signupForm.style.display = "none";//hide signup form
        loginForm.style.display = "block";//display login form
        document.getElementById("signupForm").reset();//reset input values of signup form

    });

    //send email and login info to the server
    loginBttn.addEventListener("click", (event) => {
        const loginEmail = document.getElementById("loginEmail").value.trim();
        const loginPassword = document.getElementById("loginPassword").value.trim();

        // Check if the required fields are filled
        if (!loginEmail || !loginPassword) {
            alert("Please fill in all fields.");
            event.preventDefault(); // Prevent the form from being submitted
            return;
        };

        // Send login data to the server using fetch
        fetch(`${ngrokURL}/login`, {
            method: "POST",  // Use POST method for sending data
            headers: {
                "Content-Type": "application/json",  // Send data as JSON
            },
            body: JSON.stringify({  // Stringify the data to send as JSON
                email: loginEmail,
                password: loginPassword
            })
        })
            .then(response => {
                // Handle response from the server
                if (!response.ok) {
                    throw new Error("Login failed");
                }
                return response.json();  // Parse JSON from the server response
            })
            .then(data => {
                console.log("Login successful:", data);
                window.location.href = '/vacancies';
            })
            .catch(error => {
                console.error("Error during login:", error);
                alert("Login failed. Please try again.");
            });

    });



    signupBttn.addEventListener("click", (event) => {

        signupEmail = document.getElementById("signupEmail").value.trim();
        passwordFirst = document.getElementById("passwordFirst").value.trim();
        passwordSecond = document.getElementById("passwordSecond").value.trim();
        name = document.getElementById("fullName").value.trim();
        role = document.getElementById("role").value.trim();



        // Check if the required fields are filled
        if (!signupEmail || !passwordFirst || !passwordSecond) {
            alert("Please fill in all fields.");
            event.preventDefault(); // Prevent the form from being submitted
            return;
        };

        // Check if both password fields match
        if (passwordFirst !== passwordSecond) {
            alert("Passwords do not match.");
            event.preventDefault(); // Prevent the form from being submitted
            return;
        };

        // Check email for @gmail.com
        const emailPattern = /^[a-zA-Z0-9._%+-]+@gmail\.com$/;
        if (!emailPattern.test(signupEmail)) {
            alert("Please enter a valid email address ending with @gmail.com.");
            event.preventDefault(); // Prevent the form from being submitted
            return;
        }

        fetch(`${ngrokURL}/send-code`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "ngrok-skip-browser-warning": "true"
            },
            body: JSON.stringify({
                email: signupEmail,
            }),
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
        })

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
            const response = await fetch(`${ngrokURL}/signup`, {
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

});
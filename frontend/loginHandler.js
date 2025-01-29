const ngrokURL = "https://465a-2a03-32c0-3003-5d87-7d78-11a5-9b7c-3598.ngrok-free.app";  

document.addEventListener('DOMContentLoaded', () => {

    const loginForm = document.getElementById("login-container");
    const signupForm = document.getElementById("signup-container");
    // const companyForm = document.getElementById("company-container");
    const authForm = document.getElementById("auth-container")
    const loginTag = document.getElementById("log-in-tag");//login anchor tag
    const signupTag = document.getElementById("sign-up-tag");//signup anchor tag

    const loginBttn = document.getElementById("log-in-bttn");
    const signupBttn = document.getElementById("sign-up-bttn");

    const userSection = document.getElementById("user-section");
    const companySection = document.getElementById("company-section");
    const orgzCheckBox = document.getElementById("Orgz_checkbox");

    var userType = "employee";
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


    orgzCheckBox.addEventListener("change", () => {
        console.log('Checkbox changed', orgzCheckBox.checked);  // Debugging line

        if (orgzCheckBox.checked) {
            // Checkbox is checked, show the company section and hide the user section
            userSection.style.display = "none";  // Hide the user section
            companySection.style.display = "block";  // Show the company section
            userType = "employeer";
        } else {
            // Checkbox is unchecked, show the user section and hide the company section
            userSection.style.display = "block";  // Show the user section
            companySection.style.display = "none";  // Hide the company section
            userType = "employee";
        }
    });


    // Initialize the form state on page load (user section should be visible by default)
    window.onload = () => {
        userSection.style.display = "block";  // Ensure user section is visible initially
        companySection.style.display = "none"; // Ensure company section is hidden initially
    };



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
        
        if (orgzCheckBox.checked) {
           name = document.getElementById("companyName");
        } else {
           name = document.getElementById("userFullName");
        }



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

        signupForm.style.display = "none";
        authForm.style.display = "block";
    });




    authFormBttn.addEventListener("click", (event) => {

        const authCode = document.getElementById("authCode").value.trim();
        // Send data to the server
        fetch(`${ngrokURL}/register`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",  // Sending data as JSON
            },
            body: JSON.stringify({
                name: name,
                email: signupEmail,
                password: passwordFirst,
                userType: userType,
                authCode: authCode,
            }),
        })
            .then(response => {
                if (!response.ok) {
                    throw new Error("Registration failed");
                }
                return response.json();  // Parse JSON from the response
            })
            .then(data => {
                console.log("Registration successful:", data);
              
            })
            .catch(error => {
                console.error("Error during registration:", error);
                alert("An error occurred. Please try again.");
            });

    });

});
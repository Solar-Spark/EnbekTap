

const loginForm = document.getElementById("login-container")
const signupForm = document.getElementById("signup-container")
const login = document.getElementById("log-in")//login anchor tag
const signup= document.getElementById("sign-up")//signup anchor tag


//when signup clicked
signup.addEventListener("click", ()=> {
    loginForm.style.display = "none";//hide login form
    signupForm.style.display = "block";//display signup form
    document.getElementById("loginForm").reset();//reset input values of login form
});

//when login clicked
login.addEventListener("click", ()=> {
    signupForm.style.display = "none";//hide signup form
    loginForm.style.display = "block";//display login form
    document.getElementById("signupForm").reset();//reset input values of signup form
            
});

const localhost = "http://localhost:8080";
document.addEventListener("DOMContentLoaded", function () {
  const username = document.getElementById("username");
  const fullname = document.getElementById("fullname");
  const email = document.getElementById("email");
  const role = document.getElementById("role");
  const bio = document.getElementById("bio");
  const token = localStorage.getItem("access_token");
  const logoutbttn = document.getElementById("logout");

  async function getUserData() {
    try {
      const response = await fetch(`${localhost}/auth/profile`, {
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
        credentials: "include",
      });

      if (!response.ok) {
        const errorData = await response.json();
        console.error("API Error:", errorData);
        throw new Error(errorData.error || "Failed to fetch user data");
      }

      return await response.json();
    } catch (error) {
      console.error("Failed to fetch profile:", error);
    }
  }

  logoutbttn.addEventListener("click", async (event) => {
    event.preventDefault(); // Prevent default form submission if applicable

    try {
      const token = localStorage.getItem("access_token"); // Make sure token is available

      const response = await fetch(`${localhost}/auth/logout`, {
        // ✅ Correct endpoint
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`, // ✅ Send token if required by backend
        },
        credentials: "include", // ✅ Required if using cookies
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
      localStorage.removeItem("role");

      // ✅ Redirect to login page
      window.location.href = "./login.html";
    } catch (error) {
      console.error("Logout error:", error);
      alert(error.message || "Logout failed. Please try again.");
    }
  });

  getUserData().then((userData) => {
    if (userData) {
      fullname.textContent = userData.FullName;
      email.textContent = userData.Email;
      role.textContent = userData.Role;
    }
  });
});

const ngrokURL = "https://8c7f-212-96-87-84.ngrok-free.app";
const userTableBody = document.getElementById("userTableBody");
const openCreateModal = document.getElementById("openCreateModal");
const closeCreateModal = document.getElementById("closeCreateModal");
const createModal = document.getElementById("createModal");
const createUserButton = document.getElementById("createUser");

const closeEditModal = document.getElementById("closeEditModal");
const editModal = document.getElementById("editModal");
const updateUserButton = document.getElementById("updateUser");

// Fetch all users
async function fetchUsers() {
    try {
        const response = await fetch(`${ngrokURL}/admin/users`, {
            headers: { "ngrok-skip-browser-warning": "true" }
        });
        const data = await response.json();
        
        userTableBody.innerHTML = "";
        data.users.forEach(user => {
            const row = document.createElement("tr");
            row.innerHTML = `
                <td>${user.UserID}</td>
                <td>${user.FullName}</td>
                <td>${user.Email}</td>
                <td>${user.Role}</td>
                <td>${user.Verified}</td>
                <td>${user.VerificationToken}</td>
                <td>
                    <button onclick="openEditUser(${user.UserID}, '${user.FullName}', '${user.Email}', '${user.Role}', '${user.Password}', '${user.Verified}', '${user.VerificationToken}')">Edit</button>
                    <button onclick="deleteUser(${user.UserID})">Delete</button>
                </td>
            `;
            userTableBody.appendChild(row);
        });
    } catch (error) {
        console.error("Error fetching users:", error);
    }
}

// Open Create User Modal
openCreateModal.addEventListener("click", () => {
    createModal.style.display = "block";
});

// Close Create User Modal
closeCreateModal.addEventListener("click", () => {
    createModal.style.display = "none";
});

// Create New User
createUserButton.addEventListener("click", async () => {
    const username = document.getElementById("createUsername").value;
    const email = document.getElementById("createEmail").value;
    const password = document.getElementById("createPassword").value;
    const role = document.getElementById("createRole").value;
    const verified = document.getElementById("createVerified").value;
    const verificationToken = document.getElementById("createVerificationCode").value;

    try {
        await fetch(`${ngrokURL}/admin/createuser`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "ngrok-skip-browser-warning": "true"
            },
            body: JSON.stringify({ FullName: username, Email: email, Password: password, Role: role, Verified: verified, VerificationToken: verificationToken })
        });

        createModal.style.display = "none";
        fetchUsers();
    } catch (error) {
        console.error("Error creating user:", error);
    }
});

// Open Edit User Modal
function openEditUser(id, username, email, role, password, verified, verificationToken) {
    document.getElementById("editUserId").value = id;
    document.getElementById("editUsername").value = username;
    document.getElementById("editEmail").value = email;
    document.getElementById("editRole").value = role;
    document.getElementById("editPassword").value = password;
    document.getElementById("editVerified").value = verified;
    document.getElementById("editVerificationCode").value = verificationToken;

    editModal.style.display = "block";
}

// Close Edit User Modal
closeEditModal.addEventListener("click", () => {
    editModal.style.display = "none";
});

// Update User
updateUserButton.addEventListener("click", async () => {
    const id = document.getElementById("editUserId").value;
    const username = document.getElementById("editUsername").value;
    const email = document.getElementById("editEmail").value;
    const role = document.getElementById("editRole").value;
    const password = document.getElementById("editPassword").value;
    const verified = document.getElementById("editVerified").value;
    const verificationToken = document.getElementById("editVerificationCode").value;

    try {
        await fetch(`${ngrokURL}/admin/updateuser?id=${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "ngrok-skip-browser-warning": "true"
            },
            body: JSON.stringify({ FullName: username, Email: email, Role: role, Password: password, Verified: verified, VerificationToken: verificationToken })
        });

        editModal.style.display = "none";
        fetchUsers();
    } catch (error) {
        console.error("Error updating user:", error);
    }
});

// Delete User
async function deleteUser(id) {
    await fetch(`${ngrokURL}/admin/deleteuser?id=${id}`, { method: "DELETE" });
    fetchUsers();
}

// Load users on page load
fetchUsers();

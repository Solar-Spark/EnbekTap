document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("jobForm");


    const searchButton = document.querySelector("button#search");
    const searchInput = document.querySelector("input#search");
    const jobCardsContainer = document.getElementById("jobCards");

    const createPostButton = document.querySelector("#createPost");
    const createModal = document.getElementById("createModal");
    const closeCreateModal = document.getElementById("closeCreateModal");

    const editModal = document.getElementById("editModal");
    const closeEditModal = document.getElementById("closeEditModal");


  

    const jobTypeDropdown = document.getElementById("jobTypeDropdown");
    const sortDropdown = document.getElementById("sortDropdown");
    const resetButton = document.getElementById("reset");


    const loginForm = document.getElementById("login-form")
    const signupForm = document.getElementById("signup-form")
    const login = document.getElementById("log-in")
    const signup= document.getElementById("sign-up")


    signup.addEventListener("click", ()=> {
        loginForm.style.display = "none";
        signupForm.style.display = "block";
        
    });


    login.addEventListener("click", ()=> {
        signupForm.style.display = "none";
        loginForm.style.display = "block";

        
    });


    createPostButton.addEventListener("click", ()=> {
        createModal.style.display = "block";

    })

    closeCreateModal.addEventListener("click", () => {
        createModal.style.display = "none";
    });

    form.addEventListener("submit", async (event) => {
        event.preventDefault();
        const jobName = document.getElementById("jobName").value;
        const salary = document.getElementById("salary").value;
        const jobType = document.querySelector("input[name='JobType']:checked").value;
        const description = document.getElementById("description").value;

        try {
            const response = await fetch("http://localhost:8080/createvacancy", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    Vacancy: jobName,
                    Salary: parseInt(salary),
                    JobType: jobType,
                    Description: description,
                }),
            });

            if (!response.ok) {
                const text = await response.text();
                throw new Error(`Server Error: ${response.status} - ${text}`);
            }

            const result = await response.json();
            alert(result.message);
            form.reset();
            createModal.style.display = "none";
            loadJobCards();
        } catch (error) {
            console.error("Error submitting form:", error);
            alert("Error submitting the form. Please try again.");
        }
    });

    async function loadJobCards() {
        const jobType = jobTypeDropdown.value;
        const sortBy = sortDropdown.value;
        const url = new URL("http://localhost:8080/vacancies");
    
        if (jobType) {
            url.searchParams.append("jobType", jobType);
        }
        if (sortBy) {
            url.searchParams.append("sort", sortBy);
        }
    
        try {
            const response = await fetch(url);
    
            if (!response.ok) {
                const text = await response.text();
                throw new Error(`Server Error: ${response.status} - ${text}`);
            }
    
            const vacancies = await response.json();
            renderJobCards(vacancies);
        } catch (error) {
            console.error("Error fetching vacancies:", error);
            alert("Error fetching job postings. Please try again.");
        }
    }
    

  // Render job cards on the page
  function renderJobCards(vacancies) {
    jobCardsContainer.innerHTML = "";
    if (vacancies.length === 0) {
        jobCardsContainer.innerHTML = "<h1>No vacancies posted</h1>";
    }
    vacancies.forEach((vacancy) => {
        const card = document.createElement("div");
        card.className = "job-card";
        card.innerHTML = `
            <h2>${vacancy.Vacancy}</h2>
            <p><strong>Salary:</strong> $${vacancy.Salary}</p>
            <p><strong>Type:</strong> ${vacancy.JobType}</p>
            <p>${vacancy.Description}</p>
            <button class="edit-button" data-id="${vacancy.VacancyID}">Edit</button>
            <button class="delete-button" data-id="${vacancy.VacancyID}">Delete</button>
        `;
        jobCardsContainer.appendChild(card);
    });

    document.querySelectorAll(".edit-button").forEach(button =>
        button.addEventListener("click", openEditModal)
    );

    document.querySelectorAll(".delete-button").forEach(button =>
        button.addEventListener("click", deleteVacancy)
    );
}

    function openEditModal(event) {
        const id = event.target.dataset.id;
        fetch(`http://localhost:8080/vacancy?id=${id}`)
            .then(response => {
                if (!response.ok) {
                    throw new Error(`Failed to fetch vacancy. Status: ${response.status}`);
                }
                return response.json();
            })
            .then(vacancy => {
                document.getElementById("editVacancyID").value = id;
                document.getElementById("editJobName").value = vacancy.Vacancy;
                document.getElementById("editSalary").value = vacancy.Salary;
    
                // Ensure correct JobType radio button is selected
                const fullTimeButton = document.getElementById("editFullTime");
                const partTimeButton = document.getElementById("editPartTime");
    
                if (vacancy.JobType === "Full-Time") {
                    fullTimeButton.checked = true;
                } else if (vacancy.JobType === "Part-Time") {
                    partTimeButton.checked = true;
                }
    
                document.getElementById("editDescription").value = vacancy.Description;
                editModal.style.display = "block"; // Open modal
            })
            .catch(error => {
                console.error("Error loading vacancy for editing:", error);
                alert("Failed to load vacancy details for editing. Please try again.");
            });
    }
    
    

    closeEditModal.addEventListener("click", () => {
        editModal.style.display = "none";
    });

    editForm.addEventListener("submit", async (event) => {
        event.preventDefault();
        const id = document.getElementById("editVacancyID").value;
        const jobName = document.getElementById("editJobName").value;
        const salary = document.getElementById("editSalary").value;
        const jobType = document.querySelector("input[name='EditJobType']:checked").value;
        const description = document.getElementById("editDescription").value;
    
        try {
            const response = await fetch(`http://localhost:8080/updatevacancy?id=${id}`, { // Changed to use query parameter
                method: "PUT",
                headers: { 
                    "Content-Type": "application/json" 
                },
                body: JSON.stringify({
                    Vacancy: jobName,
                    Salary: parseInt(salary),
                    JobType: jobType,
                    Description: description,
                }),
            });
    
            if (!response.ok) {
                const text = await response.text();
                throw new Error(`Server Error: ${response.status} - ${text}`);
            }
    
            const result = await response.json();
            alert(result.message);
            editModal.style.display = "none";
            loadJobCards();
        } catch (error) {
            console.error("Error editing vacancy:", error);
            alert("Error editing the vacancy. Please try again.");
        }
    });
    async function deleteVacancy(event) {
        const id = event.target.dataset.id;
        if (!confirm("Are you sure you want to delete this vacancy?")) return;
    
        try {
            const response = await fetch(`http://localhost:8080/deletevacancy?id=${id}`, { // Changed to use query parameter
                method: "DELETE",
                headers: {
                    'Content-Type': 'application/json'
                }
            });
    
            if (!response.ok) {
                const text = await response.text();
                throw new Error(`Server Error: ${response.status} - ${text}`);
            }
    
            const result = await response.json();
            alert(result.message);
            loadJobCards();
        } catch (error) {
            console.error("Error deleting vacancy:", error);
            alert("Error deleting the vacancy. Please try again.");
        }
    }

searchButton.addEventListener("click", async () => {
    const searchQuery = searchInput.value.trim();
    if (!searchQuery || isNaN(searchQuery)) {
        alert("Please enter a valid numeric vacancy ID");
        return;
    }

    try {
        const response = await fetch(`http://localhost:8080/vacancy?id=${searchQuery}`);
        if (!response.ok) {
            const text = await response.text();
            throw new Error(`Server Error: ${response.status} - ${text}`);
        }
        const vacancy = await response.json();
        if (vacancy) {
            renderJobCard(vacancy);
        } else {
            alert("No vacancy found with that ID.");
        }
    } catch (error) {
        console.error("Error during search:", error);
        alert("Error searching for vacancy. Please try again.");
    }
});


    function renderJobCard(vacancy) {
        jobCardsContainer.innerHTML = "";
        const card = document.createElement("div");
        card.className = "job-card";
        card.innerHTML = `
            <h2>${vacancy.Vacancy}</h2>
            <p><strong>Salary:</strong> $${vacancy.Salary}</p>
            <p><strong>Type:</strong> ${vacancy.JobType}</p>
            <p>${vacancy.Description}</p>
            <button class="edit-button" data-id="${vacancy.VacancyID}">Edit</button>
            <button class="delete-button" data-id="${vacancy.VacancyID}">Delete</button>
        `;
        jobCardsContainer.appendChild(card);

        document.querySelectorAll(".edit-button").forEach(button =>
            button.addEventListener("click", openEditModal)
        );
    
        document.querySelectorAll(".delete-button").forEach(button =>
            button.addEventListener("click", deleteVacancy)
        );
    }

    searchInput.addEventListener("blur", () => {
        if (!searchInput.value.trim()) {
            loadJobCards();
        }
    });

     // Reset button functionality to reset the form and reload job cards
     resetButton.addEventListener("click", () => {
        form.reset();
        jobTypeDropdown.selectedIndex = 0;  // Reset job type dropdown to the default option
        sortDropdown.selectedIndex = 0;  // Reset sort dropdown to the default option
        loadJobCards();  // Reload all job cards without any filters
    });

    loadJobCards();
});

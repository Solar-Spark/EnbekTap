document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("jobForm");
    const searchButton = document.querySelector("button#search");
    const searchInput = document.querySelector("input#search");
    const jobCardsContainer = document.getElementById("jobCards");

    const editModal = document.getElementById("editModal");
    const closeEditModal = document.getElementById("closeEditModal");
    const editForm = document.getElementById("editForm");

    const jobTypeDropdown = document.getElementById("jobTypeDropdown");
    const sortDropdown = document.getElementById("sortDropdown");
    const resetButton = document.getElementById("reset");

    // Add these variables at the top of your script
    let currentPage = 1;
    const pageSize = 9;

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
        } catch (error) {
            console.error("Error submitting form:", error);
            alert("Error submitting the form. Please try again.");
        }
    });

    // Update the loadJobCards function
    async function loadJobCards() {
        const jobType = jobTypeDropdown.value;
        const sortBy = sortDropdown.value;
        const url = new URL("http://localhost:8080/vacancies");
        
        url.searchParams.append("page", window.currentPage);  // Use window.currentPage
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
    
            const data = await response.json();
            renderJobCards(data.vacancies);
            renderPagination(data.currentPage, data.totalPages);
        } catch (error) {
            console.error("Error fetching vacancies:", error);
            alert("Error fetching job postings. Please try again.");
        }
    }
    
    // Add pagination rendering function
    function renderPagination(currentPage, totalPages) {
        const paginationContainer = document.getElementById("pagination");
        paginationContainer.innerHTML = "";
    
        // Previous button
        if (currentPage > 1) {
            const prevButton = document.createElement("button");
            prevButton.textContent = "Previous";
            prevButton.addEventListener("click", () => {
                window.currentPage--;  // Use window to ensure global scope
                loadJobCards();
            });
            paginationContainer.appendChild(prevButton);
        }
    
        // Page numbers
        for (let i = 1; i <= totalPages; i++) {
            const pageButton = document.createElement("button");
            pageButton.textContent = i;
            pageButton.classList.toggle("active", i === currentPage);
            pageButton.addEventListener("click", () => {
                window.currentPage = i;  // Use window to ensure global scope
                loadJobCards();
            });
            paginationContainer.appendChild(pageButton);
        }
    
        // Next button
        if (currentPage < totalPages) {
            const nextButton = document.createElement("button");
            nextButton.textContent = "Next";
            nextButton.addEventListener("click", () => {
                window.currentPage++;  // Use window to ensure global scope
                loadJobCards();
            });
            paginationContainer.appendChild(nextButton);
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
    
                if (vacancy.JobType === "Full Time") {
                    fullTimeButton.checked = true;
                } else if (vacancy.JobType === "Part Time") {
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
        `;
        jobCardsContainer.appendChild(card);
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

    jobTypeDropdown.addEventListener('change', () => {
        currentPage = 1;
        loadJobCards();
    });

    sortDropdown.addEventListener('change', () => {
        currentPage = 1;
        loadJobCards();
    });

    loadJobCards();
});

document.getElementById('supportForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    
    const formData = new FormData();
    formData.append('subject', document.getElementById('subject').value);
    formData.append('message', document.getElementById('message').value);
    
    const attachments = document.getElementById('attachments').files;
    for (let i = 0; i < attachments.length; i++) {
        formData.append('attachments', attachments[i]);
    }

    try {
        const response = await fetch('http://localhost:8080/support/contact', {
            method: 'POST',
            body: formData
        });

        if (!response.ok) {
            throw new Error('Failed to send message');
        }

        const result = await response.json();
        alert(result.message);
        document.getElementById('supportForm').reset();
    } catch (error) {
        console.error('Error:', error);
        alert('Failed to send message. Please try again.');
    }
});

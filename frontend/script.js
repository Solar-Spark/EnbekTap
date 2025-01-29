const ngrokURL = "https://465a-2a03-32c0-3003-5d87-7d78-11a5-9b7c-3598.ngrok-free.app";  

document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("jobForm");
    const searchButton = document.querySelector("button#search");
    const searchInput = document.querySelector("input#search");
    const jobCardsContainer = document.getElementById("jobCards");

    const editModal = document.getElementById("editModal");
    const closeEditModal = document.getElementById("closeEditModal");
    const editForm = document.getElementById("editForm");

    const createPostButton = document.querySelector("#createPost");
    const createModal = document.getElementById("createModal");
    const closeCreateModal = document.getElementById("closeCreateModal");

    const jobTypeDropdown = document.getElementById("jobTypeDropdown");
    const sortDropdown = document.getElementById("sortDropdown");
    const resetButton = document.getElementById("reset");

    let currentPage = 1;
    const pageSize = 9;

    createPostButton.addEventListener("click", () => {
        createModal.style.display = "block";
    });

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
            const response = await fetch(`${ngrokURL}/createvacancy`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "ngrok-skip-browser-warning": "true"
                },
                body: JSON.stringify({
                    Vacancy: jobName,
                    Salary: parseInt(salary),
                    JobType: jobType,
                    Description: description,
                }),
            });

            if (!response.ok) {
                throw new Error(`Server Error: ${response.status} - ${await response.text()}`);
            }

            const result = await response.json();
            alert(result.message);
            createModal.style.display = "none";
            form.reset();
        } catch (error) {
            console.error("Error submitting form:", error);
            alert("Error submitting the form. Please try again.");
        }
    });

    async function loadJobCards() {
        const jobType = jobTypeDropdown.value;
        const sortBy = sortDropdown.value;
        const url = new URL(`${ngrokURL}/vacancies`);
        
        url.searchParams.append("page", currentPage);
        if (jobType) {
            url.searchParams.append("jobType", jobType);
        }
        if (sortBy) {
            url.searchParams.append("sort", sortBy);
        }
    
        try {
            const response = await fetch(url, {
                headers: { "ngrok-skip-browser-warning": "true" }
            });

            if (!response.ok) {
                throw new Error(`Server Error: ${response.status} - ${await response.text()}`);
            }

            const data = await response.json();
            renderJobCards(data.vacancies);
            renderPagination(data.currentPage, data.totalPages);
        } catch (error) {
            console.error("Error fetching vacancies:", error);
            alert("Error fetching job postings. Please try again.");
        }
    }

    function renderPagination(currentPage, totalPages) {
        const paginationContainer = document.getElementById("pagination");
        paginationContainer.innerHTML = "";

        if (currentPage > 1) {
            const prevButton = document.createElement("button");
            prevButton.textContent = "Previous";
            prevButton.addEventListener("click", () => {
                currentPage--;
                loadJobCards();
            });
            paginationContainer.appendChild(prevButton);
        }

        for (let i = 1; i <= totalPages; i++) {
            const pageButton = document.createElement("button");
            pageButton.textContent = i;
            pageButton.classList.toggle("active", i === currentPage);
            pageButton.addEventListener("click", () => {
                currentPage = i;
                loadJobCards();
            });
            paginationContainer.appendChild(pageButton);
        }

        if (currentPage < totalPages) {
            const nextButton = document.createElement("button");
            nextButton.textContent = "Next";
            nextButton.addEventListener("click", () => {
                currentPage++;
                loadJobCards();
            });
            paginationContainer.appendChild(nextButton);
        }
    }

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
            `;
            jobCardsContainer.appendChild(card);
        });
    }

    searchButton.addEventListener("click", async () => {
        const searchQuery = searchInput.value.trim();
        if (!searchQuery || isNaN(searchQuery)) {
            alert("Please enter a valid numeric vacancy ID");
            return;
        }

        try {
            const response = await fetch(`${ngrokURL}/vacancy?id=${searchQuery}`, {
                headers: { "ngrok-skip-browser-warning": "true" }
            });

            if (!response.ok) {
                throw new Error(`Server Error: ${response.status} - ${await response.text()}`);
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

    resetButton.addEventListener("click", () => {
        form.reset();
        jobTypeDropdown.selectedIndex = 0;
        sortDropdown.selectedIndex = 0;
        loadJobCards();
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
        const response = await fetch(`${ngrokURL}/support/contact`, {
            method: 'POST',
            headers: { "ngrok-skip-browser-warning": "true" },
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

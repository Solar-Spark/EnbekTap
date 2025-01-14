


// loadJobCards.js
export const loadJobCards = async (openEditModal, deleteVacancy) => {
    const jobCardsContainer = document.getElementById("jobCards");
    const jobTypeDropdown = document.getElementById("jobTypeDropdown");
    const sortDropdown = document.getElementById("sortDropdown");

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
        renderJobCards(vacancies, openEditModal, deleteVacancy);
    } catch (error) {
        console.error("Error fetching vacancies:", error);
        alert("Error fetching job postings. Please try again.");
    }
};

// Render job cards on the page
const renderJobCards = (vacancies, openEditModal, deleteVacancy) => {
    const jobCardsContainer = document.getElementById("jobCards");
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

     // Make sure event listeners are added to newly created elements
     document.querySelectorAll(".edit-button").forEach(button =>
        button.addEventListener("click", openEditModal)
    );

    document.querySelectorAll(".delete-button").forEach(button =>
        button.addEventListener("click", deleteVacancy)
    );
};

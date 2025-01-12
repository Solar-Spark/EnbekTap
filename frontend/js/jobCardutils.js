export async function loadJobCards() {
    const jobType = document.getElementById("jobTypeDropdown").value;
    const sortBy = document.getElementById("sortDropdown").value;
    const url = new URL("http://localhost:8080/vacancies");

    if (jobType) url.searchParams.append("jobType", jobType);
    if (sortBy) url.searchParams.append("sort", sortBy);

    try {
        const response = await fetch(url);
        if (!response.ok) throw new Error("Failed to fetch vacancies");

        const vacancies = await response.json();
        renderJobCards(vacancies);
    } catch (error) {
        console.error("Error fetching vacancies:", error);
        alert("Error fetching job postings. Please try again.");
    }
}

export function renderJobCards(vacancies) {
    const jobCardsContainer = document.getElementById("jobCards");
    jobCardsContainer.innerHTML = "";

    if (vacancies.length === 0) {
        jobCardsContainer.innerHTML = "<h1>No vacancies posted</h1>";
        return;
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

    setupCardActions();
}

function setupCardActions() {
    document.querySelectorAll(".edit-button").forEach(button =>
        button.addEventListener("click", (event) => {
            import("./modalHandlers.js").then(module => module.openEditModal(event));
        })
    );

    document.querySelectorAll(".delete-button").forEach(button =>
        button.addEventListener("click", deleteVacancy)
    );
}

export async function deleteVacancy(event) {
    const id = event.target.dataset.id;
    if (!confirm("Are you sure you want to delete this vacancy?")) return;

    try {
        const response = await fetch(`http://localhost:8080/deletevacancy?id=${id}`, {
            method: "DELETE",
            headers: { "Content-Type": "application/json" },
        });

        if (!response.ok) throw new Error("Failed to delete the vacancy");

        const result = await response.json();
        alert(result.message);
        loadJobCards();
    } catch (error) {
        console.error("Error deleting vacancy:", error);
        alert("Error deleting the vacancy. Please try again.");
    }
}

export function searchVacancy(query) {
    if (!query || isNaN(query)) {
        alert("Please enter a valid numeric vacancy ID");
        return;
    }

    fetch(`http://localhost:8080/vacancy?id=${query}`)
        .then(response => {
            if (!response.ok) throw new Error("Failed to fetch vacancy");
            return response.json();
        })
        .then(renderJobCard)
        .catch(error => {
            console.error("Error searching vacancy:", error);
            alert("Error searching for vacancy. Please try again.");
        });
}

export function resetFilters() {
    document.getElementById("jobForm").reset();
    document.getElementById("jobTypeDropdown").selectedIndex = 0;
    document.getElementById("sortDropdown").selectedIndex = 0;
    loadJobCards();
}

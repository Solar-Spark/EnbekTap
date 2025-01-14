// searchJobCards.js
export const searchVacancy = async (event, loadJobCards) => {
    const searchQuery = document.querySelector("input#search").value.trim();
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
};

export const renderJobCard = (vacancy) => {
    const jobCardsContainer = document.getElementById("jobCards");
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
};



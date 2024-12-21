document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("jobForm");
    const searchButton = document.querySelector("button#search"); // Search button
    const searchInput = document.querySelector("input#search"); // Search input field
    const jobCardsContainer = document.getElementById("jobCards");

    // Handle form submission via POST
    form.addEventListener("submit", async (event) => {
        event.preventDefault();

        const jobName = document.getElementById("jobName").value;
        const salary = document.getElementById("salary").value;
        const jobType = document.querySelector("input[name='JobType']:checked").value;
        const description = document.getElementById("description").value;

        try {
            const response = await fetch("http://localhost:8080/vacancies", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
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
            form.reset();
            loadJobCards(); // Reload job cards
        } catch (error) {
            console.error("Error submitting form:", error);
            alert("Error submitting the form. Please try again.");
        }
    });

    // Fetch and display vacancies via GET
    async function loadJobCards() {
        try {
            const response = await fetch("http://localhost:8080/vacancies");
            if (!response.ok) {
                const text = await response.text();
                throw new Error(`Server Error: ${response.status} - ${text}`);
            }

            const vacancies = await response.json();
            renderJobCards(vacancies);
        } catch (error) {
            console.error("Error fetching vacancies:", error);
        }
    }

    // Render all job cards
    function renderJobCards(vacancies) {
        jobCardsContainer.innerHTML = ""; // Clear existing cards
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

    // Search for a job by ID
    searchButton.addEventListener("click", async () => {
        const searchQuery = searchInput.value.trim();
        if (!searchQuery || isNaN(searchQuery)) {
            alert("Please enter a valid numeric vacancy ID");
            return;
        }

        try {
            const response = await fetch(`http://localhost:8080/vacancies?id=${searchQuery}`);
            if (!response.ok) {
                const text = await response.text();
                throw new Error(`Server Error: ${response.status} - ${text}`);
            }

            const vacancy = await response.json();
            renderJobCard(vacancy);
        } catch (error) {
            console.error("Error during search:", error);
            alert("Error searching for vacancy. Please try again.");
        }
    });

    // Render a single job card
    function renderJobCard(vacancy) {
        jobCardsContainer.innerHTML = ""; // Clear existing cards
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

    // Load all vacancies when the search bar is empty and loses focus
    searchInput.addEventListener("blur", () => {
        if (!searchInput.value.trim()) {
            loadJobCards(); // Load all vacancies
        }
    });

    // Initial load of all vacancies
    loadJobCards();
});

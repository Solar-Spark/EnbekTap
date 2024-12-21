document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("jobForm");
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

    // Render job cards
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

    // Initial load
    loadJobCards();
});

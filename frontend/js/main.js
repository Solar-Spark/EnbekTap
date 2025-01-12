import { loadJobCards } from "./jobCards.js";
import { dropdownHandler } from "./handlers.js";
import { openEditModal } from "./editModal.js";
import { deleteVacancy } from "./delete.js";
import { searchVacancy, renderJobCard } from "./search.js";

document.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById("jobForm");
    const searchButton = document.querySelector("button#search");
    const searchInput = document.querySelector("input#search");
    const jobCardsContainer = document.getElementById("jobCards");

    const jobTypeDropdown = document.getElementById("jobTypeDropdown");
    const sortDropdown = document.getElementById("sortDropdown");
    const resetButton = document.getElementById("reset");

    // Form submission event
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
            loadJobCards();
        } catch (error) {
            console.error("Error submitting form:", error);
            alert("Error submitting the form. Please try again.");
        }
    });

    // Dropdown event listeners for filtering and sorting
    jobTypeDropdown.addEventListener("change", () => dropdownHandler(jobTypeDropdown, sortDropdown, loadJobCards));
    sortDropdown.addEventListener("change", () => dropdownHandler(jobTypeDropdown, sortDropdown, loadJobCards));

    // Search button click event
    searchButton.addEventListener("click", async () => {
        const searchQuery = searchInput.value.trim();
        if (!searchQuery || isNaN(searchQuery)) {
            alert("Please enter a valid numeric vacancy ID");
            return;
        }
        await searchVacancy(searchQuery, renderJobCard, jobCardsContainer);
    });

    // Reset button functionality
    resetButton.addEventListener("click", () => {
        form.reset();
        jobTypeDropdown.selectedIndex = 0; // Reset job type dropdown
        sortDropdown.selectedIndex = 0;   // Reset sort dropdown
        loadJobCards();                  // Reload job cards without filters
    });

    // Blur event to reload all job cards when search input loses focus
    searchInput.addEventListener("blur", () => {
        if (!searchInput.value.trim()) {
            loadJobCards();
        }
    });

    // Initial loading of job cards
    loadJobCards();
});

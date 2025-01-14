

export function initSort() {
    const jobTypeDropdown = document.getElementById("jobTypeDropdown");
    const sortDropdown = document.getElementById("sortDropdown");
    const resetButton = document.getElementById("reset");


    // Reset button functionality to reset the form and reload job cards
    resetButton.addEventListener("click", () => {
        form.reset();
        jobTypeDropdown.selectedIndex = 0;  // Reset job type dropdown to the default option
        sortDropdown.selectedIndex = 0;  // Reset sort dropdown to the default option
        loadJobCards();  // Reload all job cards without any filters
    });
}
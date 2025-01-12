export function setupModals() {
    const editModal = document.getElementById("editModal");
    const closeEditModal = document.getElementById("closeEditModal");

    closeEditModal.addEventListener("click", () => {
        editModal.style.display = "none";
    });
}

export function openEditModal(event) {
    const id = event.target.dataset.id;
    fetch(`http://localhost:8080/updatevacancy?id=${id}`)
        .then(response => {
            if (!response.ok) throw new Error("Failed to fetch vacancy");
            return response.json();
        })
        .then(vacancy => {
            document.getElementById("editVacancyID").value = id;
            document.getElementById("editJobName").value = vacancy.Vacancy;
            document.getElementById("editSalary").value = vacancy.Salary;
            document.getElementById("editDescription").value = vacancy.Description;

            const fullTimeButton = document.getElementById("editFullTime");
            const partTimeButton = document.getElementById("editPartTime");

            fullTimeButton.checked = vacancy.JobType === "Full Time";
            partTimeButton.checked = vacancy.JobType === "Part Time";

            document.getElementById("editModal").style.display = "block";
        })
        .catch(error => {
            console.error("Error loading vacancy for editing:", error);
            alert("Failed to load vacancy details for editing. Please try again.");
        });
}

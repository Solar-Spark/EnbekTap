// editJobCards.js
export const openEditModal = (event) => {
    const id = event.target.dataset.id;
    fetch(`http://localhost:8080/vacancy?id=${id}`)
        .then(response => response.json())
        .then(vacancy => {
            document.getElementById("editVacancyID").value = id;
            document.getElementById("editJobName").value = vacancy.Vacancy;
            document.getElementById("editSalary").value = vacancy.Salary;

            const fullTimeButton = document.getElementById("editFullTime");
            const partTimeButton = document.getElementById("editPartTime");
            
            fullTimeButton.checked = vacancy.JobType === "Full-Time";
            partTimeButton.checked = vacancy.JobType === "Part-Time";
            document.getElementById("editDescription").value = vacancy.Description;
            document.getElementById("editModal").style.display = "block"; // Open modal
        })
        .catch(error => {
            console.error("Error loading vacancy for editing:", error);
            alert("Failed to load vacancy details for editing. Please try again.");
        });
};

export const handleEditJobFormSubmit = async (event, loadJobCards) => {
    event.preventDefault();
    const id = document.getElementById("editVacancyID").value;
    const jobName = document.getElementById("editJobName").value;
    const salary = document.getElementById("editSalary").value;
    const jobType = document.querySelector("input[name='EditJobType']:checked").value;
    const description = document.getElementById("editDescription").value;

    try {
        const response = await fetch(`http://localhost:8080/updatevacancy?id=${id}`, {
            method: "PUT",
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
        document.getElementById("editModal").style.display = "none";
        loadJobCards();
    } catch (error) {
        console.error("Error editing vacancy:", error);
        alert("Error editing the vacancy. Please try again.");
    }
};

export const closeEditModal = () => {
    document.getElementById("editModal").style.display = "none";
};

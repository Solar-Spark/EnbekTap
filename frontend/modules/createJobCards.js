// createJobCards.js
export const handleCreateJobFormSubmit = async (event, form, loadJobCards) => {
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
        document.getElementById("createModal").style.display = "none";
        loadJobCards();
    } catch (error) {
        console.error("Error submitting form:", error);
        alert("Error submitting the form. Please try again.");
    }
};

export const openCreateModal = () => {
    document.getElementById("createModal").style.display = "block";
};

export const closeCreateModal = () => {
    document.getElementById("createModal").style.display = "none";
};

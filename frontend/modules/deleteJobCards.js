// deleteJobCards.js
export const deleteVacancy = async (event, loadJobCards) => {
    const id = event.target.dataset.id;
    if (!confirm("Are you sure you want to delete this vacancy?")) return;

    try {
        const response = await fetch(`http://localhost:8080/deletevacancy?id=${id}`, {
            method: "DELETE",
            headers: {
                'Content-Type': 'application/json'
            }
        });

        if (!response.ok) {
            const text = await response.text();
            throw new Error(`Server Error: ${response.status} - ${text}`);
        }

        const result = await response.json();
        alert(result.message);
        loadJobCards();
    } catch (error) {
        console.error("Error deleting vacancy:", error);
        alert("Error deleting the vacancy. Please try again.");
    }
};

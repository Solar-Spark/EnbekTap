document.getElementById('jobForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent form from submitting normally

    // Get values from the form fields
    const jobName = document.getElementById('jobName').value;
    const salary = document.getElementById('salary').value;
    const jobType = document.querySelector('input[name="jobType"]:checked')?.value; // Get the selected radio button value
    const description = document.getElementById('description').value;

    // Check if a job type is selected
    if (!jobType) {
        alert('Please select a job type (Full Time or Part Time).');
        return; // Prevent card creation if no job type is selected
    }

    // Create a new card
    const newCard = document.createElement('div');
    newCard.classList.add('job-card');

    newCard.innerHTML = `
        <h3>${jobName}</h3>
        <p class="salary">Salary: $${salary}</p>
        <p class="job-type">Job Type: ${jobType}</p>
        <p class="description">${description}</p>
        <button class="edit-btn">Edit</button>
        <button class="delete-btn">Delete</button>
    `;

    // Append the new card to the job cards container
    document.getElementById('jobCards').appendChild(newCard);

    // Add event listener for the "Delete" button
    newCard.querySelector('.delete-btn').addEventListener('click', function() {
        newCard.remove(); // Remove the card from the DOM
    });

    // Add event listener for the "Edit" button
    newCard.querySelector('.edit-btn').addEventListener('click', function() {
        // Prefill the form with the current card's data
        document.getElementById('jobName').value = jobName;
        document.getElementById('salary').value = salary;
        document.querySelector(`input[name="jobType"][value="${jobType}"]`).checked = true;
        document.getElementById('description').value = description;

        // Remove the card when editing (optional)
        newCard.remove();
    });

    // Clear the form after creating the card
    document.getElementById('jobForm').reset();
});

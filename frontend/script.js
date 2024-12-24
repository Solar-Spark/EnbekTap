document.getElementById('jobForm').addEventListener('submit', function(event) {
    event.preventDefault(); // Prevent form from submitting normally

    const jobData = {
        Vacancy: document.getElementById('jobName').value,
        Salary: parseInt(document.getElementById('salary').value, 10),
        JobType: document.querySelector('input[name="JobType"]:checked').value,
        Description: document.getElementById('description').value
    };

    fetch('/vacancies', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        
        body: JSON.stringify(jobData)
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.text(); // Use text() to handle empty responses
    })
    .then(text => {
        const data = text ? JSON.parse(text) : {}; // Parse only if text is not empty
        if (data.status === 'success') {
            alert('Job successfully created!');
            // Optionally, update the UI to show the new job card
        } else {
            alert('Error: ' + data.message);
        }
    })
    .catch(error => {
        console.error('Error:', error);
    });

    // Check if a job type is selected
    if (!jobData.JobType) {
        alert('Please select a job type (Full Time or Part Time).');
        return; // Prevent card creation if no job type is selected
    }

    // Create a new card
    const newCard = document.createElement('div');
    newCard.classList.add('job-card');

    newCard.innerHTML = `
        <h3>${jobData.jobName}</h3>
        <p class="salary">Salary: $${jobData.salary}</p>
        <p class="job-type">Job Type: ${jobData.JobType}</p>
        <p class="description">${jobData.description}</p>
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

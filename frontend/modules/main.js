// main.js
import { loadJobCards } from './loadJobCards.js';
import { handleCreateJobFormSubmit, openCreateModal, closeCreateModal } from './createJobCards.js';
import { deleteVacancy } from './deleteJobCards.js';
import { openEditModal, handleEditJobFormSubmit, closeEditModal } from './editJobCards.js';
import { searchVacancy } from './searchJobCards.js';

document.getElementById("createPost").addEventListener("click", openCreateModal);
document.getElementById("closeCreateModal").addEventListener("click", closeCreateModal);

document.getElementById("jobForm").addEventListener("submit", (event) => handleCreateJobFormSubmit(event, document.getElementById("jobForm"), loadJobCards));
document.getElementById("search").addEventListener("click", (event) => searchVacancy(event, loadJobCards(openEditModal, deleteVacancy)));

document.getElementById("editForm").addEventListener("submit", (event) => handleEditJobFormSubmit(event, loadJobCards(openEditModal, deleteVacancy)));
document.getElementById("closeEditModal").addEventListener("click", closeEditModal);

loadJobCards();

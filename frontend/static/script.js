function toggleEdit(button) {
    const form = button.closest('form');
    const todoTitle = form.querySelector('.todo-title');
    const editTitle = form.querySelector('.edit-title');
    const updateBtn = form.querySelector('.update-btn');
    form.classList.toggle('edit-mode');
    if (form.classList.contains('edit-mode')) {
      todoTitle.style.display = 'none';
      editTitle.style.display = 'inline-block';
      updateBtn.style.display = 'inline-block';
      editTitle.value = todoTitle.textContent;
    } else {
      todoTitle.style.display = 'inline-block';
      editTitle.style.display = 'none';
      updateBtn.style.display = 'none';
    }
  }
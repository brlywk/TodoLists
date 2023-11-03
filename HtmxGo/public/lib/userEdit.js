// Checks if a userId has been stored in localStorage, otherwise
// set it
function getOrCreateUserId() {
  let userId = localStorage.getItem("userId");

  if (!userId) {
    userId = crypto.randomUUID();
    localStorage.setItem("userId", userId);
  }

  return userId;
}

// Allows the user to specify a userId
function setUserId(userId) {
  if (!userId) return;

  localStorage.setItem("userId", userId);
}

// initialise userId display and edit form
document.addEventListener("DOMContentLoaded", () => {
  // fire trigger our 'triggerLoad' event to have the todo list do an initial fetch
  htmx.trigger("#todo-list", "triggerLoad");

  // get all the elements!
  const userIdContainer = document.querySelector("#userIdContainer");
  const userIdEditForm = document.querySelector("#userIdEditForm");
  const userIdDisplay = document.querySelector("#userIdDisplay");
  const userIdEditField = document.querySelector("#userIdEditField");
  const userIdEditButton = document.querySelector("#userIdEditButton");
  const userIdCancelButton = document.querySelector("#userIdCancelButton");

  const currentUserId = getOrCreateUserId();

  // helpers
  function startEdit() {
    userIdEditForm.classList.add("shown");
    userIdEditForm.classList.remove("hidden");
    userIdContainer.classList.remove("shown");
    userIdContainer.classList.add("hidden");
  }

  function endOrCancelEdit() {
    userIdContainer.classList.add("shown");
    userIdContainer.classList.remove("hidden");
    userIdEditForm.classList.remove("shown");
    userIdEditForm.classList.add("hidden");
  }

  // only really makes sense if all elements are present
  if (
    userIdContainer &&
    userIdDisplay &&
    userIdEditField &&
    userIdEditForm &&
    userIdEditButton &&
    userIdCancelButton
  ) {
    userIdDisplay.innerHTML = currentUserId;
    userIdEditField.value = currentUserId;

    // save new user id
    userIdEditForm.addEventListener("submit", (event) => {
      event.preventDefault();
      const newUserId = userIdEditField.value;

      if (newUserId && newUserId !== currentUserId) {
        localStorage.setItem("userId", userIdEditField.value);
        userIdDisplay.innerHTML = getOrCreateUserId();

        // Run HTMX request and reload todo list
        htmx.ajax(
          "GET",
          `/api/changeUserId?old=${currentUserId}&new=${newUserId}`,
          "#todo-list",
        );
      }

      // hide form again
      endOrCancelEdit();
    });

    // show form and hide display
    userIdEditButton.addEventListener("click", () => {
      startEdit();

      // wait for the field to be visible before trying to focus
      setTimeout(() => {
        userIdEditField.select();
      }, 0);
    });

    // Esc
    userIdEditField.addEventListener("keydown", (event) => {
      if (event.key === "Escape") {
        endOrCancelEdit();
      }
    });

    // cancel button
    userIdCancelButton.addEventListener("click", () => {
      endOrCancelEdit();
    });
  }
});

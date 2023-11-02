const todoEditForm = document.querySelector(
  "#form_{{- .UserId -}}_{{- .Id -}}",
);
const todoInputField = document.querySelector(
  "#input_{{- .UserId -}}_{{- .Id -}}",
);
const todoSaveButton = document.querySelector(
  "#save_{{- .UserId -}}_{{- .Id -}}",
);
const todoCancelButton = document.querySelector(
  "#cancel_{{- .UserId -}}_{{- .Id -}}",
);

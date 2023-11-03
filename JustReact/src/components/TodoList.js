import { useEffect, useState } from "react";
import TodoListFilter from "./TodoListFilter";
import TodoListDisplay from "./TodoListDisplay";
import AddTodo from "./AddTodo";

const TodoList = () => {
  // variables
  const storageKey = "todos";

  // create display list based on current filter
  const createDisplayList = () => {
    let displayList = todoList;

    if (currentFilter === "open") displayList = openTodoList;
    if (currentFilter === "completed") displayList = completedTodoList;

    return displayList;
  };

  // States
  const [todoList, setTodoList] = useState(null);
  const [currentFilter, setCurrentFilter] = useState("all");

  // Derived from states
  const openTodoList = todoList?.filter((item) => item.done === false) ?? [];
  const completedTodoList = todoList?.filter((item) => item.done === true) ?? [];
  const displayList = createDisplayList();

  // load to-do list from local storage
  useEffect(() => {
    let storedTodos = JSON.parse(localStorage.getItem(storageKey)) ?? [];
    console.log("Initial Load", storedTodos);
    setTodoList(storedTodos);
  }, []);

  // generate a new todo list item
  const createTodo = (newText) => {
    return {
      id: crypto.randomUUID(),
      text: newText,
      done: false,
    };
  };

  // persist todolist when state is updated
  useEffect(() => {
    // prevent initialn load from overwriting an existing list...
    if (!todoList) return;
    localStorage.setItem(storageKey, JSON.stringify(todoList));
  }, [todoList]);

  // find a todo item by id
  const getTodoById = (id) => {
    return todoList.find((entry) => entry.id === id);
  };

  // Add Handlers
  const addItemHandler = (newText) => {
    const newTodo = createTodo(newText);
    const newTodoList = todoList.concat([newTodo]);
    setCurrentFilter("all");

    setTodoList(newTodoList);
    localStorage.setItem(storageKey, JSON.stringify(newTodoList));
  };

  // Filter Handlers
  const filterHandler = (filterType) => {
    setCurrentFilter(filterType);
  };

  // Item Handlers
  // mark to-do item as done
  const doneHandler = (id, isDone) => {
    const todoEntry = getTodoById(id);

    todoEntry.done = isDone;
    setTodoList([...todoList]);
  };

  // edit to-do item
  const editHandler = (id, newText) => {
    const todoEntry = getTodoById(id);
    todoEntry.text = newText;
    setTodoList([...todoList]);
  };

  // delete to-do item
  const deleteHandler = (id) => {
    const todoEntryIndex = todoList.findIndex((entry) => entry.id === id);

    if (todoEntryIndex !== -1) {
      todoList.splice(todoEntryIndex, 1);
      setTodoList([...todoList]);
    }
  };

  return (
    <>
      <AddTodo addItemHandler={addItemHandler} />
      {todoList && todoList.length > 0 && (
        <TodoListFilter
          filterHandler={filterHandler}
          currentFilter={currentFilter}
          allTodos={todoList.length}
          openTodos={openTodoList.length}
          completedTodos={completedTodoList.length}
        />
      )}
      <TodoListDisplay
        todoList={displayList}
        doneHandler={doneHandler}
        editHandler={editHandler}
        deleteHandler={deleteHandler}
        currentFilter={currentFilter}
        filterHandler={filterHandler}
      />
    </>
  );
};

export default TodoList;

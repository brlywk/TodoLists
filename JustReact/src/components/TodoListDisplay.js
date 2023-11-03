import TodoItem from "./TodoItem";

const TodoListDisplay = ({ todoList, doneHandler, editHandler, deleteHandler, currentFilter, filterHandler }) => {
  // RENDER
  return (
    <div className="flex flex-col gap-1 md:gap-4 w-full md:w-4/5 lg:3/4 xl:w-3/6">
      {todoList &&
        todoList.map((t) => (
          <TodoItem
            key={t.id}
            id={t.id}
            text={t.text}
            done={t.done}
            doneHandler={doneHandler}
            editHandler={editHandler}
            deleteHandler={deleteHandler}
          />
        ))}
      {(!todoList || todoList.length === 0) && (
        <div className="text-center my-8 font-bold">
          {currentFilter === "all" && "You have not added any todos yet. Use the input field above to get started!"}
          {currentFilter !== "all" && (
            <div className="flex flex-col items-center justify-center">
              <p>You currently don't have any todos in this category.</p>
              <p className="mt-4">
                <button
                  className="cursor-pointer font-normal rounded-2xl border border-black/20 border-solid hover:bg-black/20 py-2 px-4  transition duration-300"
                  onClick={() => filterHandler("all")}
                >
                  Show all todos
                </button>
              </p>
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default TodoListDisplay;

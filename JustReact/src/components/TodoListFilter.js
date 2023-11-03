const TodoListFilter = ({ filterHandler, currentFilter, allTodos, openTodos, completedTodos }) => {
  // Links we use in filter (and b/c we're lazy and don't wont to write tailwind classes multiple times)
  const filterLinks = [
    {
      text: "All",
      filter: "all",
    },
    {
      text: "Open",
      filter: "open",
    },
    {
      text: "Completed",
      filter: "completed",
    },
  ];

  // helpers
  const getTodosForFilter = (filter) => {
    if (filter === "all") return allTodos;
    if (filter === "open") return openTodos;
    if (filter === "completed") return completedTodos;
  };

  return (
    <div className="w-2/3 mx-auto text-center flex flex-row gap-4 justify-center my-10">
      {filterLinks.map((link) => (
        <button
          className={
            "border border-black/20 border-solid hover:bg-black/20 px-4 py-2 rounded-2xl transition duration-300 cursor-pointer flex flex-row items-center gap-2" +
            (currentFilter === link.filter ? " bg-black/10" : "")
          }
          onClick={() => filterHandler(link.filter)}
          key={link.text}
        >
          {link.text}{" "}
          {getTodosForFilter(link.filter) > 0 && (
            <span className="text-center bg-black/20 rounded-full py-1 px-2 text-[0.5rem] ">
              {getTodosForFilter(link.filter)}
            </span>
          )}
        </button>
      ))}
    </div>
  );
};

export default TodoListFilter;

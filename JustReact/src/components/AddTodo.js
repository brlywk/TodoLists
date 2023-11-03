import { useRef } from "react";

const AddTodo = ({ addItemHandler }) => {
  // Refs
  const inputRef = useRef(null);

  // Handlers
  const addItemInputHandler = (event) => {
    event.preventDefault();
    const newText = inputRef.current.value;

    if (newText) addItemHandler(newText);

    inputRef.current.value = "";
  };

  return (
    <div className="w-full md:w-4/5 lg:3/4 xl:w-3/6 bg-black/5 p-2 border border-black/10 rounded-2xl shadow-xl backdrop-blur-lg mb-4 md:mb-8">
      <form
        className="flex flex-row flex-nowrap gap-1 md:gap-2 justify-between relative items-center"
        onSubmit={addItemInputHandler}
      >
        <input
          type="text"
          placeholder="Start typing here..."
          className="w-full border-0 bg-transparent text-xl placeholder:italic placeholder:text-black/20 focus:ring-0"
          ref={inputRef}
          autoFocus
        />
        <button
          type="submit"
          className="border border-black/20 border-solid hover:bg-black/20 rounded-2xl py-2 px-6 transition duration-300 flex flex-row flex-nowrap justify-around align-middle gap-2"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24">
            <path
              fill="currentColor"
              d="M11 17h2v-4h4v-2h-4V7h-2v4H7v2h4v4Zm1 5q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"
            />
          </svg>
          Add
        </button>
      </form>
    </div>
  );
};

export default AddTodo;

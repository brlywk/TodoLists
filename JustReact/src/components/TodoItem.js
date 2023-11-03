import { useState } from "react";
import IconButton from "./IconButton";
import YesNoButtons from "./YesNoButtons";

const TodoItem = ({ id, text, done, doneHandler, editHandler, deleteHandler }) => {
  // Variables
  const editPath =
    "M5 19h1.4l8.625-8.625l-1.4-1.4L5 17.6V19ZM19.3 8.925l-4.25-4.2l1.4-1.4q.575-.575 1.413-.575t1.412.575l1.4 1.4q.575.575.6 1.388t-.55 1.387L19.3 8.925ZM4 21q-.425 0-.713-.288T3 20v-2.825q0-.2.075-.388t.225-.337l10.3-10.3l4.25 4.25l-10.3 10.3q-.15.15-.337.225T6.825 21H4ZM14.325 9.675l-.7-.7l1.4 1.4l-.7-.7Z";
  const deletePath =
    "M7 21q-.825 0-1.413-.588T5 19V6q-.425 0-.713-.288T4 5q0-.425.288-.713T5 4h4q0-.425.288-.713T10 3h4q.425 0 .713.288T15 4h4q.425 0 .713.288T20 5q0 .425-.288.713T19 6v13q0 .825-.588 1.413T17 21H7ZM17 6H7v13h10V6ZM9 17h2V8H9v9Zm4 0h2V8h-2v9ZM7 6v13V6Z";
  const checkMarkPath =
    "m10 13.6l5.9-5.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-6.6 6.6q-.3.3-.7.3t-.7-.3l-2.6-2.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9Z";
  const abortPath =
    "M18.3 5.71a.996.996 0 0 0-1.41 0L12 10.59L7.11 5.7A.996.996 0 1 0 5.7 7.11L10.59 12L5.7 16.89a.996.996 0 1 0 1.41 1.41L12 13.41l4.89 4.89a.996.996 0 1 0 1.41-1.41L13.41 12l4.89-4.89c.38-.38.38-1.02 0-1.4z";
  const yesClasses = [
    "bg-green-400/20",
    "rounded",
    "hover:bg-green-400",
    "hover:fill-white",
    "transition",
    "duration-300",
  ];
  const noClasses = ["bg-red-400/20", "rounded", "hover:bg-red-400", "hover:fill-white", "transition", "duration-300"];

  // States
  const [editMode, setEditMode] = useState(false);
  const [deleteMode, setDeleteMode] = useState(false);
  const [isDone, setIsDone] = useState(done);
  const [editFieldValue, setEditFieldValue] = useState(text);

  // Handlers
  // handle editing of item
  const editFinished = (event) => {
    if (event.key === "Escape") {
      setEditMode(false);
    }

    if (event.key === "Enter" || event.type === "click") {
      setEditMode(false);
      if (editFieldValue !== "") editHandler(id, editFieldValue);
    }
  };

  // deletion was confirmed
  const deleteConfirmed = () => {
    deleteHandler(id);
  };

  // handle marking an item as done
  const markAsDone = () => {
    setIsDone(!isDone);
    doneHandler(id, !isDone);
  };

  const handleEditFieldChange = (event) => {
    if (!event.currentTarget.value) return;
    setEditFieldValue(event.currentTarget.value);
  };

  // RENDER
  return (
    <div className="bg-white/20 p-4 border border-black/10 rounded-2xl shadow-xl backdrop-blur-lg hover:border-black/50">
      <div className="flex flex-row flex-nowrap gap-1 md:gap-2 justify-between relative items-center">
        {/* Check */}
        <input
          type="checkbox"
          className="cursor-pointer bg-transparent rounded-full text-green-500 focus:ring-0 focus:ring-transparent focus:ring-offset-transparent hover:border-2 hover:border-green-500"
          checked={isDone}
          onChange={markAsDone}
        />

        {/* Todo Text */}
        {!editMode && (
          <span
            className={"w-full cursor-pointer" + (isDone ? " line-through text-black/50" : "")}
            onDoubleClick={() => setEditMode(true)}
            onClick={markAsDone}
          >
            {text}
          </span>
        )}

        {/*Edit Input Field*/}
        {editMode && (
          <input
            type="text"
            className="w-full border-0 bg-transparent p-0 focus:ring-0 rounded outline outline-1 outline-black/20 outline-offset-2 focus:outline-1 focus:outline-offset-2 focus:outline-black/50"
            defaultValue={editFieldValue}
            onChange={handleEditFieldChange}
            onFocus={(event) => event.currentTarget.select()}
            onKeyDown={editFinished}
            autoFocus
          />
        )}

        {/* Edit */}
        {!editMode && !deleteMode && (
          <IconButton
            svgPath={editPath}
            classes={["fill-black/50", "hover:fill-black"]}
            clickHandler={() => setEditMode(true)}
          />
        )}

        {/* Edit Mode Yes / No Buttons */}
        {editMode && (
          <YesNoButtons
            yesHandler={editFinished}
            yesSvgPath={checkMarkPath}
            yesClasses={yesClasses}
            noHandler={() => setEditMode(false)}
            noSvgPath={abortPath}
            noClasses={noClasses}
          />
        )}

        {/* Delete */}
        {!editMode && !deleteMode && (
          <IconButton
            svgPath={deletePath}
            classes={["fill-black/50", "hover:fill-red-500"]}
            clickHandler={() => setDeleteMode(true)}
          />
        )}

        {/* Delete Mode Yes / No Buttons */}
        {deleteMode && (
          <YesNoButtons
            yesHandler={deleteConfirmed}
            yesSvgPath={checkMarkPath}
            yesClasses={yesClasses}
            noHandler={() => setDeleteMode(false)}
            noSvgPath={abortPath}
            noClasses={noClasses}
          />
        )}
      </div>
    </div>
  );
};

export default TodoItem;

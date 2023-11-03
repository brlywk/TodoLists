import { useEffect, useState } from "react";

const ScrollToTop = ({ scrollDistance }) => {
  // Variables
  const svgPath =
    "M11 11.8V15q0 .425.288.713T12 16q.425 0 .713-.288T13 15v-3.2l.9.9q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-2.6-2.6q-.3-.3-.7-.3t-.7.3l-2.6 2.6q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l.9-.9ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z";
  const defaultClasses = [
    "fixed",
    "w-full",
    "max-md:left-0",
    "bottom-0",
    "md:bottom-4",
    "md:right-4",
    "md:w-auto",
    "transition",
    "duration-300",
    "ease-in-out",
  ];

  // State
  const [isVisibile, setIsVisible] = useState(false);

  // add event listener to scroll event
  useEffect(() => {
    window.addEventListener("scroll", () => {
      if (window.scrollY > scrollDistance) {
        setIsVisible(true);
      } else {
        setIsVisible(false);
      }
    });
  }, [scrollDistance]);

  // click handler
  const clickHandler = () => {
    window.scrollTo({
      behavior: "smooth",
      top: 0,
    });
  };

  return (
    <div className={defaultClasses.join(" ") + " " + (isVisibile ? "opacity-100" : "hidden opacity-0")}>
      <button
        className="group w-full bg-black/50 p-2 flex flex-row items-center justify-center md:bg-transparent md:w-auto md:border md:border-black/20 md:rounded-2xl max-md:hover:bg-black/75 md:hover:border-black/20 md:hover:bg-black/20 transition duration-300"
        onClick={clickHandler}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          className="w-10 h-10 lg:w-12 lg:h-12 stroke-white fill-transparent md:stroke-none md:fill-black/50 md:group-hover:fill-black"
        >
          <path d={svgPath} />
        </svg>
      </button>
    </div>
  );
};

export default ScrollToTop;

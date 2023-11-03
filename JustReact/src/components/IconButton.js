const IconButton = ({ svgPath, clickHandler, classes }) => {
  return (
    <button onClick={clickHandler}>
      <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" className={classes.join(" ")}>
        <path d={svgPath} />
      </svg>
    </button>
  );
};

export default IconButton;

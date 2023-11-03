import IconButton from "./IconButton";

const YesNoButtons = ({ yesHandler, noHandler, yesSvgPath, noSvgPath, yesClasses, noClasses }) => {
  return (
    <>
      <IconButton svgPath={yesSvgPath} classes={yesClasses} clickHandler={yesHandler} />
      <IconButton svgPath={noSvgPath} classes={noClasses} clickHandler={() => noHandler(false)} />
    </>
  );
};

export default YesNoButtons;

const Footer = () => {
  // Emotions while coding... :P
  const madeWithEmotes = ["♥️", "♥️", "♥️", "♥️", "😭", "🤬", "🤯", "😵‍💫"];

  // get random emote
  const getRandomEmote = () => {
    const rnd = Math.floor(Math.random() * madeWithEmotes.length);
    return madeWithEmotes[rnd];
  };

  return (
    <footer className="max-md:mb-24">
      <div className="text-center text-sm text-black/50">
        Made with {getRandomEmote()} and a little bit of help from{" "}
        <a href="https://react.dev/" target="_blank" rel="noreferrer" className="text-black hover:underline">
          React
        </a>{" "}
        and{" "}
        <a href="https://tailwindcss.com/" target="_blank" rel="noreferrer" className="text-black hover:underline">
          Tailwind
        </a>
      </div>
    </footer>
  );
};

export default Footer;

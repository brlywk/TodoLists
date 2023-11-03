import Header from "./components/Header";
import Main from "./components/Main";
import Footer from "./components/Footer";
import ScrollToTop from "./components/ScrollToTop";

function App() {
  // some trickery to randomize background gradient
  const gradients = [
    ["from-orange-100", "to-rose-100"],
    ["from-gray-50", "to-gray-300"],
    ["from-red-50", "to-amber-50"],
    ["from-cyan-100", "to-pink-300"],
  ];
  const randomGradient = gradients[Math.floor(Math.random() * gradients.length)];
  document.documentElement.classList.add(...randomGradient);

  return (
    <div className="w-full my-4 px-2 md:my-8 justify-center">
      <Header />
      <Main />
      <Footer />
      <ScrollToTop scrollDistance={200} />
    </div>
  );
}

export default App;

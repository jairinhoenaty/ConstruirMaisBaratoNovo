import { useState } from "react";
import Navigation from "./Navigation";

function Header() {
  const [currentPage, setCurrentPage] = useState(() => {
    const isLoggedIn = localStorage.getItem("isLoggedIn") === "true";
    return isLoggedIn ? "professional-panel" : "home";
  });
  const [carouselPage, setCarouselPage] = useState("H");

  return (
    <div className="h-20">
      <Navigation
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        carouselPage={"H"}
        setCarouselPage={setCarouselPage}
      />
    </div>
  );
}
export default Header;

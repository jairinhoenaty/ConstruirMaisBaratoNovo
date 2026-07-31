import { useLocation } from "react-router-dom";
import CarouselNovo from "./CarouselNovo";
import SearchResultsBanner from "./SearchResultsBanner";

function BannerHeader() {
  const location = useLocation();

  if (location.pathname === "/dashboard" || location.pathname === "/checkout") {
    return <div></div>;
  }

  return (
    <div>
      {" "}
      {location.pathname === "/search-results" ? (
        <SearchResultsBanner />
      ) : (
        <CarouselNovo page={location.pathname as string} />
      )}
    </div>
  );
}

export default BannerHeader;

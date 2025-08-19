import { useLocation } from "react-router-dom";
import CarouselNovo from "./CarouselNovo";

function BannerHeader() {
  const location = useLocation();

  return (
    <div>
      {" "}
      {location.pathname != "/dashboard" && (
        <CarouselNovo  page={location.pathname as string} />  
      )}
    </div>
  );
}

export default BannerHeader;

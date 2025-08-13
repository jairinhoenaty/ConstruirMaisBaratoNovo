import { useState, useEffect } from "react";
import { ChevronLeft, ChevronRight } from "lucide-react";
import { BannerService } from "../services/BannerService";
import { useLocation } from "react-router-dom";

function Carousel(page: any) {
  const [currentSlide, setCurrentSlide] = useState(0);
  const [images, setImages] = useState([{ image: "", title: "", link: "" }]);
  const nextSlide = () => {
    setCurrentSlide((prev) => (prev + 1) % images.length);
  };

  const prevSlide = () => {
    setCurrentSlide((prev) => (prev - 1 + images.length) % images.length);
  };

  const location = useLocation();

  useEffect(() => {
    let cityId = 0;
    const fetchData = async () => {
      if (page.page == "/search") {
        page = "S";
      } else if (page.page == "/marketplace") {
        page = "M";
      } else if (page.page == "/privacy") {
        page = "P";
      } else if (page.page == "/login") {
        page = "L";
      } else if (page.page == "/register") {
        page = "R";
      } else if (page.page == "/dashboard") {
        page = "D";
      } else if (page.page == "/professional-panel") {
        page = "A";
      } else if (page.page == "/search-results") {
        page = "C";
        cityId = parseInt(location.state.selectedCity);
      } else {
        page = "H";
      }
      const data = { page: page, cityId: cityId, regionId: 0 };
      const response = await BannerService.getBannerByPagePublic(data);
      if ((response.status == 200)) {
        const json = await response.data;
        if (json.length === 0) {
          const response_home = await BannerService.getBannerByPagePublic({
            page: "H",
            cityId: 0,
            regionId:0,
          });
          if ((response_home.status == 200)) {
            const json_home = await response_home.data;
            setImages(json_home);
          }
        } else {
          setImages(json);
        }
      }
    };

    fetchData();
    const timer = setInterval(nextSlide, 3000);
    return () => clearInterval(timer);
  }, [page]);

  return (
    <div className="relative h-[400px] overflow-hidden">
      {images.map((image, index) => (
        <div
          key={index}
          className={`absolute w-full h-full transition-transform duration-500 ease-in-out`}
          style={{
            transform: `translateX(${(index - currentSlide) * 100}%)`,
          }}
        >
          {index}
          <img
            src={"data:image;base64," + image.image}
            alt={image.title}
            onClick={
              image.link != ""
                ? () => window.open(image.link, "_blank")
                : () => {}
            }
            className="w-full h-full object-cover"
            style={{ cursor: image.link != "" ? "pointer" : "" }}
          />
          <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/70 to-transparent p-8"></div>
        </div>
      ))}
      <button
        onClick={prevSlide}
        className="absolute left-4 top-1/2 -translate-y-1/2 bg-white/80 p-2 rounded-full hover:bg-white transition-colors"
      >
        <ChevronLeft className="w-6 h-6" />
      </button>
      <button
        onClick={nextSlide}
        className="absolute right-4 top-1/2 -translate-y-1/2 bg-white/80 p-2 rounded-full hover:bg-white transition-colors"
      >
        <ChevronRight className="w-6 h-6" />
      </button>
    </div>
  );
}

export default Carousel;

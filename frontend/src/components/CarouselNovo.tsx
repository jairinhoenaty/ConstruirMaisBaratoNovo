import { useState, useEffect } from "react";
import { BannerService } from "../services/BannerService";
import { useLocation } from "react-router-dom";
import Slider from "react-slick";
import { RegionService } from "../services";

function CarouselNovo(page: any) {
  const [images, setImages] = useState([]);
  const location = useLocation();
    const [isLoading, setIsLoading] = useState(false);
  //const [selectedRegion,setSelectedRegion] = useState(0);
  let region = 0;

  useEffect(() => {
    
    let cityId = 0;
    
    const fetchData = async () => {
      setIsLoading(true);
      region = 0;
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
        const cityIdSearch = parseInt(location.state.selectedCity);
        try {
          const resu_region = await RegionService.getRegionbyCity(cityIdSearch);
          if (resu_region.status == 200) {
            region = await resu_region.data.id;
            //setSelectedRegion(parseInt(resu_region.data.id));
          }
        } catch (error) {
          console.error("Erro ao obter região do usuário:", error);
        }

      } else {
        page = "H";
      }

      let json_images=[];
      let images_region = [];
      const data = { page: page, cityId: cityId,regionId: 0 };
      const response = await BannerService.getBannerByPagePublic(data);
      if (response.status == 200) {
        const json = await response.data;
        if (json.length === 0) {
          const response_home = await BannerService.getBannerByPagePublic({
            page: "H",
            cityId: 0,
            regionId: 0,
          });
          if (response_home.status == 200) {
            const json_home = await response_home.data;
            json_images = json_home;
            //setImages(json_home);
            //setIsLoading(false);
          }
        } else {
          json_images= json;
          //setImages(json);
        }

        if ((localStorage.getItem("isLoggedIn") === "true") &&  page !== "C") {
          const cityId_logged = parseInt(localStorage.getItem("city_id")??""); 
          try {
            const resu_region = await RegionService.getRegionbyCity(cityId_logged);
            if (resu_region.status==200){
              region = await resu_region.data.id;
              //setSelectedRegion(parseInt(resu_region.data.id));
            }
          } catch (error) {
            console.error("Erro ao obter região do usuário:", error);   
          }

        }

        if (region !== 0)  {
          const response_home = await BannerService.getBannerByPagePublic({
            page: "-",
            cityId: 0,
            regionId: region,
          });
          if (response_home.status == 200) {
            const json_route = await response_home.data;
            images_region = json_route;
            //setImages(mergedObject);
          }
        }
  
      }
      let merge_json=[];
      if (page == "C" && images_region.length > 0) {
        merge_json = images_region;
      } else {
        if ((localStorage.getItem("isLoggedIn") === "true") &&  page !== "C" && images_region.length > 0 ) {
          merge_json = images_region;
        }
        else {
          merge_json = json_images;
        }
        
      }
      console.log("merge_json", merge_json);
      
      setImages(merge_json);
      setIsLoading(false);
    };

    fetchData();
  }, [page.page]);

  const settings = {
    dots: true,
    //infinite: true,
    speed: 500,
    slidesToShow: 1,
    slidesToScroll: 1,
    initialSlide: 1,
    autoplay: true,
    autoplaySpeed: 4000,
    
//    prevArrow: '<button type="button" class="slick-custom-arrow slick-prev"> < </button>',
//    nextArrow: '<button type="button" class="slick-custom-arrow slick-next"> > </button>',
    responsive: [
      {
        breakpoint: 1024,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1,
          infinite: true,
          dots: true,
        },
      },
      {
        breakpoint: 600,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1,
          initialSlide: 1,
        },
      },
      {
        breakpoint: 480,
        settings: {
          slidesToShow: 1,
          slidesToScroll: 1,
        },
      },
    ],
  };
  return (
    <div className="slider-container max-w-[1400px] mx-auto px-4">
      <link
        rel="stylesheet"
        type="text/css"
        //charset="UTF-8"
        href="https://cdnjs.cloudflare.com/ajax/libs/slick-carousel/1.6.0/slick.min.css"
      />
      <link
        rel="stylesheet"
        type="text/css"
        href="https://cdnjs.cloudflare.com/ajax/libs/slick-carousel/1.6.0/slick-theme.min.css"
      />

      <Slider {...settings}>
        {isLoading && (
          <div className="flex justify-center items-center h-[400px]">
            <div className="loader"></div>
          </div>
        )}

        {images.map((image, index) => (
          <div>
            <img
              //src={"data:image;base64," + image.image}
              src={atob(image.image)}
              alt={image.title}
              onClick={
                image.link != ""
                  ? () => window.open(image.link, "_blank")
                  : () => {}
              }
              //className="w-full h-full object-cover"
              //              className="w-full h-[300px] md:h-[400px] xl:h-[500px] object-cover rounded-lg"
              className="w-full h-[400px] object-cover rounded-lg shadow-lg transition-transform duration-300 hover:scale-105"
              style={{
                cursor: image.link != "" ? "pointer" : "",
                height: "400px",
              }}
            />
            <div className="absolute bottom-0 left-0 right-0 bg-gradient-to-t from-black/10 to-transparent p-8">
              {index}
              {images.length}
            </div>
          </div>
        ))}
      </Slider>
    </div>
  );
}

export default CarouselNovo;

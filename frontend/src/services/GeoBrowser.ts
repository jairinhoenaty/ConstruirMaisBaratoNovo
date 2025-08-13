import React, { useState, useEffect } from "react";

const GeoBrowserService = () => {
  const [location, setLocation] = useState(null);
  const [city, setCity] = useState("");
  const [state, setState] = useState("");

  useEffect(() => {
    // Solicitar permissão para acessar a localização
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const { latitude, longitude } = position.coords;
          setLocation({ latitude, longitude });
          // Chamada para a API de geocodificação (exemplo com Google Maps - substitua pela sua API)
          // Substitua a chave de API pela sua
          fetch(
            `https://maps.googleapis.com/maps/api/geocode/json?latlng=${latitude},${longitude}&key=SUA_CHAVE_DE_API`
          )
            .then((response) => response.json())
            .then((data) => {
              // Analisar a resposta para obter cidade e estado
              const addressComponents = data.results[0].address_components;
              let foundCity = null;
              let foundState = null;

              for (const component of addressComponents) {
                if (component.types.includes("locality")) {
                  foundCity = component.long_name;
                } else if (component.types.includes("administrative_level_1")) {
                  foundState = component.long_name;
                }
              }
              setCity(foundCity);
              setState(foundState);
            })
            .catch((error) => console.error("Erro na chamada à API:", error));
        },
        (error) => {
          console.error("Erro ao obter a localização:", error);
        }
      );
    } else {
      console.error("Geolocalização não suportada pelo browser");
    }
  }, []);

  return (
    <div>
      {location && (
        <div>
          <p>Latitude: {location.latitude}</p>
          <p>Longitude: {location.longitude}</p>
          <p>Cidade: {city}</p>
          <p>Estado: {state}</p>
        </div>
      )}
    </div>
  );
};

export default App;

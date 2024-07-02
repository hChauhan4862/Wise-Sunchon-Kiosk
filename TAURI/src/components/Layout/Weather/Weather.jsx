import React from "react";

const WeatherComponent = ({ weatherCondition }) => {
  const weatherIcons = {
    sunny: "icon-sunny.png",
    cloudy: "icon-cloudy.png",
    rainy: "icon-rainy.png",
    snowy: "icon-snowy.png",
    thunderstorm: "icon-thunderstorm.png",
    foggy: "icon-foggy.png",
    clear: "icon-clear.png",
    partlyCloudy: "icon-partly-cloudy.png",
    windy: "icon-windy.png",
    mist: "icon-mist.png",
    hail: "icon-hail.png",
    drizzle: "icon-drizzle.png",
    tornado: "icon-tornado.png",
    sandstorm: "icon-sandstorm.png",
    blizzard: "icon-blizzard.png",
    smoky: "icon-smoky.png",
    freezingRain: "icon-freezing-rain.png",
    overcast: "icon-overcast.png",
    hailstorm: "icon-hailstorm.png",
    rainbow: "icon-rainbow.png",
  };

  const getWeatherIcon = (condition) => {
    const iconFilename = weatherIcons[condition];
    return iconFilename ? `images/${iconFilename}` : null;
  };

  const weatherIcon = getWeatherIcon(weatherCondition);
  console.log("weather condition", weatherCondition);
  console.log("weather iconsnnss", weatherIcon);
  return (
    <div className="weather-icon">
      {weatherIcon && <img src={weatherIcon} alt="Weather Icon" />}
    </div>
  );
};

export default WeatherComponent;

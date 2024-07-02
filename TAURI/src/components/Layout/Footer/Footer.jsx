import React, { useState, useEffect } from "react";

import { useLocation } from "react-router-dom";
// import TimerWithDate from "./timeWithDate";
import LanguageComponent from "../Language/Language";
import axios from "axios";
import WeatherComponent from "../Weather/Weather";
import { timeInterval, TimerWithDate } from "./timeWithDate";
import { useTranslation } from "react-i18next";

const Footer = () => {
  const pathName = useLocation().pathname;
  const [time, setTime] = useState(TimerWithDate());
  const [weatherData, setWeatherData] = useState([]);
  const [city, setCity] = useState("Suncheon");
  const [country, setCountry] = useState("");
  const [weatherCondition, setWeatherCondition] = useState([]);
  const [temperature, settemperature] = useState();

  const API_KEY = "3e2d332a74ccf9786d820679a31cc8ce";
  function getWeatherData() {
    axios
      .get(
        `https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${API_KEY}`
      )
      .then((res) => {
        setWeatherData(res.data);
        setWeatherCondition(res.data.weather[0].description);
        setCity(res.data.name);
        setCountry(res.data.sys.country);
        settemperature(res.data.main.feels_like);
      })
      .catch((err) => {
        console.log(err);
      });
  }
  useEffect(() => {
    getWeatherData()
    setInterval(() => {
      getWeatherData()
    }, 1000 * 60 * 3)
  }, []);
  const formattedTemperature = (temperature - 273).toFixed(1);

  useEffect(() => {
    const interval = setInterval(() => {
      timeInterval()
    }, 1000)
    return () => clearInterval(interval)
  }, []);
  const { t } = useTranslation();
  return (
    <div id="footer-main">
      <div id={pathName === "/" ? "footer-widget" : "footer-widget2"}>
        {/* --------------------start  language component in footer------------------------------------------- */}
        <LanguageComponent />
        {/* --------------------end  language component in footer------------------------------------------- */}
        <div className="weather-widget">
          {/* -----------------------------------------start weather status-------------------------------- */}
          <div className="weather-container">
            <div
              className={
                pathName === "/" ? "weather-detail" : "weather-detail2"
              }
            >
              <h6>
                {formattedTemperature}°C <span> {weatherCondition} </span>
              </h6>
              <p>
                {" "}
                {city}, {country}
              </p>
            </div>
            <div className="weather-icon">
              <img
                src="/images/icon-rainy-day.png"
                alt="Rainy Cloud Icon"
              ></img>
            </div>

            {/* <WeatherComponent weatherCondition={weatherCondition} /> */}
          </div>
          {/* -----------------------------------------end weather status-------------------------------- */}
          <div className="date-time-container">
            {/* ----------------------------------------start date time---------------------------------- */}
            <div className={pathName === "/" ? "date-detail" : "date-detail2"}>
              <h6> {t(time.day)}</h6>
              <p>
                {" "}
                {time.date} {t(time.month)}, {time.year}
              </p>
            </div>
            <div className={pathName === "/" ? "time-detail" : "time-detail2"}>
              <h6 id="hm_ampm"></h6>
            </div>
            {/* ----------------------------------------end date time---------------------------------- */}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Footer;

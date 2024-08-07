import i18n from "i18next";
import { initReactI18next } from "react-i18next";
import LanguageDetector from "i18next-browser-languagedetector";
import HttpApi from "i18next-http-backend";

const languageCode = localStorage.getItem('languageCode')
const options = {
  order: ["cookie", "htmlTag", "localStorage", "path", "subdomain"],
  caches: ["cookie"],
};
i18n
  .use(initReactI18next) // passes i18n down to react-i18next
  .use(LanguageDetector)
  .use(HttpApi)
  .init({
    fallbackLng: languageCode,
    lng:  languageCode, // default language,
    supportedLngs: ["en", "zh", "ko", "th", "ja", "ar", "fr", "es"],
    detection: options,
    backend: {
      loadPath: "/locales/{{lng}}/translation.json",
    },
  });

export default i18n;

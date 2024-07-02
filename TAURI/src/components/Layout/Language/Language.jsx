import React, { useState } from "react";
import { useTranslation } from "react-i18next";
import { Link } from "react-router-dom";
import i18n from "./i18n_react";
import { languages } from "./languageList";
import { useDispatch } from "react-redux";
import { setSelectedLanguage } from "../../../redux/actions/setup";

const LanguageComponent = () => {
  const dispatch = useDispatch()
  const [activeLanguage, setActiveLanguage] = useState("USA");
  const { t } = useTranslation();

  const handleLanguageClick = (name, code) => {
    setActiveLanguage(name);
    i18n.changeLanguage(code);
    const languageName = name === 'KOREAN' ? 'KOREAN' : 'ENGLISH'
    dispatch(setSelectedLanguage(languageName))
  };

  return (
    <div className="language-widget">
      <h4>{t("select Language")} <img src="/images/arrow-right.png" width={20}/> </h4>
      <ul>
        {languages &&
          languages.map((language) => (
            <li key={language.id}>
              <Link
                to=""
                onClick={(e) => {
                  e.preventDefault();
                  handleLanguageClick(language.name, language.code);
                }}
                className={activeLanguage === language.name ? "active" : ""}
              >
                <img src={language.img} alt={`Flag ${language.name}`} />
              </Link>
            </li>
          ))}
      </ul>
    </div>
  );
};

export default LanguageComponent;

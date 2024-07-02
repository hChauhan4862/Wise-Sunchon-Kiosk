import React, { useState, useEffect } from "react";
import Header from "./Header/Header";
import Footer from "./Footer/Footer";
import { useLocation } from "react-router-dom";
import { SessionTimer } from "../SessionTimer/sessionLogoutTimer";

const Layout = ({ children }) => {
  const path = useLocation().pathname;
  const [pathName, setPathName] = useState(path);

  useEffect(() => {
    setPathName(path);
  }, [path]);

  return (
    <>
      <div
        onTouchStart={() => SessionTimer()}
        // onClick={() => SessionTimer()}
        style={
          pathName !== "/"
            ? {
              backgroundImage: `url("/images/bg-main.jpg")`,
              width: "1920px",
              height: "1080px",
            }
            : null
        }
      >
        {pathName === "/" ? (
          <video id="background-video" autoPlay loop muted>
            <source src="/video/background-02.mp4" type="video/mp4" />
          </video>
        ) : null}
        <Header />
        {children}
        <Footer />
      </div>
    </>
  );
};

export default Layout;
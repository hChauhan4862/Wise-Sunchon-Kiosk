import "./App.css";
import Home from "./pages/Home/Home";
import { BrowserRouter } from "react-router-dom";
import Layout from "./components/Layout";
import { useDispatch, useSelector } from "react-redux";
import Router from "./Routes";
import { useEffect } from "react";
import { setTimers, hcLiveSyncFun } from "./redux/actions/setup";
import { actionSeatsMapLoad } from "./redux/actions/seats";
import { ToastContainer, toast } from "react-toastify";
import withReactContent from "sweetalert2-react-content";
import "react-toastify/dist/ReactToastify.css";

import { invoke } from "@tauri-apps/api";
import { changeAPIBaseUrl } from "./apis";

const timers = {
  popup: 20, // 20
  worning: 1,
  success: 1,
  error: 1,
  print: 10,  // 10
  floorPage: 90, // 16
  bookingsDetail: 60,
  keyboard: 15,
};

// const timers = {
//   popup: 1000,
//   worning: 2000,
//   success: 2000,
//   error: 2000,
//   print: 1000,
//   floorPage: 3000,
//   bookingsDetail: 10000
// }

const App = () => {
  const dispatch = useDispatch();
  const mapUrls = useSelector((state) => state.seatsReducer.seatsMapUrls);
  const hcLiveSync = useSelector((state) => state.setupReducer.hcLiveSync);

  useEffect(() => {
    invoke("get_server_url").then((result) => {
      changeAPIBaseUrl(result);
      dispatch(hcLiveSyncFun());
    });
  }, []);

  useEffect(() => {
    dispatch(setTimers(timers));
    const interval = setInterval(() => { dispatch(hcLiveSyncFun()); }, 2000);
    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    if (hcLiveSync?.softAvailability?.length) {
      hcLiveSync?.softAvailability?.forEach(SE => {
        SE?.SECTORS.forEach(e => {
          if (!mapUrls || mapUrls[e.SECTOR_IMAGE] == null) {
            console.log(mapUrls, "@@@@@@@@@@@@@@@@", e.SECTOR_IMAGE, (mapUrls) ? mapUrls[e.SECTOR_IMAGE] : "")
            dispatch(actionSeatsMapLoad(e.SECTOR_IMAGE));
          }
        });
      });
    }
    if(hcLiveSync?.config?.Timers){
      dispatch(setTimers(hcLiveSync?.config?.Timers));
    }
  }, [hcLiveSync]);

  return (
    <BrowserRouter>
      <Layout>
        <ToastContainer
          position="bottom-center"
          autoClose={timers?.error * 1000}
          hideProgressBar={false}
          // newestOnTop={false}
          rtl={false}
          theme="light"
        />
        <Router />
      </Layout>
    </BrowserRouter>
  );
};

export default App;

import axios from "axios"
import axiosTauriApiAdapter from 'axios-tauri-api-adapter';

let BASEURL = "";

export const API = axios.create({adapter: axiosTauriApiAdapter});
export const API2 = axios.create({adapter: axiosTauriApiAdapter});

export const getHcLiveSync = () => API.get('sync')
export const getHcLogin = (username) => API.get(`USER/validate/${username}`)
export const getSeatsList = (sector) => API.get(`MAP/deepAvailabilityBySector/${sector}`)
export const getMap = (url) => API.get("MAP/sector_image?map="+url, {responseType: 'json', timeout: 100000})
export const seatBooking = (data) => API.post('SEAT/bookSeat', data)
export const returnSeat = (data) => API.post('SEAT/returnSeat', data)
export const extendSeat = (data) => API.post("SEAT/extendSeat", data)
export const changeSeat = (data) => API.post('SEAT/moveSeat', data)
export const confirmSeat = (data) => API.post('SEAT/confirmSeat', data)

export const printTickit = (data) => API.post(window.ElectronAPI ?
    `http://localhost:${window.ElectronAPI.wnConfig.CLIENT_PORT}/kiosk/print` : "http://172.17.208.1:8083/kiosk/", data) 

// Change API base url
export const changeAPIBaseUrl = (url) => {
    API.defaults.baseURL = url
    BASEURL = url
}

// Get API base url
export const getAPIBaseUrl = () => BASEURL.replace("/V1/KIOSK","")

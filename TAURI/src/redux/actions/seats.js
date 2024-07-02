import * as api from '../../apis'


export const actionSeatsList = (sector) => async (dispatch) => {
    try {
        // dispatch({ type: 'SET_SeatsListBySector', data: null })
        const res = await api.getSeatsList(sector)
        dispatch({ type: 'SET_SeatsListBySector', data: res.data.result })
    } catch (error) {
        // dispatch({ type: 'SET_SeatsListBySector', data: null })
        console.log(error)
    }
}

export const actionSeatsMapLoad = (path) => async (dispatch) => {
    try {
        console.log("Requesting Map: "+path);
        
        dispatch({ type: 'ADD_SeatsMapUrl', data: {url:path,data:"WAIT"} })
        
        const res = await api.getMap(path)
        let d = res?.data
        // d = JSON.parse(d)
        if ((d?.result?.length || 0) > 0) {
            dispatch({ type: 'ADD_SeatsMapUrl', data: {url:path,data:d?.result} })
        } else {
            console.log("Could Not Download")
        }
    } catch (error) {
        dispatch({ type: 'ADD_SeatsMapUrl', data: {url:path,data:null} })
        
        console.log("ERROR", error)
    }
}